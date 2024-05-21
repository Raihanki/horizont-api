package entity

import (
	"time"

	"github.com/shopspring/decimal"
)

type Product struct {
	ID              uint            `json:"id"`
	Name            string          `json:"name"`
	Slug            string          `json:"slug"`
	CategoryID      uint            `json:"category_id"`
	Category        Category        `json:"category" db:",prefix=categories."`
	Price           decimal.Decimal `json:"price"`
	Description     string          `json:"description"`
	ProductCity     string          `json:"product_city"`
	ProductProvince string          `json:"product_province"`
	Stock           uint            `json:"stock"`
	Image           string          `json:"image"`
	IsActive        bool            `json:"is_active"`
	CreatedAt       time.Time       `json:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at"`
}
