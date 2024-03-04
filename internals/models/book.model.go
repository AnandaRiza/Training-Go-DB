package models

type BookModel struct {
	Id          int     `db:"id"`
	Title       string  `db:"title" json:"title" form:"title"`
	Description *string `db:"description" json:"description" form:"description"`
	Author      string  `db:"author" json:"author" form:"author"`
	Picture     *string `db:"picture" json:"picture" form:"picture"`
}
