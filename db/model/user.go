package model

type User struct {
	CompanyID uint   `gorm:"not null" json:"company_id"`
	Name      string `gorm:"not null" json:"name"`
	Email     string `gorm:"unique;not null" json:"email"`
	Password  string `gorm:"not null" json:"password"` // Password is a hash of the password
}
