package http

import (
	"net/http"

	"github.com/morikuni/failure"

	"github.com/asawo/api/errors"
)

func getHttpStatus(err errors.ServiceError) int {
	errMap := errorMap()
	statusCode, ok := errMap[err.GetCode()]

	if !ok {
		statusCode = http.StatusInternalServerError
	}

	return statusCode
}

func errorMap() map[failure.Code]int {
	// Define only the errors that support HTTP status codes
	return map[failure.Code]int{
		errors.TemporaryUnavailable: http.StatusServiceUnavailable,
		errors.Timeout:              http.StatusRequestTimeout,
		errors.InvalidRequest:       http.StatusBadRequest,
		errors.NotFound:             http.StatusNotFound,
		errors.Unauthorized:         http.StatusUnauthorized,
		errors.PermissionDenied:     http.StatusForbidden,
		errors.MethodNotAllowed:     http.StatusMethodNotAllowed,
		errors.Internal:             http.StatusInternalServerError,
		errors.Unknown:              http.StatusInternalServerError,
		errors.AlreadyExists:        http.StatusConflict,
	}
}
