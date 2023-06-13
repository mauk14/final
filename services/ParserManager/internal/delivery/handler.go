package delivery

import (
	"context"
	"final/services/ParserManager/internal/domain"
	"final/services/ParserManager/internal/useCase"
	"final/services/ParserManager/proto/parser"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type GrpcServer struct {
	parserUseCase useCase.ParserUseCase
}

func (s *GrpcServer) transformBookRPC(bk *domain.Book) *parser.Book {

	if bk == nil {
		return nil
	}

	createAt := &timestamppb.Timestamp{
		Seconds: bk.CreateAt.Unix(),
	}
	res := &parser.Book{
		Id:       bk.Id,
		Title:    bk.Title,
		Author:   bk.Author,
		Genres:   bk.Genres,
		Year:     bk.Year,
		CreateAt: createAt,
	}
	return res
}

func (s *GrpcServer) GetNewBooks(ctx context.Context, _ *parser.RequestRM) (*parser.ResponseRM, error) {
	books, err := s.parserUseCase.GetNewBooks(ctx)
	if err != nil {
		return &parser.ResponseRM{Books: nil, Errors: []*parser.ErrorMessage{&parser.ErrorMessage{Message: err.Error()}}}, err
	}
	var booksPf []*parser.Book
	for i := range books {
		booksPf = append(booksPf, s.transformBookRPC(books[i]))
	}

	return &parser.ResponseRM{Books: booksPf, Errors: nil}, nil
}

func NewParserManagerServer(parserUseCase useCase.ParserUseCase) parser.ParserManagerServer {
	return &GrpcServer{parserUseCase: parserUseCase}
}
