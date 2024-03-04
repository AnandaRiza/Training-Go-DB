package pkg

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	_ "github.com/joho/godotenv/autoload"
)

type Payload struct {
	Email string
	jwt.RegisteredClaims
}

func NewPayLoad(email string) *Payload {
	return &Payload{Email: email, RegisteredClaims: jwt.RegisteredClaims{
		Issuer:    os.Getenv("JWT_ISSUER"),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 2)),
	}}

}

func (p *Payload) CreateToken() (string, error) {
	jwtSecret := os.Getenv("JWT_SECRET")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, p)
	result, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}
	return result, nil
}

func VerifyToken(token string) (*Payload, error) {
	jwtSecret := os.Getenv("JWT_SECRET")
	parsedToken, err := jwt.ParseWithClaims(token, &Payload{}, func(t *jwt.Token) (interface{}, error) {
		// ambil dulu jwt secret nya
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return nil, err
	}

	payload := parsedToken.Claims.(*Payload)
	return payload, nil
}
