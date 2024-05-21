package entity

import (
	"time"

	"github.com/shopspring/decimal"
)

type Transaction struct {
	ID                uint            `json:"id"`
	UserID            uint            `json:"user_id"`
	User              User            `json:"user"`
	ProductID         uint            `json:"product_id"`
	Product           Product         `json:"product"`
	Quantity          uint            `json:"quantity"`
	TotalProductPrice decimal.Decimal `json:"total_product_price"`
	ShippingPrice     decimal.Decimal `json:"shipping_price"`
	TotalPrice        decimal.Decimal `json:"total_price"`
	Address           string          `json:"address"`
	Status            string          `json:"status"`
	ShipmentNumber    *string         `json:"shipment_number"`
	CreatedAt         time.Time       `json:"created_at"`
	UpdatedAt         time.Time       `json:"updated_at"`
}
