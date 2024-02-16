package model

type Company struct {
	ID               uint              `gorm:"primaryKey" json:"id"`
	Name             string            `gorm:"not null" json:"name"`
	Representative   string            `gorm:"not null" json:"representative"`
	PhoneNumber      string            `gorm:"not null" json:"phone_number"`
	PostalCode       string            `gorm:"not null" json:"postal_code"`
	Address          string            `gorm:"not null" json:"address"`
	Users            []User            `gorm:"foreignKey:CompanyID" json:"users"`
	ServiceProviders []ServiceProvider `gorm:"foreignKey:CompanyID" json:"service_providers"`
	Invoices         []Invoice         `gorm:"foreignKey:CompanyID" json:"invoices"`
}
