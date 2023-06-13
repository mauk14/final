package main

import (
	"context"
	"final/services/SearchManager/proto/search"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

func main() {
	title := "chert"
	conn, err := grpc.Dial(":4001", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	c := search.NewSearchManagerClient(conn)

	books, err := c.Get(context.Background(), &search.Filter{Title: title, Size: 2, Sort: "id"})

	fmt.Println(books)
}
