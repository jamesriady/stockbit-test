package usecase

import (
	"movie-service/domain"
	"movie-service/pkg/log"
	"sync"
)

type movieUsecase struct {
	movieRepository domain.MovieRepository
}

func NewMovieUsecase(movieRepository domain.MovieRepository) domain.MovieUsecase {
	return &movieUsecase{movieRepository: movieRepository}
}

func (u *movieUsecase) getMovieDetail(wg *sync.WaitGroup, movieId string, movieCh chan domain.Movie) {
	defer wg.Done()

	log.Get().Infof("[START] get movie by id %s", movieId)
	movie, err := u.movieRepository.GetById(movieId)
	if err != nil {
		log.Get().Errorf("failed to get movie on id %s. Err: %s", movieId, err.Error())
		return
	}
	log.Get().Infof("[END] success getting movie on id %s", movie.IdmbId)
	movieCh <- movie
}

func (u *movieUsecase) GetByTitleAndPage(title string, page int) ([]domain.Movie, error) {
	log.Get().Infof("[START] get list of movies by title %s, page %d", title, page)
	listMovie, err := u.movieRepository.GetByTitleAndPage(title, page)
	if err != nil {
		log.Get().Infof("failed to get list of movies by title %s, page %d. Err: %s", title, page, err.Error())
		return nil, err
	}
	log.Get().Infof("[END] success getting list of movies by title %s, page %d", title, page)
	wg := sync.WaitGroup{}
	wg.Add(len(listMovie))
	movieCh := make(chan domain.Movie, len(listMovie))

	for _, movie := range listMovie {
		go u.getMovieDetail(&wg, movie.IdmbId, movieCh)
	}

	wg.Wait()
	close(movieCh)
	var movies []domain.Movie
	for movie := range movieCh {
		movies = append(movies, movie)
	}

	return movies, nil
}
