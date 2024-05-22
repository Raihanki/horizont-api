package repositories

import (
	"database/sql"
	"errors"

	log "github.com/sirupsen/logrus"

	"github.com/Raihanki/horizont-api/entity"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type ProductRepository interface {
	GetAll(ctx *fiber.Ctx, tx *sqlx.Tx) ([]entity.Product, error)
	GetBySlug(ctx *fiber.Ctx, tx *sqlx.Tx, productSlug string) (entity.Product, error)
	Create(ctx *fiber.Ctx, tx *sqlx.Tx, product entity.Product) (entity.Product, error)
	Update(ctx *fiber.Ctx, tx *sqlx.Tx, product entity.Product, productSlug string) (entity.Product, error)
	Unactivate(ctx *fiber.Ctx, tx *sqlx.Tx, productSlug string, active bool) error
}

type ProductRepositoryImpl struct {
}

func NewProductRepository() ProductRepository {
	return &ProductRepositoryImpl{}
}

func (repository *ProductRepositoryImpl) GetAll(ctx *fiber.Ctx, tx *sqlx.Tx) ([]entity.Product, error) {
	var newProduct []struct {
		entity.Product
		entity.ProductCategory
	}
	query := "SELECT products.*, categories.name as category_name, categories.id as category_id_2 FROM products INNER JOIN categories ON products.category_id = categories.id WHERE products.is_active = true"
	errGetProducts := tx.SelectContext(ctx.Context(), &newProduct, query)
	if errGetProducts != nil {
		log.Error("Error GetAllProducts: ", errGetProducts)
		return []entity.Product{}, errGetProducts
	}

	products := []entity.Product{}
	for _, product := range newProduct {
		products = append(products, entity.Product{
			ID:              product.ID,
			Name:            product.Name,
			Slug:            product.Slug,
			Price:           product.Price,
			Description:     product.Description,
			ProductCity:     product.ProductCity,
			ProductProvince: product.ProductProvince,
			Stock:           product.Stock,
			Image:           product.Image,
			IsActive:        product.IsActive,
			Category: entity.ProductCategory{
				CategoryID:   product.ProductCategory.CategoryID,
				CategoryName: product.ProductCategory.CategoryName,
			},
			CreatedAt: product.CreatedAt,
		})
	}

	return products, nil
}

func (repository *ProductRepositoryImpl) GetBySlug(ctx *fiber.Ctx, tx *sqlx.Tx, slug string) (entity.Product, error) {
	var tempProduct struct {
		entity.Product
		entity.ProductCategory
	}
	query := "SELECT products.*, categories.id as category_id_2, categories.name as category_name FROM products INNER JOIN categories ON products.category_id = categories.id WHERE slug = ? AND products.is_active = true"
	errGetProduct := tx.GetContext(ctx.Context(), &tempProduct, query, slug)
	if errors.Is(errGetProduct, sql.ErrNoRows) {
		return entity.Product{}, errGetProduct
	}
	if errGetProduct != nil {
		log.Error("Error GetProduct: ", errGetProduct)
		return entity.Product{}, errGetProduct
	}

	product := entity.Product{
		ID:              tempProduct.ID,
		Name:            tempProduct.Name,
		Slug:            tempProduct.Slug,
		Price:           tempProduct.Price,
		Description:     tempProduct.Description,
		ProductCity:     tempProduct.ProductCity,
		ProductProvince: tempProduct.ProductProvince,
		Stock:           tempProduct.Stock,
		Image:           tempProduct.Image,
		IsActive:        tempProduct.IsActive,
		Category: entity.ProductCategory{
			CategoryID:   tempProduct.ProductCategory.CategoryID,
			CategoryName: tempProduct.ProductCategory.CategoryName,
		},
		CreatedAt: tempProduct.CreatedAt,
	}

	return product, nil
}

func (repository *ProductRepositoryImpl) Create(ctx *fiber.Ctx, tx *sqlx.Tx, product entity.Product) (entity.Product, error) {
	query := "INSERT INTO products (name, slug, price, description, product_city, product_province, stock, image, is_active, category_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	_, errCreateProduct := tx.ExecContext(ctx.Context(), query, product.Name, product.Slug, product.Price, product.Description, product.ProductCity, product.ProductProvince, product.Stock, product.Image, product.IsActive, product.CategoryID)
	if errCreateProduct != nil {
		log.Error("Error CreateProduct: ", errCreateProduct)
		return entity.Product{}, errCreateProduct
	}

	var tempProduct struct {
		entity.Product
		entity.ProductCategory
	}
	errNewProduct := tx.GetContext(ctx.Context(), &tempProduct, "SELECT products.*, categories.id as category_id_2, categories.name as category_name FROM products INNER JOIN categories ON products.category_id = categories.id WHERE slug = ? AND products.is_active = true", product.Slug)
	if errNewProduct != nil {
		log.Error("Error CreateProduct: ", errNewProduct)
		return entity.Product{}, errNewProduct
	}

	newProduct := entity.Product{
		ID:              tempProduct.ID,
		Name:            tempProduct.Name,
		Slug:            tempProduct.Slug,
		Price:           tempProduct.Price,
		Description:     tempProduct.Description,
		ProductCity:     tempProduct.ProductCity,
		ProductProvince: tempProduct.ProductProvince,
		Stock:           tempProduct.Stock,
		Image:           tempProduct.Image,
		IsActive:        tempProduct.IsActive,
		Category: entity.ProductCategory{
			CategoryID:   tempProduct.ProductCategory.CategoryID,
			CategoryName: tempProduct.ProductCategory.CategoryName,
		},
		CreatedAt: tempProduct.CreatedAt,
	}

	_ = tx.Commit()
	return newProduct, nil
}

func (repository *ProductRepositoryImpl) Update(ctx *fiber.Ctx, tx *sqlx.Tx, product entity.Product, productSlug string) (entity.Product, error) {
	query := "UPDATE products SET name = ?, slug = ?, price = ?, description = ?, product_city = ?, product_province = ?, stock = ?, image = ?, is_active = ?, category_id = ? WHERE slug = ? AND is_active = true"
	_, errUpdateProduct := tx.ExecContext(ctx.Context(), query, product.Name, product.Slug, product.Price, product.Description, product.ProductCity, product.ProductProvince, product.Stock, product.Image, product.IsActive, product.CategoryID, productSlug)
	if errUpdateProduct != nil {
		log.Error("Error UpdateProduct: ", errUpdateProduct)
		return entity.Product{}, errUpdateProduct
	}

	var updatedProductTemp struct {
		entity.Product
		entity.ProductCategory
	}
	errNewProduct := tx.GetContext(ctx.Context(), &updatedProductTemp, "SELECT products.*, categories.id as category_id_2, categories.name as category_name FROM products INNER JOIN categories ON products.category_id = categories.id WHERE slug = ? AND products.is_active = true", product.Slug)
	if errNewProduct != nil {
		log.Error("Error GetProduct: ", errNewProduct)
		return entity.Product{}, errNewProduct
	}

	updatedProduct := entity.Product{
		ID:              updatedProductTemp.ID,
		Name:            updatedProductTemp.Name,
		Slug:            updatedProductTemp.Slug,
		Price:           updatedProductTemp.Price,
		Description:     updatedProductTemp.Description,
		ProductCity:     updatedProductTemp.ProductCity,
		ProductProvince: updatedProductTemp.ProductProvince,
		Stock:           updatedProductTemp.Stock,
		Image:           updatedProductTemp.Image,
		IsActive:        updatedProductTemp.IsActive,
		Category: entity.ProductCategory{
			CategoryID:   updatedProductTemp.ProductCategory.CategoryID,
			CategoryName: updatedProductTemp.ProductCategory.CategoryName,
		},
		CreatedAt: updatedProductTemp.CreatedAt,
	}

	_ = tx.Commit()
	return updatedProduct, nil
}

func (repository *ProductRepositoryImpl) Unactivate(ctx *fiber.Ctx, tx *sqlx.Tx, productSlug string, active bool) error {
	query := "UPDATE products SET is_active = ? WHERE slug = ?"
	_, errUnactivateProduct := tx.ExecContext(ctx.Context(), query, active, productSlug)
	if errUnactivateProduct != nil {
		log.Error("Error UnactivateProduct: ", errUnactivateProduct)
		return errUnactivateProduct
	}

	_ = tx.Commit()
	return nil
}
