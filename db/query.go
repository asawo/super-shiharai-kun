package db

import (
	"time"

	"github.com/asawo/api/db/model"
	"github.com/asawo/api/errors"
	"gorm.io/gorm"
)

func (d *db) GetUserByEmail(tx *gorm.DB, email string) (*model.User, errors.ServiceError) {
	client := d.getClient(tx)

	var user model.User
	if err := client.
		Where("email = ?", email).
		First(&user).
		Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NewFromError(err, errors.NotFound)
		}
		return nil, errors.NewFromError(err, errors.Internal)

	}

	return &user, nil
}

func (d *db) CreateInvoice(tx *gorm.DB, invoice *model.Invoice) errors.ServiceError {
	client := d.getClient(tx)
	if err := client.Create(invoice).Error; err != nil {
		return errors.NewFromError(err, errors.Internal)
	}
	return nil
}

func (d *db) ListInvoices(tx *gorm.DB, companyId uint, start, end time.Time) ([]*model.Invoice, errors.ServiceError) {
	client := d.getClient(tx)

	var invoices []*model.Invoice
	query := client.Where("company_id = ?", companyId)
	if !start.IsZero() {
		query = query.Where("due_date >= ?", start)
	}
	if !end.IsZero() {
		query = query.Where("due_date <= ?", end)
	}
	query = query.Where("status != ?", model.Paid)
	if err := query.Order("due_date").Find(&invoices).Error; err != nil {
		return nil, errors.NewFromError(err, errors.Internal)
	}

	return invoices, nil
}

func (d *db) ListServiceProvidersByCompanyID(tx *gorm.DB, companyId uint) ([]*model.ServiceProvider, errors.ServiceError) {
	client := d.getClient(tx)

	var serviceProviders []*model.ServiceProvider
	query := client.Where("company_id = ?", companyId)

	if err := query.Find(&serviceProviders).Error; err != nil {
		return nil, errors.NewFromError(err, errors.Internal)
	}

	return serviceProviders, nil
}

func (d *db) ListBankAccountsByServiceProviderID(tx *gorm.DB, serviceProviderId uint) ([]*model.BankAccount, errors.ServiceError) {
	client := d.getClient(tx)

	var bankAccounts []*model.BankAccount
	if err := client.
		Where("service_provider_id = ?", serviceProviderId).
		Find(&bankAccounts).
		Error; err != nil {
		return nil, errors.NewFromError(err, errors.Internal)
	}

	return bankAccounts, nil
}
