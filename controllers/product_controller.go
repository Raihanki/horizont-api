package controllers

import (
	"database/sql"
	"errors"

	"github.com/Raihanki/horizont-api/helpers"
	"github.com/Raihanki/horizont-api/requests"
	"github.com/Raihanki/horizont-api/services"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type ProductController interface {
	Index(ctx *fiber.Ctx) error
	Show(ctx *fiber.Ctx) error
	Store(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	UpdateStatus(ctx *fiber.Ctx) error
}

type ProductControllerImpl struct {
	productService services.ProductService
}

func NewProductController(productService services.ProductService) ProductController {
	return &ProductControllerImpl{productService}
}

func (controller *ProductControllerImpl) Index(ctx *fiber.Ctx) error {
	products, errProducts := controller.productService.GetAllProducts(ctx)
	if errProducts != nil {
		return helpers.ApiResponse(ctx, fiber.StatusInternalServerError, "Failed to get products", nil)
	}

	return helpers.ApiResponse(ctx, fiber.StatusOK, "Success get products", products)
}

func (controller *ProductControllerImpl) Show(ctx *fiber.Ctx) error {
	productSlug := ctx.Params("product_slug")
	product, errProduct := controller.productService.GetOneProduct(ctx, productSlug)
	if errors.Is(errProduct, sql.ErrNoRows) {
		return helpers.ApiResponse(ctx, fiber.StatusNotFound, "Product not found", nil)
	}
	if errProduct != nil {
		return helpers.ApiResponse(ctx, fiber.StatusInternalServerError, "Failed to get product", nil)
	}

	return helpers.ApiResponse(ctx, fiber.StatusOK, "Success get product", product)
}

func (controller *ProductControllerImpl) Store(ctx *fiber.Ctx) error {
	productRequest := requests.ProductRequest{}
	errBody := ctx.BodyParser(&productRequest)
	if errBody != nil {
		return helpers.ApiResponse(ctx, fiber.StatusBadRequest, "Failed to process request", nil)
	}

	errValidate := validator.New().Struct(productRequest)
	if errValidate != nil {
		errorRepsonses := []helpers.ValidatrionError{}
		for _, err := range errValidate.(validator.ValidationErrors) {
			errorRepsonses = append(errorRepsonses, helpers.ValidatrionError{
				Key:     err.Field(),
				Tag:     err.Tag(),
				Message: err.Value(),
			})
		}
		return helpers.ApiResponse(ctx, fiber.StatusBadRequest, "Failed to validate request", errorRepsonses)
	}

	product, errProduct := controller.productService.CreateProduct(ctx, productRequest)
	if errors.Is(errProduct, sql.ErrNoRows) {
		return helpers.ApiResponse(ctx, fiber.StatusNotFound, "Category not found", nil)
	}
	if errProduct != nil {
		return helpers.ApiResponse(ctx, fiber.StatusInternalServerError, "Failed to create product", nil)
	}

	return helpers.ApiResponse(ctx, fiber.StatusCreated, "Success create product", product)
}

func (controller *ProductControllerImpl) Update(ctx *fiber.Ctx) error {
	productSlug := ctx.Params("product_slug")

	productRequest := requests.ProductRequest{}
	errBody := ctx.BodyParser(&productRequest)
	if errBody != nil {
		return helpers.ApiResponse(ctx, fiber.StatusBadRequest, "Failed to process request", nil)
	}

	errValidate := validator.New().Struct(productRequest)
	if errValidate != nil {
		errorRepsonses := []helpers.ValidatrionError{}
		for _, err := range errValidate.(validator.ValidationErrors) {
			errorRepsonses = append(errorRepsonses, helpers.ValidatrionError{
				Key:     err.Field(),
				Tag:     err.Tag(),
				Message: err.Value(),
			})
		}
		return helpers.ApiResponse(ctx, fiber.StatusBadRequest, "Failed to validate request", errorRepsonses)
	}

	product, errProduct := controller.productService.UpdateProduct(ctx, productRequest, productSlug)
	if errors.Is(errProduct, errors.New("pr-404")) {
		return helpers.ApiResponse(ctx, fiber.StatusNotFound, "Product not found", nil)
	}
	if errors.Is(errProduct, errors.New("ct-404")) {
		return helpers.ApiResponse(ctx, fiber.StatusNotFound, "Category not found", nil)
	}
	if errProduct != nil {
		return helpers.ApiResponse(ctx, fiber.StatusInternalServerError, "Failed to update product", nil)
	}

	return helpers.ApiResponse(ctx, fiber.StatusOK, "Success update product", product)
}

func (controller *ProductControllerImpl) UpdateStatus(ctx *fiber.Ctx) error {
	productSlug := ctx.Params("product_slug")

	errProduct := controller.productService.UnactivateProduct(ctx, productSlug)
	if errors.Is(errProduct, sql.ErrNoRows) {
		return helpers.ApiResponse(ctx, fiber.StatusNotFound, "Product not found", nil)
	}
	if errProduct != nil {
		return helpers.ApiResponse(ctx, fiber.StatusInternalServerError, "Failed to unactivate product", nil)
	}

	return helpers.ApiResponse(ctx, fiber.StatusOK, "Success unactivate product", nil)
}
