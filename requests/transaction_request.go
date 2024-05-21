package requests

type TransactionRequest struct {
	ProductID uint   `json:"product_id" form:"product_id" validate:"required"`
	Quantity  uint   `json:"quantity" form:"quantity" validate:"required"`
	Address   string `json:"address" form:"address" validate:"required"`
}
