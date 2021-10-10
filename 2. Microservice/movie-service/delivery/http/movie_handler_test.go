package http

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"movie-service/domain"
	"movie-service/domain/mocks"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchMovieSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)

	title := "Batman"
	mockMovieUsecase := mocks.NewMockMovieUsecase(ctrl)
	handler := movieHandler{mUsecase: mockMovieUsecase}
	mockMovieUsecase.EXPECT().GetByTitleAndPage(title, 1).Return([]domain.Movie{}, nil)
	router := httprouter.New()
	router.GET("/movies", handler.FetchMovie)

	req, err := http.NewRequest(http.MethodGet, "/movies?title="+title, nil)
	assert.NoError(t, err)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	result := resp.Result()
	assert.Equal(t, http.StatusOK, result.StatusCode)
}

func TestFetchMovieParamNotValid(t *testing.T) {
	ctrl := gomock.NewController(t)

	title := "Ba"
	mockMovieUsecase := mocks.NewMockMovieUsecase(ctrl)
	handler := movieHandler{mUsecase: mockMovieUsecase}
	mockMovieUsecase.EXPECT().GetByTitleAndPage(title, 1).Return([]domain.Movie{}, nil)
	router := httprouter.New()
	router.GET("/movies", handler.FetchMovie)

	req, err := http.NewRequest(http.MethodGet, "/movies?title="+title, nil)
	assert.NoError(t, err)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	result := resp.Result()
	assert.Equal(t, http.StatusBadRequest, result.StatusCode)
}

func TestFetchMovieFailed(t *testing.T) {
	ctrl := gomock.NewController(t)

	title := "Batman"
	mockMovieUsecase := mocks.NewMockMovieUsecase(ctrl)
	handler := movieHandler{mUsecase: mockMovieUsecase}
	mockMovieUsecase.EXPECT().GetByTitleAndPage(title, 1).Return(nil, errors.New("got error"))
	router := httprouter.New()
	router.GET("/movies", handler.FetchMovie)

	req, err := http.NewRequest(http.MethodGet, "/movies?title="+title, nil)
	assert.NoError(t, err)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	result := resp.Result()
	assert.Equal(t, http.StatusInternalServerError, result.StatusCode)
}
