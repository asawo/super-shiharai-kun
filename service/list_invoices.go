package service

import (
	"context"
	"time"

	"github.com/asawo/api/db/model"
	"github.com/asawo/api/errors"
)

type ListInvoicesRequest struct {
	Email     string
	Password  string
	StartDate time.Time
	EndDate   time.Time
}

func (s *Impl) ListInvoices(ctx context.Context, req *ListInvoicesRequest) ([]*model.Invoice, errors.ServiceError) {
	// Validate and authenticate user
	user, serr := s.validateUser(req.Email, req.Password)
	if serr != nil {
		return nil, serr
	}

	// Get Invoices sorted by due date, between start date and end date
	invoices, serr := s.db.ListInvoices(nil, user.CompanyID, req.StartDate, req.EndDate)
	if serr != nil {
		return nil, serr
	}

	return invoices, nil
}
