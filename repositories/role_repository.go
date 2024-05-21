package repositories

import (
	"github.com/Raihanki/horizont-api/entity"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type RoleRepository interface {
	GetAll() ([]entity.Role, error)
	GetByID(roleID uint) (entity.Role, error)
	Create(role entity.Role) (entity.Role, error)
	Update(role entity.Role, categoryID uint) (entity.Role, error)
}

type RoleRepositoryImpl struct {
	tx  sqlx.Tx
	ctx *fiber.Ctx
}

func NewRoleRepository(tx sqlx.Tx, ctx *fiber.Ctx) RoleRepository {
	return &RoleRepositoryImpl{tx, ctx}
}

func (repository *RoleRepositoryImpl) GetAll() ([]entity.Role, error) {
	var roles []entity.Role
	query := "SELECT * FROM roles"
	errGetRoles := repository.tx.SelectContext(repository.ctx.Context(), &roles, query)
	if errGetRoles != nil {
		return []entity.Role{}, errGetRoles
	}

	return roles, nil
}

func (repository *RoleRepositoryImpl) GetByID(roleID uint) (entity.Role, error) {
	role := entity.Role{}
	query := "SELECT * FROM roles WHERE id = ?"
	errGetRole := repository.tx.GetContext(repository.ctx.Context(), &role, query, roleID)
	if errGetRole != nil {
		return entity.Role{}, errGetRole
	}

	return role, nil
}

func (repository *RoleRepositoryImpl) Create(role entity.Role) (entity.Role, error) {
	query := "INSERT INTO roles (name) VALUES (?)"
	_, errCreateRole := repository.tx.ExecContext(repository.ctx.Context(), query, role.Name)
	if errCreateRole != nil {
		return entity.Role{}, errCreateRole
	}

	return role, nil
}

func (repository *RoleRepositoryImpl) Update(role entity.Role, roleID uint) (entity.Role, error) {
	query := "UPDATE roles SET name = ? WHERE id = ?"
	_, errUpdateRole := repository.tx.ExecContext(repository.ctx.Context(), query, role.Name, roleID)
	if errUpdateRole != nil {
		return entity.Role{}, errUpdateRole
	}

	return role, nil
}
