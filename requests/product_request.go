package requests

import "github.com/shopspring/decimal"

type ProductRequest struct {
	Name            string          `json:"name" form:"name" validate:"required,max=255"`
	CategoryID      uint            `json:"category_id" form:"category_id" validate:"required,numeric"`
	Price           decimal.Decimal `json:"price" form:"price" validate:"required"`
	Description     string          `json:"description" form:"description" validate:"required"`
	ProductCity     string          `json:"product_city" form:"product_city" validate:"required"`
	ProductProvince string          `json:"product_province" form:"product_province" validate:"required"`
	Stock           uint            `json:"stock" form:"stock" validate:"required"`
	Image           string          `json:"image" form:"image" validate:"required"`
	IsActive        bool            `json:"is_active" form:"is_active"`
}
