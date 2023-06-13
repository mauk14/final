package delivery

import (
	"context"
	"final/pkg/messages/library/book"
	"final/services/BookManager/internal/domain"
	"final/services/BookManager/internal/useCase"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type GrpcServer struct {
	bookUseCase useCase.BookUseCase
}

func (s *GrpcServer) transformBookRPC(bk *domain.Book) *book.Book {

	if bk == nil {
		return nil
	}

	createAt := &timestamppb.Timestamp{
		Seconds: bk.CreateAt.Unix(),
	}
	res := &book.Book{
		Id:       bk.Id,
		Title:    bk.Title,
		Author:   bk.Author,
		Genres:   bk.Genres,
		CreateAt: createAt,
	}
	return res
}

func (s *GrpcServer) transformBookData(bk *book.Book) *domain.Book {
	createAt := time.Unix(bk.GetCreateAt().GetSeconds(), 0)
	res := &domain.Book{
		Id:       bk.Id,
		Title:    bk.Title,
		Author:   bk.Author,
		Genres:   bk.Genres,
		Year:     bk.Year,
		CreateAt: createAt,
	}
	return res
}

func (s *GrpcServer) Insert(ctx context.Context, bk *book.Book) (*book.ResponseBM, error) {
	err := s.bookUseCase.Create(ctx, s.transformBookData(bk))
	if err != nil {
		return &book.ResponseBM{TotalSuccess: 0, Errors: []*book.ErrorMessage{{Message: err.Error()}}}, err
	}

	return &book.ResponseBM{TotalSuccess: 1}, nil
}

func (s *GrpcServer) Update(ctx context.Context, bk *book.Book) (*book.ResponseBM, error) {
	err := s.bookUseCase.Update(ctx, s.transformBookData(bk))
	if err != nil {
		return &book.ResponseBM{TotalSuccess: 0, Errors: []*book.ErrorMessage{{Message: err.Error()}}}, err
	}

	return &book.ResponseBM{TotalSuccess: 1}, nil
}

func (s *GrpcServer) Delete(ctx context.Context, dr *book.DeleteRequest) (*book.ResponseBM, error) {
	err := s.bookUseCase.Delete(ctx, dr.Id)
	if err != nil {
		return &book.ResponseBM{TotalSuccess: 0, Errors: []*book.ErrorMessage{{Message: err.Error()}}}, err
	}

	return &book.ResponseBM{TotalSuccess: 1}, nil
}

func NewBookManagerServer(bookUseCase useCase.BookUseCase) book.BookManagerServer {
	return &GrpcServer{bookUseCase: bookUseCase}
}
