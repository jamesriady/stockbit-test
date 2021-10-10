package usecase

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"movie-service/domain"
	"movie-service/domain/mocks"
	"movie-service/pkg/log"
	"testing"
)

func TestGetByTitleAndPage(t *testing.T) {
	log.InitializeLog()

	ctrl := gomock.NewController(t)
	mockMovieRepository := mocks.NewMockMovieRepository(ctrl)

	t.Run("success", func(t *testing.T) {
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
		mockMovies := []domain.Movie{mockMovie}

		title := "Batman"
		page := 1
		mockMovieRepository.EXPECT().GetByTitleAndPage(title, page).Return(mockMovies, nil)
		mockMovieRepository.EXPECT().GetById("tt0094712").Return(mockMovie, nil)
		usecase := NewMovieUsecase(mockMovieRepository)
		movies, err := usecase.GetByTitleAndPage(title, page)
		assert.Equal(t, movies, mockMovies)
		assert.NoError(t, err)
		assert.Len(t, movies, len(mockMovies))
	})

	t.Run("failed", func(t *testing.T) {
		title := "Batman"
		page := 1
		var movies []domain.Movie
		mockMovieRepository.EXPECT().GetByTitleAndPage(title, page).Return(movies, errors.New("got error"))
		usecase := NewMovieUsecase(mockMovieRepository)
		expectedMovies, err := usecase.GetByTitleAndPage(title, page)
		assert.Equal(t, expectedMovies, movies)
		assert.Error(t, err)
	})
}
