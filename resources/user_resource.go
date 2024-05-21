package resources

type UserResource struct {
	ID        uint         `json:"id"`
	Name      string       `json:"name"`
	Email     string       `json:"email"`
	Role      RoleResource `json:"role"`
	CreatedAt string       `json:"created_at"`
}
