package main

import (
	"github.com/julienschmidt/httprouter"
	"google.golang.org/grpc"
	"movie-service/config"
	service_grpc "movie-service/delivery/grpc"
	service_http "movie-service/delivery/http"
	"movie-service/repository/service"
	"movie-service/usecase"
	"net"
	"net/http"

	"movie-service/pkg/log"
)

func main() {
	log.InitializeLog()
	config.InitializeAppConfig()

	movieServiceCfg := config.AppConfig.MovieService
	client := &http.Client{
		Timeout: movieServiceCfg.TimeoutInSec.Duration,
	}
	movieRepository := service.NewServiceMovieRepository(client, movieServiceCfg.BaseUrl, movieServiceCfg.SecretKey)
	movieUsecase := usecase.NewMovieUsecase(movieRepository)

	listener, err := net.Listen("tcp", ":"+config.AppConfig.GRPC.Port)
	if err != nil {
		panic(err)
	}

	gServer := grpc.NewServer()
	service_grpc.NewMovieServerGrpc(gServer, movieUsecase)
	log.Get().Info("GRPC Server run at port ", config.AppConfig.GRPC.Port)

	go func() {
		err = gServer.Serve(listener)
		if err != nil {
			panic(err)
		}
	}()

	router := httprouter.New()
	service_http.NewMovieHandler(router, movieUsecase)
	log.Get().Fatal(http.ListenAndServe(":8080", router))
}
