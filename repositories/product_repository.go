package repositories

import (
	"github.com/Raihanki/horizont-api/entity"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type ProductRepository interface {
	GetAll() ([]entity.Product, error)
	GetBySlug(slug string) (entity.Product, error)
	Create(product entity.Product) (entity.Product, error)
	Update(product entity.Product, productID uint) (entity.Product, error)
}

type ProductRepositoryImpl struct {
	tx  *sqlx.Tx
	ctx *fiber.Ctx
}

func NewProductRepository(tx *sqlx.Tx, ctx *fiber.Ctx) ProductRepository {
	return &ProductRepositoryImpl{tx, ctx}
}

func (repository *ProductRepositoryImpl) GetAll() ([]entity.Product, error) {
	var products []entity.Product
	query := "SELECT * FROM products INNER JOIN categories ON products.category_id = categories.id"
	errGetProducts := repository.tx.SelectContext(repository.ctx.Context(), &products, query)

	if errGetProducts != nil {
		return []entity.Product{}, errGetProducts
	}

	return products, nil
}

func (repository *ProductRepositoryImpl) GetBySlug(slug string) (entity.Product, error) {
	product := entity.Product{}
	query := "SELECT * FROM products INNER JOIN categories ON products.category_id = categories.id WHERE products.slug = ?"
	errGetProduct := repository.tx.GetContext(repository.ctx.Context(), &product, query, slug)

	if errGetProduct != nil {
		return entity.Product{}, errGetProduct
	}

	return product, nil
}

func (repository *ProductRepositoryImpl) Create(product entity.Product) (entity.Product, error) {
	query := "INSERT INTO products (name, slug, price, description, product_city, product_province, stock, image, is_active, category_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	_, errCreateProduct := repository.tx.ExecContext(repository.ctx.Context(), query, product.Name, product.Slug, product.Price, product.Description, product.ProductCity, product.ProductProvince, product.Stock, product.Image, product.IsActive, product.CategoryID)
	if errCreateProduct != nil {
		return entity.Product{}, errCreateProduct
	}

	return product, nil
}

func (repository *ProductRepositoryImpl) Update(product entity.Product, productID uint) (entity.Product, error) {
	query := "UPDATE products SET name = ?, slug = ?, price = ?, description = ?, product_city = ?, product_province = ?, stock = ?, image = ?, is_active = ?, category_id = ? WHERE id = ?"
	_, errUpdateProduct := repository.tx.ExecContext(repository.ctx.Context(), query, product.Name, product.Slug, product.Price, product.Description, product.ProductCity, product.ProductProvince, product.Stock, product.Image, product.IsActive, product.CategoryID, productID)
	if errUpdateProduct != nil {
		return entity.Product{}, errUpdateProduct
	}

	return product, nil
}
