package db

//go:generate mockgen -source db.go -destination db_mock.go -package db

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/asawo/api/config"
	"github.com/asawo/api/db/model"
	"github.com/asawo/api/errors"
)

type DB interface {
	// CreateInvoice creates an invoice
	CreateInvoice(tx *gorm.DB, invoice *model.Invoice) errors.ServiceError
	// ListInvoices returns a list of invoices
	ListInvoices(tx *gorm.DB, companyId uint, start, time time.Time) ([]*model.Invoice, errors.ServiceError)
	// GetUserByEmail returns a user with an email
	GetUserByEmail(tx *gorm.DB, email string) (*model.User, errors.ServiceError)
	// ListServiceProvidersByCompanyID returns a list of service providers belonging to a company
	ListServiceProvidersByCompanyID(tx *gorm.DB, companyId uint) ([]*model.ServiceProvider, errors.ServiceError)
	// ListBankAccountsByServiceProviderID returns a list of bank accounts associated with a service provider
	ListBankAccountsByServiceProviderID(tx *gorm.DB, serviceProviderId uint) ([]*model.BankAccount, errors.ServiceError)
}

type db struct {
	client *gorm.DB
}

const retrySleepPeriod = 5 * time.Second

// InitDB initialize the DB connector
func InitDB(env *config.Env) (*gorm.DB, error) {
	db, err := connectDB(env, 10)
	if err != nil {
		return nil, err
	}
	if err := migrate(db); err != nil {
		return nil, err
	}

	return db, nil
}

// connectDB creates a DB connection
func connectDB(env *config.Env, retryCount int) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo",
		env.DBHost,
		env.DBUser,
		env.DBPassword,
		env.DBName,
		env.DBPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		if retryCount > 0 {
			time.Sleep(retrySleepPeriod)
			return connectDB(env, retryCount-1)
		}
		return nil, fmt.Errorf("failed to connect database")
	}

	return db, nil
}

// run the migration
func migrate(db *gorm.DB) error {
	if err := db.Migrator().AutoMigrate(
		&model.Company{},
		&model.ServiceProvider{},
		&model.BankAccount{},
		&model.Invoice{},
		&model.User{},
	); err != nil {
		return err
	}

	return nil
}

// New creates a new DB instance
func New(env *config.Env) (DB, error) {
	client, err := InitDB(env)
	if err != nil {
		return nil, err
	}

	return &db{
		client: client,
	}, nil
}

func (d *db) getClient(tx *gorm.DB) *gorm.DB {
	client := d.client
	if tx != nil {
		client = tx
	}
	return client
}
