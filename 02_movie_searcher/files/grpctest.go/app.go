package main

import (
	"context"
	"log"

	"github.com/pikomonde/i-view-stockbit/02_movie_searcher/delivery/grpchandler/moviesearch"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9044", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := moviesearch.NewMovieSearchServiceClient(conn)

	request := &moviesearch.MessageRequest{
		Searchword: "Batman",
		Pagination: 2,
	}

	response, err := c.GetAndSaveMoviesBySearch(context.Background(), request)
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.Response)

}
