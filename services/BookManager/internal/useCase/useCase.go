package useCase

import (
	"context"
	"final/services/BookManager/internal/domain"
	"final/services/BookManager/internal/repository"
)

type BookUseCase interface {
	Create(context.Context, *domain.Book) error
	Update(context.Context, *domain.Book) error
	Delete(context.Context, int64) error
}

type bookUseCase struct {
	bookRep repository.BookRepository
}

func (b *bookUseCase) Create(ctx context.Context, book *domain.Book) error {
	return b.bookRep.Insert(ctx, book)
}

func (b *bookUseCase) Update(ctx context.Context, book *domain.Book) error {
	return b.bookRep.Update(ctx, book)
}

func (b *bookUseCase) Delete(ctx context.Context, id int64) error {
	return b.bookRep.Delete(ctx, id)
}

func NewBookUseCase(bookRep repository.BookRepository) BookUseCase {
	return &bookUseCase{bookRep: bookRep}
}
