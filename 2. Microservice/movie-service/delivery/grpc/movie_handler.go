package grpc

import (
	"context"
	"encoding/json"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
	server_grpc "movie-service/delivery/grpc/proto"
	"movie-service/domain"
)

type server struct {
	mUsecase domain.MovieUsecase
}

func NewMovieServerGrpc(gServer *grpc.Server, movieUsecase domain.MovieUsecase) {
	movieServer := &server{
		mUsecase: movieUsecase,
	}
	server_grpc.RegisterMovieHandlerServer(gServer, movieServer)
	reflection.Register(gServer)
}

func (s *server) transformMovieData(movie *domain.Movie) (*server_grpc.Movie, error) {
	movieInByte, err := json.Marshal(movie)
	if err != nil {
		return nil, err
	}
	messageMovie := &server_grpc.Movie{}
	err = protojson.Unmarshal(movieInByte, messageMovie)
	return messageMovie, nil
}

func (s *server) FetchMovie(ctx context.Context, request *server_grpc.FetchMovieRequest) (*server_grpc.MovieData, error) {
	if request == nil {
		return nil, domain.ErrInvalidParam
	}

	title := request.GetTitle()
	if title == "" || len(title) < 3 {
		return nil, domain.ErrParamTitleMinCharacter
	}

	page := request.GetPage()
	if page <= 0 {
		page = 1
	}

	movies, err := s.mUsecase.GetByTitleAndPage(title, int(page))
	if err != nil {
		return nil, err
	}

	movieMessages := make([]*server_grpc.Movie, len(movies))
	for i, movie := range movies {
		msg, err := s.transformMovieData(&movie)
		if err != nil {
			return nil, err
		}
		movieMessages[i] = msg
	}

	return &server_grpc.MovieData{
		Movies: movieMessages,
	}, nil
}
