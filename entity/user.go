package entity

import "time"

type User struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	RoleId    uint      `json:"role_id"`
	Role      Role      `json:"role" db:",prefix:roles."`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
