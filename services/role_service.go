package services

import (
	"github.com/Raihanki/horizont-api/entity"
	"github.com/Raihanki/horizont-api/repositories"
	"github.com/Raihanki/horizont-api/requests"
	"github.com/Raihanki/horizont-api/resources"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type RoleService interface {
	GetAllRoles(ctx *fiber.Ctx) ([]resources.RoleResource, error)
	CreateRole(ctx *fiber.Ctx, roleRequest requests.RoleRequest) (resources.RoleResource, error)
	UpdateRole(ctx *fiber.Ctx, roleRequest requests.RoleRequest, roleID uint) (resources.RoleResource, error)
}

type RoleServiceImpl struct {
	db             *sqlx.DB
	roleRepository repositories.RoleRepository
}

func NewRoleService(db *sqlx.DB, roleRepository repositories.RoleRepository) RoleService {
	return &RoleServiceImpl{db, roleRepository}
}

func (service *RoleServiceImpl) GetAllRoles(ctx *fiber.Ctx) ([]resources.RoleResource, error) {
	tx, _ := service.db.Beginx()
	defer tx.Rollback()

	roles, errRoles := service.roleRepository.GetAll(ctx, tx)
	if errRoles != nil {
		return []resources.RoleResource{}, errRoles
	}

	var roleResources []resources.RoleResource
	for _, role := range roles {
		roleResource := resources.RoleResource{
			ID:   role.ID,
			Name: role.Name,
		}
		roleResources = append(roleResources, roleResource)
	}

	return roleResources, nil
}

func (service *RoleServiceImpl) CreateRole(ctx *fiber.Ctx, roleRequest requests.RoleRequest) (resources.RoleResource, error) {
	tx, _ := service.db.Beginx()
	defer tx.Rollback()

	role, errCreateRole := service.roleRepository.Create(ctx, tx, entity.Role{
		Name: roleRequest.Name,
	})
	if errCreateRole != nil {
		return resources.RoleResource{}, errCreateRole
	}

	roleResource := resources.RoleResource{
		ID:   role.ID,
		Name: role.Name,
	}

	return roleResource, nil
}

func (service *RoleServiceImpl) UpdateRole(ctx *fiber.Ctx, roleRequest requests.RoleRequest, roleID uint) (resources.RoleResource, error) {
	tx, _ := service.db.Beginx()
	defer tx.Rollback()

	oldRole, errCheckRole := service.roleRepository.GetByID(ctx, tx, roleID)
	if errCheckRole != nil {
		return resources.RoleResource{}, errCheckRole
	}

	role, errUpdateRole := service.roleRepository.Update(ctx, tx, entity.Role{
		Name: roleRequest.Name,
	}, oldRole.ID)
	if errUpdateRole != nil {
		return resources.RoleResource{}, errUpdateRole
	}

	roleResource := resources.RoleResource{
		ID:   role.ID,
		Name: role.Name,
	}

	return roleResource, nil
}
