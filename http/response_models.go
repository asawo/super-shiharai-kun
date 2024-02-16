package http

import "github.com/asawo/api/db/model"

type ListInvoicesResponse struct {
	Result string           `json:"result"`
	Data   []*model.Invoice `json:"data"`
}

type CreateInvoiceResponse struct {
	Result string         `json:"result"`
	Data   *model.Invoice `json:"data"`
}
