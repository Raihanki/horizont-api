package requests

type CategoryRequest struct {
	Name string `json:"name" form:"name" validate:"required,max=255"`
}
