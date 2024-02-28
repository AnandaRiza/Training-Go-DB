package models

type AuthModel struct {
	Id       int    `db:"id"`
	Email    string `db:"email" json:"email" form:"email"`
	Password string `db:"password" json:"password" form:"password"`
}
