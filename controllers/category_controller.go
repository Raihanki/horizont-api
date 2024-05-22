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

type CategoryController interface {
	Index(ctx *fiber.Ctx) error
	Store(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
}

type CategoryControllerImpl struct {
	categoryService services.CategoryService
}

func NewCategoryController(categoryService services.CategoryService) CategoryController {
	return &CategoryControllerImpl{categoryService}
}

func (controller *CategoryControllerImpl) Index(ctx *fiber.Ctx) error {
	categories, errCreateCategory := controller.categoryService.GetAllCategories(ctx)
	if errCreateCategory != nil {
		return helpers.ApiResponse(ctx, fiber.StatusInternalServerError, "Failed to get categories", nil)
	}

	return helpers.ApiResponse(ctx, fiber.StatusOK, "Success get categories", categories)
}

func (controller *CategoryControllerImpl) Store(ctx *fiber.Ctx) error {
	categoryRequest := requests.CategoryRequest{}
	errBody := ctx.BodyParser(&categoryRequest)
	if errBody != nil {
		return helpers.ApiResponse(ctx, fiber.StatusBadRequest, "Failed to parse request body", nil)
	}

	errValidate := validator.New().Struct(categoryRequest)
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

	category, errCreateCategory := controller.categoryService.CreateCategory(ctx, categoryRequest)
	if errCreateCategory != nil {
		return helpers.ApiResponse(ctx, fiber.StatusInternalServerError, "Failed to create category", nil)
	}

	return helpers.ApiResponse(ctx, fiber.StatusCreated, "Success create category", category)
}

func (controller *CategoryControllerImpl) Update(ctx *fiber.Ctx) error {
	categoryIdParam := ctx.Params("category_id")
	categoryID, _ := strconv.Atoi(categoryIdParam)

	categoryRequest := requests.CategoryRequest{}
	errBody := ctx.BodyParser(&categoryRequest)
	if errBody != nil {
		return helpers.ApiResponse(ctx, fiber.StatusBadRequest, "Failed to parse request body", nil)
	}

	errValidate := validator.New().Struct(categoryRequest)
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

	categories, errCreateCategory := controller.categoryService.UpdateCategory(ctx, categoryRequest, uint(categoryID))
	if errors.Is(errCreateCategory, sql.ErrNoRows) {
		return helpers.ApiResponse(ctx, fiber.StatusNotFound, "Category not found", nil)
	}
	if errCreateCategory != nil {
		return helpers.ApiResponse(ctx, fiber.StatusInternalServerError, "Failed to update category", nil)
	}

	return helpers.ApiResponse(ctx, fiber.StatusCreated, "Success update category", categories)
}
