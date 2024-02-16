package seed

import (
	"fmt"
	"time"

	"github.com/asawo/api/auth"
	"github.com/asawo/api/db/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// This function is purely for local testing purporses.
// addTestData can be used for entering test data via gorm.
// However, it's recommended to just use the make seed-db command.
func AddTestData(db *gorm.DB) {
	// Create sample data
	company := model.Company{
		ID:             1,
		Name:           "Example Corporation",
		Representative: "John Doe",
		PhoneNumber:    "123456789",
		PostalCode:     "12345",
		Address:        "123 Main Street",
	}
	db.Create(&company)

	company2 := model.Company{
		ID:             2,
		Name:           "Example Corporation 3",
		Representative: "Ben Doe",
		PhoneNumber:    "123456789",
		PostalCode:     "12345",
		Address:        "123 Main Street",
	}
	db.Create(&company2)

	user := model.User{
		Name:      "Alice",
		Email:     "alice@example.com",
		CompanyID: company.ID,
	}
	pass, err := auth.HashPassword("password123")
	if err != nil {
		fmt.Print(err)
	}

	user.Password = pass

	db.Create(&user)

	user2 := model.User{
		Name:      "Arthur",
		Email:     "arthur@example.com",
		CompanyID: company2.ID,
	}

	pass2, err := auth.HashPassword("password123")
	if err != nil {
		fmt.Print(err)
	}
	user2.Password = pass2

	db.Create(&user2)

	serviceProvider := model.ServiceProvider{
		CompanyID:      company.ID,
		Name:           "serviceProvider Corporation",
		Representative: "Jane Smith",
		PhoneNumber:    "987654321",
		PostalCode:     "54321",
		Address:        "456 Oak Avenue",
	}
	db.Create(&serviceProvider)

	serviceProvider2 := model.ServiceProvider{
		CompanyID:      company2.ID,
		Name:           "serviceProvider Corporation 2",
		Representative: "Jane Smith",
		PhoneNumber:    "987654321",
		PostalCode:     "54321",
		Address:        "456 Oak Avenue",
	}
	db.Create(&serviceProvider2)

	serviceProvider3 := model.ServiceProvider{
		CompanyID:      company2.ID,
		Name:           "serviceProvider Corporation 3",
		Representative: "Jane Smith",
		PhoneNumber:    "987654321",
		PostalCode:     "54321",
		Address:        "456 Oak Avenue",
	}
	db.Create(&serviceProvider3)

	bankAccount := model.BankAccount{
		ServiceProviderID: serviceProvider.ID,
		BankName:          "Bank of Example",
		BranchName:        "Main Branch",
		AccountNumber:     "1234567890",
		AccountName:       "Jane Smith",
	}
	db.Create(&bankAccount)

	bankAccount2 := model.BankAccount{
		ServiceProviderID: serviceProvider2.ID,
		BankName:          "Bank of Example 2",
		BranchName:        "Main Branch",
		AccountNumber:     "1234567890",
		AccountName:       "Jane Smith",
	}
	db.Create(&bankAccount2)

	bankAccount3 := model.BankAccount{
		ServiceProviderID: serviceProvider3.ID,
		BankName:          "Bank of Example",
		BranchName:        "Main Branch",
		AccountNumber:     "1234567890",
		AccountName:       "Jane Smith",
	}
	db.Create(&bankAccount3)

	invoice := model.Invoice{
		ID:                uuid.New(),
		IssueDate:         time.Now(),
		PaymentAmount:     100.00,
		Commission:        4.00,
		CommissionRate:    0.04,
		Tax:               1.10,
		TaxRate:           0.10,
		Amount:            104.40,
		DueDate:           time.Now().AddDate(0, 0, 30),
		Status:            model.Outstanding,
		CompanyID:         company.ID,
		ServiceProviderID: serviceProvider.ID,
	}
	db.Create(&invoice)

	invoice2 := model.Invoice{
		ID:                uuid.New(),
		IssueDate:         time.Now(),
		PaymentAmount:     100.00,
		Commission:        4.00,
		CommissionRate:    0.04,
		Tax:               1.10,
		TaxRate:           0.10,
		Amount:            104.40,
		DueDate:           time.Now().AddDate(0, 0, 30),
		Status:            model.Outstanding,
		CompanyID:         company2.ID,
		ServiceProviderID: serviceProvider2.ID,
	}

	db.Create(&invoice2)

	fmt.Println("Sample data created successfully!")
}
