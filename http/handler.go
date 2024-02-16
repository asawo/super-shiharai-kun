package http

import (
	"context"
	"net/http"

	"github.com/morikuni/failure"

	"github.com/asawo/api/errors"
)

func (s *Server) invoicesHandler(ctx context.Context) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			listInvoices(s, w, r)
		case http.MethodPost:
			createInvoice(s, w, r)
		default:
			s.handleError(ctx, w, r, errors.New(errors.MethodNotAllowed, failure.Messagef("method %s not allowed", r.Method)))
		}
	})
}
