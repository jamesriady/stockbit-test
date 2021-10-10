package domain

type ListMovie struct {
	Search       []Movie `json:"Search"`
	TotalResults string  `json:"totalResults"`
	Response     string  `json:"Response"`
}

type Movie struct {
	Title      string        `json:"Title"`
	Year       string        `json:"Year"`
	Rated      string        `json:"Rated"`
	Released   string        `json:"Released"`
	Runtime    string        `json:"Runtime"`
	Genre      string        `json:"Genre"`
	Director   string        `json:"Director"`
	Writer     string        `json:"Writer"`
	Actors     string        `json:"Actors"`
	Plot       string        `json:"Plot"`
	Language   string        `json:"Language"`
	Country    string        `json:"Country"`
	Awards     string        `json:"Awards"`
	Poster     string        `json:"Poster"`
	Ratings    []MovieRating `json:"Ratings"`
	Metascore  string        `json:"Metascore"`
	ImdbRating string        `json:"imdbRating"`
	ImdbVotes  string        `json:"imdbVotes"`
	IdmbId     string        `json:"imdbID"`
	Type       string        `json:"Type"`
	DVD        string        `json:"DVD"`
	BoxOffice  string        `json:"BoxOffice"`
	Production string        `json:"Production"`
	Website    string        `json:"Website"`
	Response   string        `json:"Response"`
}

type MovieRating struct {
	Source string `json:"source"`
	Value  string `json:"value"`
}
type MovieUsecase interface {
	GetByTitleAndPage(title string, page int) ([]Movie, error)
}

type MovieRepository interface {
	GetByTitleAndPage(title string, page int) ([]Movie, error)
	GetById(id string) (Movie, error)
}
