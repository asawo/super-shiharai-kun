package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/asawo/api/errors"
	"github.com/asawo/api/service"
	"github.com/morikuni/failure"
)

type CreateInvoiceRequestBody struct {
	PaymentAmount     float64   `json:"payment_amount"`
	ServiceProviderID uint      `json:"service_provider_id"`
	DueDate           time.Time `json:"due_date"`
}

func createInvoice(s *Server, w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	req, serr := validateCreateInvoiceRequest(r)
	if serr != nil {
		s.handleError(ctx, w, r, serr)
		return
	}

	invoice, serr := s.service.CreateInvoice(ctx, req)
	if serr != nil {
		s.handleError(ctx, w, r, serr)
		return
	}
	response := CreateInvoiceResponse{
		Result: "OK",
		Data:   invoice,
	}

	if err := writeJSON(w, http.StatusOK, &response); err != nil {
		s.logger.Error(fmt.Sprintf("failed to write response: %v", err))
	}
}

func validateCreateInvoiceRequest(r *http.Request) (*service.CreateInvoiceRequest, errors.ServiceError) {
	// Check method
	if r.Method != http.MethodPost {
		return nil, errors.New(errors.MethodNotAllowed, failure.Messagef("status method %s is not allowed", r.Method))
	}

	// Extract content type from header
	ct := r.Header.Get("Content-Type")
	if ct != "" {
		mediaType := strings.ToLower(strings.TrimSpace(strings.Split(ct, ";")[0]))
		if mediaType != "application/json" {
			return nil, errors.New(errors.InvalidRequest, failure.Message("invalid content type header"))
		}
	}

	// Extract email and password from header
	email, password, ok := r.BasicAuth()
	if !ok {
		return nil, errors.New(errors.Unauthorized, failure.Message("invalid basic auth header"))
	}

	// Get issue date and/or due date
	decoder := json.NewDecoder(r.Body)
	var req CreateInvoiceRequestBody
	if err := decoder.Decode(&req); err != nil {
		return nil, errors.New(errors.InvalidRequest, failure.Messagef("failed to decode request body %v", err))
	}

	// Payment amount should be above 0
	if req.PaymentAmount <= 0 {
		return nil, errors.New(errors.InvalidRequest, failure.Message("payment_amount should be more than 0"))
	}

	// Due date can't be backdated
	if req.DueDate.Before(time.Now()) {
		return nil, errors.New(errors.InvalidRequest, failure.Message("due_date should be in the future"))
	}

	return &service.CreateInvoiceRequest{
		Email:             email,
		Password:          password,
		PaymentAmount:     req.PaymentAmount,
		ServiceProviderId: req.ServiceProviderID,
		DueDate:           req.DueDate,
	}, nil
}
