package main

import (
	"context"
	"final/services/ParserManager/proto/parser"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial(":4002", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	c := parser.NewParserManagerClient(conn)

	books, err := c.GetNewBooks(context.Background(), &parser.RequestRM{})

	fmt.Println(books)
}
