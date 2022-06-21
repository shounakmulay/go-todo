package resmodel

import (
	"net/http"
)

type Response struct {
	Status int `json:"status"`
	Body   any `json:"body"`
}

func Success(data any) Response {
	return Response{
		Status: http.StatusOK,
		Body:   data,
	}
}
