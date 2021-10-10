package service

import (
	"encoding/json"
	"movie-service/domain"
	"movie-service/repository"
	"net/http"
	"strconv"
)

type serviceMovieRepository struct {
	client             *http.Client
	baseUrl, secretKey string
}

func NewServiceMovieRepository(client *http.Client, baseUrl, secretKey string) domain.MovieRepository {
	return &serviceMovieRepository{
		client:    client,
		baseUrl:   baseUrl,
		secretKey: secretKey,
	}
}

func (r *serviceMovieRepository) doApiCall(method, endpoint string, resp interface{}, queryParam map[string]string) error {
	queryParam["apikey"] = r.secretKey
	req, err := repository.GenerateApiRequest(method, r.baseUrl, endpoint, nil, queryParam)
	if err != nil {
		return err
	}
	clientResp, err := r.client.Do(req)
	if err != nil {
		return err
	}
	defer clientResp.Body.Close()

	return json.NewDecoder(clientResp.Body).Decode(&resp)
}

func (r *serviceMovieRepository) GetByTitleAndPage(title string, page int) ([]domain.Movie, error) {
	var movieResponse domain.ListMovie
	queryParam := map[string]string{
		"s":    title,
		"page": strconv.Itoa(page),
	}
	err := r.doApiCall(http.MethodGet, "", &movieResponse, queryParam)
	return movieResponse.Search, err
}

func (r *serviceMovieRepository) GetById(id string) (domain.Movie, error) {
	var movie domain.Movie

	queryParam := map[string]string{
		"i": id,
	}
	err := r.doApiCall(http.MethodGet, "", &movie, queryParam)
	return movie, err
}
