package repositories

import (
	"bcas/bookstore-go/internals/models"

	"github.com/jmoiron/sqlx"
)

type AuthRepo struct {
	*sqlx.DB
}

func InitAuthRepo(db *sqlx.DB) *AuthRepo {
	return &AuthRepo{db}
}

func (a *AuthRepo) FindByEmail(body models.AuthModel) ([]models.AuthModel, error) {
	query := "SELECT * FROM users WHERE email = ?"
	result := &[]models.AuthModel{}
	if err := a.Select(result, query, body.Email); err != nil {
		return nil, err
	}
	return *result, nil
}
func (a *AuthRepo) SaveUser(body models.AuthModel) error {
	query := "INSERT INTO `users`( `email`, `password`) VALUES (?,?);"
	if _, err := a.Exec(query, body.Email, body.Password); err != nil {
		return err
	}
	return nil
}
