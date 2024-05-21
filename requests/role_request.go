package requests

type RoleRequest struct {
	Name string `json:"name" form:"name" validate:"required,max=255"`
}
