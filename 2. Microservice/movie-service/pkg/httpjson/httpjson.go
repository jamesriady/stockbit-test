package httpjson

import (
	"encoding/json"
	"net/http"
)

type ResponseData struct {
	Data interface{} `json:"data"`
}

type ResponseDataError struct {
	Data   interface{} `json:"data"`
	Errors interface{} `json:"errors,omitempty"`
}

type ResponseError struct {
	Errors interface{} `json:"errors"`
}

type ErrorDetail struct {
	Message string `json:"message"`
	Code    int    `json:"code,omitempty"`
}

type response struct {
	status int
	data   interface{}
}

func writeFailedEncode(w http.ResponseWriter) {
	_ = json.NewEncoder(w).Encode(&ResponseError{Errors: &ErrorDetail{Message: "failed to encode response"}})
}

func writeResponse(w http.ResponseWriter, resp response) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(resp.status)
	err := json.NewEncoder(w).Encode(resp.data)
	if err != nil {
		writeFailedEncode(w)
	}

}

func WriteOKResponse(w http.ResponseWriter, data interface{}) {
	data = &ResponseData{
		Data: data,
	}
	writeResponse(w, response{http.StatusOK, data})
}

func WriteErrorResponse(w http.ResponseWriter, statusCode int, err error) {
	data := &ResponseError{
		Errors: &ErrorDetail{
			Message: err.Error(),
		},
	}

	writeResponse(w, response{statusCode, data})
}
