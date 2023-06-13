package main

import (
	"final/services/ParserManager/internal/delivery"
	"final/services/ParserManager/internal/useCase"
	"final/services/ParserManager/proto/parser"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	s := grpc.NewServer()
	srv := delivery.NewParserManagerServer(useCase.NewParserUseCase())
	parser.RegisterParserManagerServer(s, srv)

	l, err := net.Listen("tcp", ":4002")
	if err != nil {
		log.Fatal(err)
	}

	if err = s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
