package resmodel

import (
	"net/http"
)

type ErrorResponse struct {
	Status  int    `json:"status"`
	Error   string `json:"error"`
	Message string `json:"message"`
}

func BadRequest(err error) *ErrorResponse {
	return errorResponse(err, http.StatusBadRequest)
}

func UnprocessableEntity(err error) *ErrorResponse {
	return errorResponse(err, http.StatusUnprocessableEntity)
}

func InternalServerError(err error) *ErrorResponse {
	return errorResponse(err, http.StatusInternalServerError)
}

func Unauthorized(err error) *ErrorResponse {
	return errorResponse(err, http.StatusUnauthorized)
}

func NotFound(msg string) *ErrorResponse {
	return &ErrorResponse{
		Status:  http.StatusNotFound,
		Error:   http.StatusText(http.StatusNotFound),
		Message: msg,
	}
}

func errorResponse(err error, httpCode int) *ErrorResponse {
	return &ErrorResponse{
		Status:  httpCode,
		Error:   http.StatusText(httpCode),
		Message: err.Error(),
	}
}
