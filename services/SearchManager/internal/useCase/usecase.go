package useCase

import (
	"context"
	"final/services/SearchManager/internal/domain"
	"final/services/SearchManager/internal/repository"
)

type SearchUseCase interface {
	Get(context.Context, *domain.Filter) ([]*domain.Book, error)
}

type searchUseCase struct {
	searchRepository repository.SearchRepository
}

func (s *searchUseCase) Get(ctx context.Context, filter *domain.Filter) ([]*domain.Book, error) {
	books, err := s.searchRepository.Get(ctx, filter)
	if err != nil {
		return nil, err
	}
	return books, nil
}

func NewSearchUseCase(searchRep repository.SearchRepository) SearchUseCase {
	return &searchUseCase{searchRepository: searchRep}
}
