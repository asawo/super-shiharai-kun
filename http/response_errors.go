package http

import (
	"context"
	"fmt"
	"net/http"

	"github.com/asawo/api/errors"
)

type ErrorResponse struct {
	Result  string `json:"result"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (s *Server) handleError(ctx context.Context, w http.ResponseWriter, r *http.Request, serr errors.ServiceError) {
	s.logger.Error(fmt.Sprintf("%+v", serr))
	httpStatusCode := getHttpStatus(serr)

	respErr := &ErrorResponse{
		Result:  "error",
		Code:    getHttpStatus(serr),
		Message: serr.Message(),
	}

	if err := writeJSON(w, httpStatusCode, respErr); err != nil {
		s.logger.Error(fmt.Sprintf("failed to write response: %v", err))
	}
}
