package main

import (
	"final/pkg/messages/library/book"
	"final/pkg/postgres"
	"final/services/BookManager/internal/delivery"
	"final/services/BookManager/internal/repository"
	"final/services/BookManager/internal/useCase"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	db, err := postgres.OpenDb("localhost", "5432", "postgres", "123456", "library")
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	srv := delivery.NewBookManagerServer(useCase.NewBookUseCase(repository.NewBookRepository(db)))
	book.RegisterBookManagerServer(s, srv)

	l, err := net.Listen("tcp", ":4000")
	if err != nil {
		log.Fatal(err)
	}

	if err = s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
