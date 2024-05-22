package entity

import (
	"time"

	"github.com/shopspring/decimal"
)

type Product struct {
	ID              uint            `json:"id" db:"id"`
	Name            string          `json:"name" db:"name"`
	Slug            string          `json:"slug" db:"slug"`
	CategoryID      uint            `json:"category_id" db:"category_id"`
	Category        ProductCategory `json:"category" db:"category"`
	Price           decimal.Decimal `json:"price" db:"price"`
	Description     string          `json:"description" db:"description"`
	ProductCity     string          `json:"product_city" db:"product_city"`
	ProductProvince string          `json:"product_province" db:"product_province"`
	Stock           uint            `json:"stock" db:"stock"`
	Image           string          `json:"image" db:"image"`
	IsActive        bool            `json:"is_active" db:"is_active"`
	CreatedAt       time.Time       `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at" db:"updated_at"`
}

type ProductCategory struct {
	CategoryID   uint   `db:"category_id_2"`
	CategoryName string `db:"category_name"`
}
