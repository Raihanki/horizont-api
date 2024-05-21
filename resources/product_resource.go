package resources

import (
	"time"

	"github.com/shopspring/decimal"
)

type ProductResource struct {
	ID              uint             `json:"id"`
	Name            string           `json:"name"`
	Slug            string           `json:"slug"`
	Category        CategoryResource `json:"category"`
	Price           decimal.Decimal  `json:"price"`
	Description     string           `json:"description"`
	ProductCity     string           `json:"product_city"`
	ProductProvince string           `json:"product_province"`
	Stock           uint             `json:"stock"`
	Image           string           `json:"image"`
	IsActive        bool             `json:"is_active"`
	CreatedAt       time.Time        `json:"created_at"`
}
