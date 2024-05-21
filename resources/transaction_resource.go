package resources

import (
	"time"

	"github.com/shopspring/decimal"
)

type TransactionResource struct {
	ID                uint            `json:"id"`
	User              UserResource    `json:"user"`
	Product           ProductResource `json:"product"`
	Quantity          uint            `json:"quantity"`
	TotalProductPrice decimal.Decimal `json:"total_product_price"`
	ShippingPrice     decimal.Decimal `json:"shipping_price"`
	TotalPrice        decimal.Decimal `json:"total_price"`
	Address           string          `json:"address"`
	Status            string          `json:"status"`
	ShipmentNumber    *string         `json:"shipment_number"`
	CreatedAt         time.Time       `json:"created_at"`
}
