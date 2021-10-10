package http

import (
	"github.com/julienschmidt/httprouter"
	"movie-service/domain"
	"movie-service/pkg/httpjson"
	"net/http"
	"strconv"
)

type movieHandler struct {
	mUsecase domain.MovieUsecase
}

func NewMovieHandler(router *httprouter.Router, movieUsecase domain.MovieUsecase) {
	handler := &movieHandler{
		mUsecase: movieUsecase,
	}
	router.GET("/movies", handler.FetchMovie)
}

func (h *movieHandler) FetchMovie(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	q := r.URL.Query()
	title := q.Get("title")
	if title == "" || len(title) < 3 {
		httpjson.WriteErrorResponse(w, http.StatusBadRequest, domain.ErrParamTitleMinCharacter)
		return
	}

	pageStr := q.Get("page")
	if pageStr == "" {
		pageStr = "1"
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}
	resp, err := h.mUsecase.GetByTitleAndPage(title, page)
	if err != nil {
		httpjson.WriteErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	httpjson.WriteOKResponse(w, resp)
}
