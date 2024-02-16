package model

type BankAccount struct {
	ID                uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	ServiceProviderID uint   `gorm:"not null" json:"service_provider_id"`
	BankName          string `gorm:"not null" json:"bank_name"`
	BranchName        string `gorm:"not null" json:"branch_name"`
	AccountNumber     string `gorm:"not null" json:"account_number"`
	AccountName       string `gorm:"not null" json:"account_name"`
}
