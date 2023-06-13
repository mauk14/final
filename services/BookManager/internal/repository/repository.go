package repository

import (
	"context"
	"errors"
	"final/services/BookManager/internal/domain"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	ErrEditConflict   = errors.New("edit conflict")
	ErrRecordNotFound = errors.New("record not found")
)

type bookRepository struct {
	DB *pgxpool.Pool
}

type BookRepository interface {
	Insert(context.Context, *domain.Book) error
	Update(context.Context, *domain.Book) error
	Delete(context.Context, int64) error
}

func NewBookRepository(db *pgxpool.Pool) BookRepository {
	return &bookRepository{DB: db}
}

func (b *bookRepository) Insert(ctx context.Context, book *domain.Book) error {
	query := `INSERT INTO books (title,author, year, genres)
VALUES ($1, $2, $3, $4)
RETURNING id, created_at`

	args := []any{book.Title, book.Author, book.Year, book.Genres}

	return b.DB.QueryRow(ctx, query, args...).Scan(&book.Id, &book.CreateAt)
}

func (b *bookRepository) Update(ctx context.Context, book *domain.Book) error {
	query := `
UPDATE books
SET title = $1,author = $2, year = $3, genres = $4
WHERE id = $5`

	args := []any{
		book.Title,
		book.Author,
		book.Year,

		book.Genres,
		book.Id,
	}

	err := b.DB.QueryRow(ctx, query, args...).Scan()

	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			return ErrEditConflict
		default:
			return err
		}
	}

	return nil

}

func (b *bookRepository) Delete(ctx context.Context, id int64) error {

	if id < 1 {
		return ErrRecordNotFound
	}

	query := `
DELETE FROM books
WHERE id = $1`

	result, err := b.DB.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return ErrRecordNotFound
	}

	return nil

}
