package service

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"movie-service/domain"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetById(t *testing.T) {
	mockMovie := domain.Movie{
		Title:    "Bat*asd21",
		Year:     "1988",
		Rated:    "R",
		Released: "21 Oct 1988",
		Runtime:  "105 min",
		Genre:    "Drama, War",
		Director: "Peter Markle",
		Writer:   "William C. Anderson, George Gordon",
		Actors:   "Gene Hackman, Danny Glover, Jerry Reed",
		Plot:     "During the Vietnam War, Colonel Hambleton's aircraft is shot down over enemy territory and a frantic rescue operation ensues.",
		Language: "English",
		Country:  "United States",
		Awards:   "1 nomination",
		Poster:   "https://m.media-amazon.com/images/M/MV5BZDRmNjYwZDktOTYxZi00MTdlLWI5ZjYtYWU4MDE5MDc5NGM3L2ltYWdlXkEyXkFqcGdeQXVyNjQzNDI3NzY@._V1_SX300.jpg",
		Ratings: []domain.MovieRating{
			{
				Source: "Internet Movie Database",
				Value:  "6.5/10",
			},
		},
		Metascore:  "58",
		ImdbRating: "6.5",
		ImdbVotes:  "8,496",
		IdmbId:     "tt0094712",
		Type:       "movie",
		DVD:        "24 Apr 2007",
		BoxOffice:  "$3,966,256",
		Production: "TriStar Pictures",
		Website:    "N/A",
		Response:   "True",
	}

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.URL.String(), "/?apikey=secret&i=tt0094712")
		responseInByte, _ := json.Marshal(mockMovie)
		rw.Write(responseInByte)
	}))

	defer server.Close()

	repository := NewServiceMovieRepository(server.Client(), server.URL, "secret")
	resp, err := repository.GetById("tt0094712")
	assert.NoError(t, err)
	assert.Equal(t, mockMovie, resp)
}

func TestGetByTitleAndPage(t *testing.T) {
	mockMovies := []domain.Movie{{
		Title:    "Bat*asd21",
		Year:     "1988",
		Rated:    "R",
		Released: "21 Oct 1988",
		Runtime:  "105 min",
		Genre:    "Drama, War",
		Director: "Peter Markle",
		Writer:   "William C. Anderson, George Gordon",
		Actors:   "Gene Hackman, Danny Glover, Jerry Reed",
		Plot:     "During the Vietnam War, Colonel Hambleton's aircraft is shot down over enemy territory and a frantic rescue operation ensues.",
		Language: "English",
		Country:  "United States",
		Awards:   "1 nomination",
		Poster:   "https://m.media-amazon.com/images/M/MV5BZDRmNjYwZDktOTYxZi00MTdlLWI5ZjYtYWU4MDE5MDc5NGM3L2ltYWdlXkEyXkFqcGdeQXVyNjQzNDI3NzY@._V1_SX300.jpg",
		Ratings: []domain.MovieRating{
			{
				Source: "Internet Movie Database",
				Value:  "6.5/10",
			},
		},
		Metascore:  "58",
		ImdbRating: "6.5",
		ImdbVotes:  "8,496",
		IdmbId:     "tt0094712",
		Type:       "movie",
		DVD:        "24 Apr 2007",
		BoxOffice:  "$3,966,256",
		Production: "TriStar Pictures",
		Website:    "N/A",
		Response:   "True",
	}}

	listMovie := domain.ListMovie{
		Search: mockMovies,
	}

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.URL.String(), "/?apikey=secret&page=1&s=Batman")
		responseInByte, _ := json.Marshal(listMovie)
		rw.Write(responseInByte)
	}))

	defer server.Close()

	repository := NewServiceMovieRepository(server.Client(), server.URL, "secret")
	resp, err := repository.GetByTitleAndPage("Batman", 1)
	assert.NoError(t, err)
	assert.Equal(t, mockMovies, resp)
}
