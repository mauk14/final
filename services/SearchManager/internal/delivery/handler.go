package delivery

import (
	"context"
	"final/services/SearchManager/internal/domain"
	"final/services/SearchManager/internal/useCase"
	"final/services/SearchManager/proto/search"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
)

type GrpcServer struct {
	searchUseCase useCase.SearchUseCase
}

func (s *GrpcServer) transformFilterData(filter *search.Filter) *domain.Filter {
	res := &domain.Filter{
		Title:  filter.Title,
		Author: filter.Author,
		Genres: filter.Genres,
		Size:   filter.Size,
		Sort:   filter.Sort,
	}
	return res
}

func (s *GrpcServer) transformBookRPC(bk *domain.Book) *search.Book {

	if bk == nil {
		return nil
	}

	createAt := &timestamppb.Timestamp{
		Seconds: bk.CreateAt.Unix(),
	}
	res := &search.Book{
		Id:       bk.Id,
		Title:    bk.Title,
		Author:   bk.Author,
		Genres:   bk.Genres,
		Year:     bk.Year,
		CreateAt: createAt,
	}
	return res
}

func (s *GrpcServer) Get(ctx context.Context, filter *search.Filter) (*search.ResponseSM, error) {
	books, err := s.searchUseCase.Get(ctx, s.transformFilterData(filter))
	if err != nil {
		log.Println(err.Error())
		return &search.ResponseSM{Books: nil, Errors: []*search.ErrorMessage{{Message: err.Error()}}}, err
	}
	var booksPf []*search.Book
	for i := range books {
		booksPf = append(booksPf, s.transformBookRPC(books[i]))
	}
	return &search.ResponseSM{Books: booksPf, Errors: nil}, nil
}

func NewSearchManagerServer(searchUseCase useCase.SearchUseCase) search.SearchManagerServer {
	return &GrpcServer{searchUseCase: searchUseCase}
}
