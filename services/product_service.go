package services

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/Raihanki/horizont-api/entity"
	"github.com/Raihanki/horizont-api/repositories"
	"github.com/Raihanki/horizont-api/requests"
	"github.com/Raihanki/horizont-api/resources"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type ProductService interface {
	GetAllProducts(ctx *fiber.Ctx) ([]resources.ProductResource, error)
	GetOneProduct(ctx *fiber.Ctx, productSlug string) (resources.ProductResource, error)
	CreateProduct(ctx *fiber.Ctx, productRequest requests.ProductRequest) (resources.ProductResource, error)
	UpdateProduct(ctx *fiber.Ctx, productRequest requests.ProductRequest, productSlug string) (resources.ProductResource, error)
	UnactivateProduct(ctx *fiber.Ctx, prouctSlug string) error
}

type ProductServiceImpl struct {
	db                 *sqlx.DB
	productRepository  repositories.ProductRepository
	categoryRepository repositories.CategoryRepository
}

func NewProductService(db *sqlx.DB, productRepository repositories.ProductRepository, categoryRepository repositories.CategoryRepository) ProductService {
	return &ProductServiceImpl{db, productRepository, categoryRepository}
}

func (service *ProductServiceImpl) GetAllProducts(ctx *fiber.Ctx) ([]resources.ProductResource, error) {
	tx, _ := service.db.Beginx()
	defer tx.Rollback()

	products, errProducts := service.productRepository.GetAll(ctx, tx)
	if errProducts != nil {
		return []resources.ProductResource{}, errProducts
	}

	productResources := []resources.ProductResource{}
	for _, product := range products {
		productResources = append(productResources, resources.ProductResource{
			ID:    product.ID,
			Name:  product.Name,
			Slug:  product.Slug,
			Price: product.Price,
			Category: resources.CategoryResource{
				ID:   product.Category.CategoryID,
				Name: product.Category.CategoryName,
			},
			Description:     product.Description,
			ProductCity:     product.ProductCity,
			ProductProvince: product.ProductProvince,
			Stock:           product.Stock,
			Image:           product.Image,
			IsActive:        product.IsActive,
			CreatedAt:       product.CreatedAt,
		})
	}

	return productResources, nil
}

func (service *ProductServiceImpl) GetOneProduct(ctx *fiber.Ctx, productSlug string) (resources.ProductResource, error) {
	tx, _ := service.db.Beginx()
	defer tx.Rollback()

	product, errProduct := service.productRepository.GetBySlug(ctx, tx, productSlug)
	if errProduct != nil {
		return resources.ProductResource{}, errProduct
	}

	productResource := resources.ProductResource{
		ID:    product.ID,
		Name:  product.Name,
		Slug:  product.Slug,
		Price: product.Price,
		Category: resources.CategoryResource{
			ID:   product.Category.CategoryID,
			Name: product.Category.CategoryName,
		},
		Description:     product.Description,
		ProductCity:     product.ProductCity,
		ProductProvince: product.ProductProvince,
		Stock:           product.Stock,
		Image:           product.Image,
		IsActive:        product.IsActive,
		CreatedAt:       product.CreatedAt,
	}

	return productResource, nil
}

func (service *ProductServiceImpl) CreateProduct(ctx *fiber.Ctx, productRequest requests.ProductRequest) (resources.ProductResource, error) {
	tx, _ := service.db.Beginx()
	defer tx.Rollback()

	category, errCategory := service.categoryRepository.GetByID(ctx, tx, productRequest.CategoryID)
	if errCategory != nil {
		return resources.ProductResource{}, errCategory
	}

	//random number
	productSlug := strings.ReplaceAll(strings.ToLower(productRequest.Name), " ", "-")

	_, errCheckProduct := service.productRepository.GetBySlug(ctx, tx, productSlug)
	if errCheckProduct == nil {
		return resources.ProductResource{}, errors.New("slug found")
	}

	product, errCreateProduct := service.productRepository.Create(ctx, tx, entity.Product{
		Name:            productRequest.Name,
		Slug:            productSlug,
		CategoryID:      category.ID,
		Price:           productRequest.Price,
		Description:     productRequest.Description,
		ProductCity:     productRequest.ProductCity,
		ProductProvince: productRequest.ProductProvince,
		Stock:           productRequest.Stock,
		Image:           productRequest.Image,
		IsActive:        true,
	})
	if errCreateProduct != nil {
		return resources.ProductResource{}, errCreateProduct
	}

	productResource := resources.ProductResource{
		ID:    product.ID,
		Name:  product.Name,
		Slug:  product.Slug,
		Price: product.Price,
		Category: resources.CategoryResource{
			ID:   product.Category.CategoryID,
			Name: product.Category.CategoryName,
		},
		Description:     product.Description,
		ProductCity:     product.ProductCity,
		ProductProvince: product.ProductProvince,
		Stock:           product.Stock,
		Image:           product.Image,
		IsActive:        product.IsActive,
		CreatedAt:       product.CreatedAt,
	}

	return productResource, nil
}

func (service *ProductServiceImpl) UpdateProduct(ctx *fiber.Ctx, productRequest requests.ProductRequest, productSlug string) (resources.ProductResource, error) {
	tx, _ := service.db.Beginx()
	defer tx.Rollback()

	_, errGetOldProduct := service.productRepository.GetBySlug(ctx, tx, productSlug)
	if errors.Is(errGetOldProduct, sql.ErrNoRows) {
		return resources.ProductResource{}, errors.New("pr-404")
	}
	if errGetOldProduct != nil {
		return resources.ProductResource{}, errGetOldProduct
	}

	category, errCategory := service.categoryRepository.GetByID(ctx, tx, productRequest.CategoryID)
	if errors.Is(errCategory, sql.ErrNoRows) {
		return resources.ProductResource{}, errors.New("ct-404")
	}
	if errCategory != nil {
		return resources.ProductResource{}, errCategory
	}

	//random number
	slug := strings.ReplaceAll(strings.ToLower(productRequest.Name), " ", "-")
	product, errCreateProduct := service.productRepository.Update(ctx, tx, entity.Product{
		Name:            productRequest.Name,
		Slug:            slug,
		CategoryID:      category.ID,
		Price:           productRequest.Price,
		Description:     productRequest.Description,
		ProductCity:     productRequest.ProductCity,
		ProductProvince: productRequest.ProductProvince,
		Stock:           productRequest.Stock,
		Image:           productRequest.Image,
		IsActive:        true,
	}, productSlug)
	if errCreateProduct != nil {
		return resources.ProductResource{}, errCreateProduct
	}

	productResource := resources.ProductResource{
		ID:    product.ID,
		Name:  product.Name,
		Slug:  product.Slug,
		Price: product.Price,
		Category: resources.CategoryResource{
			ID:   product.Category.CategoryID,
			Name: product.Category.CategoryName,
		},
		Description:     product.Description,
		ProductCity:     product.ProductCity,
		ProductProvince: product.ProductProvince,
		Stock:           product.Stock,
		Image:           product.Image,
		IsActive:        product.IsActive,
		CreatedAt:       product.CreatedAt,
	}

	return productResource, nil
}

func (service *ProductServiceImpl) UnactivateProduct(ctx *fiber.Ctx, productSlug string) error {
	tx, _ := service.db.Beginx()
	defer tx.Rollback()

	product, errGetProduct := service.productRepository.GetBySlug(ctx, tx, productSlug)
	if errGetProduct != nil {
		return errGetProduct
	}

	active := !product.IsActive
	errUnactivateProduct := service.productRepository.Unactivate(ctx, tx, product.Slug, active)
	if errUnactivateProduct != nil {
		return errUnactivateProduct
	}

	return nil
}
