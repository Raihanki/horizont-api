package controllers

import (
	"database/sql"
	"errors"
	"strconv"

	"github.com/Raihanki/horizont-api/helpers"
	"github.com/Raihanki/horizont-api/requests"
	"github.com/Raihanki/horizont-api/services"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type RoleController interface {
	Index(ctx *fiber.Ctx) error
	Store(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
}

type RoleControllerImpl struct {
	roleService services.RoleService
}

func NewRoleController(roleService services.RoleService) RoleController {
	return &RoleControllerImpl{roleService}
}

func (controller *RoleControllerImpl) Index(ctx *fiber.Ctx) error {
	roles, errRoles := controller.roleService.GetAllRoles(ctx)
	if errRoles != nil {
		return helpers.ApiResponse(ctx, fiber.StatusInternalServerError, "Failed to get roles", nil)
	}

	return helpers.ApiResponse(ctx, fiber.StatusOK, "Success get roles", roles)
}

func (controller *RoleControllerImpl) Store(ctx *fiber.Ctx) error {
	roleRequest := requests.RoleRequest{}
	errBody := ctx.BodyParser(&roleRequest)
	if errBody != nil {
		return helpers.ApiResponse(ctx, fiber.StatusBadRequest, "Failed to parse request body", nil)
	}

	errValidate := validator.New().Struct(roleRequest)
	if errValidate != nil {
		errorResponse := []helpers.ValidatrionError{}
		for _, err := range errValidate.(validator.ValidationErrors) {
			errorResponse = append(errorResponse, helpers.ValidatrionError{
				Key:     err.Field(),
				Tag:     err.Tag(),
				Message: err.Value(),
			})
		}
		return helpers.ApiResponse(ctx, fiber.StatusUnprocessableEntity, "Failed to validate request", errorResponse)
	}

	roleResource, errCreateRole := controller.roleService.CreateRole(ctx, roleRequest)
	if errCreateRole != nil {
		return helpers.ApiResponse(ctx, fiber.StatusInternalServerError, "Failed to create role", nil)
	}

	return helpers.ApiResponse(ctx, fiber.StatusCreated, "Success create role", roleResource)
}

func (controller *RoleControllerImpl) Update(ctx *fiber.Ctx) error {
	roleIdParam := ctx.Params("role_id")
	roleID, _ := strconv.Atoi(roleIdParam)

	roleRequest := requests.RoleRequest{}
	errBody := ctx.BodyParser(&roleRequest)
	if errBody != nil {
		return helpers.ApiResponse(ctx, fiber.StatusBadRequest, "Failed to parse request body", nil)
	}

	errValidate := validator.New().Struct(roleRequest)
	if errValidate != nil {
		errorResponse := []helpers.ValidatrionError{}
		for _, err := range errValidate.(validator.ValidationErrors) {
			errorResponse = append(errorResponse, helpers.ValidatrionError{
				Key:     err.Field(),
				Tag:     err.Tag(),
				Message: err.Value(),
			})
		}
		return helpers.ApiResponse(ctx, fiber.StatusUnprocessableEntity, "Failed to validate request", errorResponse)
	}

	roleResource, errUpdateRole := controller.roleService.UpdateRole(ctx, roleRequest, uint(roleID))
	if errors.Is(errUpdateRole, sql.ErrNoRows) {
		return helpers.ApiResponse(ctx, fiber.StatusNotFound, "Role Not Found", nil)
	}
	if errUpdateRole != nil {
		return helpers.ApiResponse(ctx, fiber.StatusInternalServerError, "Failed to update role", nil)
	}

	return helpers.ApiResponse(ctx, fiber.StatusOK, "Success update role", roleResource)
}
