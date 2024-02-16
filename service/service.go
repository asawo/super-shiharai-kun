package service

import (
	"context"

	"github.com/asawo/api/config"
	"github.com/asawo/api/db"
	"github.com/asawo/api/db/model"
	"github.com/asawo/api/errors"
	"github.com/asawo/api/logger"
)

type Service interface {
	CreateInvoice(ctx context.Context, req *CreateInvoiceRequest) (*model.Invoice, errors.ServiceError)
	ListInvoices(ctx context.Context, req *ListInvoicesRequest) ([]*model.Invoice, errors.ServiceError)
}

type Impl struct {
	logger *logger.Log
	env    *config.Env
	db     db.DB
}

func NewService(logger *logger.Log, env *config.Env, db db.DB) *Impl {
	return &Impl{
		logger: logger,
		env:    env,
		db:     db,
	}
}
