package repositories

import (
	"github.com/Raihanki/horizont-api/entity"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type CategoryRepository interface {
	GetAll() ([]entity.Category, error)
	GetByID(categoryID uint) (entity.Category, error)
	Create(category entity.Category) (entity.Category, error)
	Update(category entity.Category, categoryID uint) (entity.Category, error)
}

type CategoryRepositoryImpl struct {
	tx  *sqlx.Tx
	ctx *fiber.Ctx
}

func NewCategoryRepository(tx *sqlx.Tx, ctx *fiber.Ctx) CategoryRepository {
	return &CategoryRepositoryImpl{tx, ctx}
}

func (repository *CategoryRepositoryImpl) GetAll() ([]entity.Category, error) {
	var categories []entity.Category
	query := "SELECT * FROM categories"
	errGetCategories := repository.tx.SelectContext(repository.ctx.Context(), &categories, query)
	if errGetCategories != nil {
		return []entity.Category{}, errGetCategories
	}

	return categories, nil
}

func (repository *CategoryRepositoryImpl) GetByID(categoryID uint) (entity.Category, error) {
	category := entity.Category{}
	query := "SELECT * FROM categories WHERE id = ?"
	errGetCategory := repository.tx.GetContext(repository.ctx.Context(), &category, query, categoryID)
	if errGetCategory != nil {
		return entity.Category{}, errGetCategory
	}

	return category, nil
}

func (repository *CategoryRepositoryImpl) Create(category entity.Category) (entity.Category, error) {
	query := "INSERT INTO categories (name) VALUES (?)"
	_, errCreateCategory := repository.tx.ExecContext(repository.ctx.Context(), query, category.Name)
	if errCreateCategory != nil {
		return entity.Category{}, errCreateCategory
	}

	return category, nil
}

func (repository *CategoryRepositoryImpl) Update(category entity.Category, categoryID uint) (entity.Category, error) {
	query := "UPDATE categories SET name = ? WHERE id = ?"
	_, errUpdateCategory := repository.tx.ExecContext(repository.ctx.Context(), query, category.Name, categoryID)
	if errUpdateCategory != nil {
		return entity.Category{}, errUpdateCategory
	}

	return category, nil
}
