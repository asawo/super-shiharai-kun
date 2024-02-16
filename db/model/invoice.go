package model

import (
	"time"

	"github.com/google/uuid"
)

type Invoice struct {
	ID                uuid.UUID `gorm:"primaryKey;autoIncrement" json:"id"`
	IssueDate         time.Time `json:"issue_date"`
	PaymentAmount     float64   `json:"payment_amount"`
	Commission        float64   `json:"commission"`
	CommissionRate    float64   `json:"commission_rate"`
	Tax               float64   `json:"tax"`
	TaxRate           float64   `json:"tax_rate"`
	Amount            float64   `json:"amount"`
	DueDate           time.Time `json:"due_date"`
	CompanyID         uint      `json:"company_id"`
	ServiceProviderID uint      `json:"service_provider_id"`
	Status            Status    `json:"status"`
}

type Status string

const (
	// 未処理
	Outstanding Status = "OUTSTANDING"
	// 処理中
	Processing Status = "PROCESSING"
	// 支払済
	Paid Status = "PAID"
	// エラー
	Error Status = "ERROR"
)
