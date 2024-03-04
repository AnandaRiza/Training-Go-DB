package repositories

import (
	"bcas/bookstore-go/internals/models"
	"strings"

	"github.com/jmoiron/sqlx"
)

type BookRepo struct {
	*sqlx.DB
}

func InitBookRepo(db *sqlx.DB) *BookRepo {
	return &BookRepo{db}
}

func (b *BookRepo) FindAll() ([]models.BookModel, error) {
	query := "SELECT * FROM books"
	result := []models.BookModel{}
	if err := b.Select(&result, query); err != nil {
		return nil, err
	}
	return result, nil
}

func (b *BookRepo) SaveBook(body models.BookModel) error {

	query := "INSERT INTO books(title, description, author) VALUES (?,?,?)"
	if _, err := b.Exec(query, body.Title, body.Description, body.Author); err != nil {
		return err
	}
	return nil

}

func (b *BookRepo) FindbyId(id int) ([]models.BookModel, error) {
	//search for books id from the books table
	query := "SELECT * FROM books WHERE id = ?"
	result := []models.BookModel{}
	if err := b.Select(&result, query); err != nil {
		return nil, err
	}
	return result, nil

}

func (b *BookRepo) DeletebyId(id int) error {
	//delete books id from books table
	query := "DELETE * FROM books WHERE id = ?"
	if _, err := b.Exec(query, id); err != nil {
		return err
	}
	return nil

}

// UpdateById updates a book record by its ID in the database.
func (b *BookRepo) UpdateById(id int, body models.BookModel) error {
	// Construct the SQL query
	query := "UPDATE books SET"
	var args []interface{}

	// Check if the title field is not empty
	if body.Title != "" {
		query += " title = ?,"
		args = append(args, body.Title)
	}

	// Check if the description field is not empty
	if body.Description != nil && *body.Description != "" {
		query += " description = ?,"
		args = append(args, *body.Description)
	}

	// Check if the author field is not empty
	if body.Author != "" {
		query += " author = ?,"
		args = append(args, body.Author)
	}

	// Remove the trailing comma from the query
	query = strings.TrimSuffix(query, ",")

	// Add the WHERE clause to the query
	query += " WHERE id = ?"
	args = append(args, id)

	// Execute the parameterized query using Exec function
	if _, err := b.Exec(query, args...); err != nil {
		return err
	}

	return nil
}
