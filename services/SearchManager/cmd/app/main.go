package main

import (
	"final/pkg/postgres"
	"final/services/SearchManager/internal/delivery"
	"final/services/SearchManager/internal/repository"
	"final/services/SearchManager/internal/useCase"
	"final/services/SearchManager/proto/search"
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
	srv := delivery.NewSearchManagerServer(useCase.NewSearchUseCase(repository.NewSearchRepository(db)))
	search.RegisterSearchManagerServer(s, srv)

	l, err := net.Listen("tcp", ":4001")
	if err != nil {
		log.Fatal(err)
	}

	if err = s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
