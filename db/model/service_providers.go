package model

type ServiceProvider struct {
	ID             uint          `gorm:"primaryKey" json:"id"`
	CompanyID      uint          `gorm:"not null" json:"company_id"`
	Name           string        `gorm:"not null" json:"name"`
	Representative string        `gorm:"not null" json:"representative"`
	PhoneNumber    string        `gorm:"not null" json:"phone_number"`
	PostalCode     string        `gorm:"not null" json:"postal_code"`
	Address        string        `gorm:"not null" json:"address"`
	BankAccounts   []BankAccount `json:"bank_accounts"`
	Invoices       []Invoice     `gorm:"foreignKey:ServiceProviderID" json:"invoices"`
}
