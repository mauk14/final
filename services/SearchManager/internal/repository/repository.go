package repository

import (
	"context"
	"final/services/SearchManager/internal/domain"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

type SearchRepository interface {
	Get(context.Context, *domain.Filter) ([]*domain.Book, error)
}

type seacrhRepository struct {
	Db *pgxpool.Pool
}

func (s *seacrhRepository) Get(ctx context.Context, filter *domain.Filter) ([]*domain.Book, error) {
	query := fmt.Sprintf(`
SELECT id, created_at, title, author, year, genres
FROM books
WHERE (to_tsvector('simple', title) @@ plainto_tsquery('simple', $1) OR $1 = '')
ANd (to_tsvector('simple', author) @@ plainto_tsquery('simple', $2) OR $2 = '')
AND (genres @> $3 OR $3 = '{}')
ORDER BY %s %s, id ASC
LIMIT $4`, filter.SortColumn(), filter.SortDirection())

	args := []any{filter.Title, filter.Author, filter.Genres, filter.Size}

	fmt.Println(filter)

	rows, err := s.Db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	books := []*domain.Book{}

	for rows.Next() {

		var book domain.Book
		err := rows.Scan(
			&book.Id,
			&book.CreateAt,
			&book.Title,
			&book.Author,
			&book.Year,
			&book.Genres,
		)
		if err != nil {
			return nil, err
		}
		books = append(books, &book)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return books, nil

}

func NewSearchRepository(db *pgxpool.Pool) SearchRepository {
	return &seacrhRepository{Db: db}
}
