package main

import (
	"context"
	"flag"
	"google.golang.org/grpc"
	server_grpc "movie-service/delivery/grpc/proto"
	"movie-service/pkg/log"
	"time"
)

const timeoutDuration = 5 * time.Second

func main() {
	log.InitializeLog()

	address := flag.String("server", "", "gRPC server in format host:port")
	flag.Parse()

	conn, err := grpc.Dial(*address, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := server_grpc.NewMovieHandlerClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDuration)
	defer cancel()
	resp, err := client.FetchMovie(ctx, &server_grpc.FetchMovieRequest{Title: "Batman", Page: 1})
	if err != nil {
		log.Get().Errorf("failed to fetch movie. Err: %s", err.Error())
		return
	}

	for _, movie := range resp.Movies {
		log.Get().Infof("ID: %s, Title: %s, Rating %s", movie.ImdbID, movie.Title, movie.Ratings)
	}
}
