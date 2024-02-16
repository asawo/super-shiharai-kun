package service

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/morikuni/failure"

	"github.com/asawo/api/db/model"
	"github.com/asawo/api/errors"
)

const (
	COMMISION_RATE float64 = 0.04
	TAX_RATE       float64 = 0.10
)

type CreateInvoiceRequest struct {
	Email             string
	Password          string
	ServiceProviderId uint
	DueDate           time.Time
	PaymentAmount     float64
}

func (s *Impl) CreateInvoice(ctx context.Context, req *CreateInvoiceRequest) (*model.Invoice, errors.ServiceError) {
	// Validate and authenticate user
	user, serr := s.validateUser(req.Email, req.Password)
	if serr != nil {
		return nil, serr
	}

	// Validate that the invoice belongs to one of the service providers associated to the user's company
	serviceProviders, serr := s.db.ListServiceProvidersByCompanyID(nil, user.CompanyID)
	if serr != nil {
		return nil, serr
	}
	if !containsServiceProvider(serviceProviders, req.ServiceProviderId) {
		return nil, errors.New(errors.PermissionDenied, failure.Messagef("service provider id %d is not registered as a service provider with your company", req.ServiceProviderId))
	}

	// Validate the ServiceProvider has at least one bank account
	bankAccounts, serr := s.db.ListBankAccountsByServiceProviderID(nil, req.ServiceProviderId)
	if serr != nil {
		return nil, serr
	}
	if len(bankAccounts) == 0 {
		return nil, errors.New(errors.NotFound, failure.Messagef("failed to find bank account associated with service provider id %d", req.ServiceProviderId))
	}

	// Calculate commission, tax, and total payment amount
	commission := req.PaymentAmount * COMMISION_RATE
	tax := commission * TAX_RATE
	total := req.PaymentAmount + (tax + commission)

	invoice := model.Invoice{
		ID:                uuid.New(),
		IssueDate:         time.Now(),
		DueDate:           req.DueDate,
		ServiceProviderID: req.ServiceProviderId,
		CompanyID:         user.CompanyID,
		Commission:        roundFloat(commission, 2),
		CommissionRate:    COMMISION_RATE,
		TaxRate:           TAX_RATE,
		Tax:               roundFloat(tax, 2),
		Amount:            roundFloat(total, 2),
		PaymentAmount:     req.PaymentAmount,
		Status:            model.Outstanding,
	}

	if serr := s.db.CreateInvoice(nil, &invoice); serr != nil {
		return nil, serr
	}

	return &invoice, nil
}

func containsServiceProvider(serviceProviders []*model.ServiceProvider, id uint) bool {
	for _, sp := range serviceProviders {
		if sp.ID == id {
			return true
		}
	}

	return false
}
