package http

import (
	"fmt"
	"net/http"

	"github.com/morikuni/failure"

	"github.com/asawo/api/errors"
	"github.com/asawo/api/service"
)

func listInvoices(s *Server, w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	req, serr := validateListInvoicesRequest(r)
	if serr != nil {
		s.handleError(ctx, w, r, serr)
		return
	}

	invoices, serr := s.service.ListInvoices(ctx, req)
	if serr != nil {
		s.handleError(ctx, w, r, serr)
		return
	}
	response := ListInvoicesResponse{
		Result: "OK",
		Data:   invoices,
	}

	if err := writeJSON(w, http.StatusOK, &response); err != nil {
		s.logger.Error(fmt.Sprintf("failed to write response: %v", err))
	}
}

func validateListInvoicesRequest(r *http.Request) (*service.ListInvoicesRequest, errors.ServiceError) {
	if r.Method != http.MethodGet {
		return nil, errors.New(errors.MethodNotAllowed, failure.Messagef("status method %s is not allowed", r.Method))
	}

	// Authorization
	email, password, ok := r.BasicAuth()
	if !ok {
		return nil, errors.New(errors.Unauthorized, failure.Message("invalid basic auth header"))
	}

	// Get issue date and/or due date
	startDate, endDate, serr := GetQueryParamsForListInvoices(r)
	if serr != nil {
		return nil, serr
	}

	// Validate startDate is before endDate
	if !startDate.IsZero() && !endDate.IsZero() && startDate.After(endDate) {
		return nil, errors.New(errors.InvalidRequest, failure.Messagef("start_date %s is after end_date %s", startDate.Format("2006-01-02"), endDate.Format("2006-01-02")))
	}

	return &service.ListInvoicesRequest{
		Email:     email,
		Password:  password,
		StartDate: startDate,
		EndDate:   endDate,
	}, nil
}
