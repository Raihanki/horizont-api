package repositories

import (
	"database/sql"
	"errors"

	log "github.com/sirupsen/logrus"

	"github.com/Raihanki/horizont-api/entity"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type RoleRepository interface {
	GetAll(ctx *fiber.Ctx, tx *sqlx.Tx) ([]entity.Role, error)
	GetByID(ctx *fiber.Ctx, tx *sqlx.Tx, roleID uint) (entity.Role, error)
	Create(ctx *fiber.Ctx, tx *sqlx.Tx, role entity.Role) (entity.Role, error)
	Update(ctx *fiber.Ctx, tx *sqlx.Tx, role entity.Role, categoryID uint) (entity.Role, error)
}

type RoleRepositoryImpl struct {
}

func NewRoleRepository() RoleRepository {
	return &RoleRepositoryImpl{}
}

func (repository *RoleRepositoryImpl) GetAll(ctx *fiber.Ctx, tx *sqlx.Tx) ([]entity.Role, error) {
	var roles []entity.Role
	query := "SELECT * FROM roles"
	errGetRoles := tx.SelectContext(ctx.Context(), &roles, query)
	if errGetRoles != nil {
		log.Error("Error GetAllRoles: ", errGetRoles)
		return []entity.Role{}, errGetRoles
	}

	return roles, nil
}

func (repository *RoleRepositoryImpl) GetByID(ctx *fiber.Ctx, tx *sqlx.Tx, roleID uint) (entity.Role, error) {
	role := entity.Role{}
	query := "SELECT * FROM roles WHERE id = ?"
	errGetRole := tx.GetContext(ctx.Context(), &role, query, roleID)
	if errors.Is(errGetRole, sql.ErrNoRows) {
		return entity.Role{}, errGetRole
	}
	if errGetRole != nil {
		log.Error("Error GetRoleByID: ", errGetRole)
		return entity.Role{}, errGetRole
	}

	return role, nil
}

func (repository *RoleRepositoryImpl) Create(ctx *fiber.Ctx, tx *sqlx.Tx, role entity.Role) (entity.Role, error) {
	query := "INSERT INTO roles (name) VALUES (?)"
	_, errCreateRole := tx.ExecContext(ctx.Context(), query, role.Name)

	newRole := entity.Role{}
	errGetRole := tx.GetContext(ctx.Context(), &newRole, "SELECT * FROM roles WHERE name = ?", role.Name)
	if errGetRole != nil {
		log.Error("Error GetRole: ", errGetRole)
		return entity.Role{}, errGetRole
	}

	if errCreateRole != nil {
		log.Error("Error CreateRole: ", errCreateRole)
		return entity.Role{}, errCreateRole
	}

	_ = tx.Commit()
	return newRole, nil
}

func (repository *RoleRepositoryImpl) Update(ctx *fiber.Ctx, tx *sqlx.Tx, role entity.Role, roleID uint) (entity.Role, error) {
	query := "UPDATE roles SET name = ? WHERE id = ?"
	_, errUpdateRole := tx.ExecContext(ctx.Context(), query, role.Name, roleID)

	updatedRole := entity.Role{}
	errGetRole := tx.GetContext(ctx.Context(), &updatedRole, "SELECT * FROM roles WHERE id = ?", roleID)
	if errGetRole != nil {
		log.Error("Error GetRole: ", errGetRole)
		return entity.Role{}, errGetRole
	}

	if errUpdateRole != nil {
		log.Error("Error UpdateRole: ", errUpdateRole)
		return entity.Role{}, errUpdateRole
	}

	_ = tx.Commit()
	return updatedRole, nil
}
