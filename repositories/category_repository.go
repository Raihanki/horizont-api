package repositories

import (
	"database/sql"
	"errors"

	"github.com/Raihanki/horizont-api/entity"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

type CategoryRepository interface {
	GetAll(ctx *fiber.Ctx, tx *sqlx.Tx) ([]entity.Category, error)
	GetByID(ctx *fiber.Ctx, tx *sqlx.Tx, categoryID uint) (entity.Category, error)
	Create(ctx *fiber.Ctx, tx *sqlx.Tx, category entity.Category) (entity.Category, error)
	Update(ctx *fiber.Ctx, tx *sqlx.Tx, category entity.Category, categoryID uint) (entity.Category, error)
}

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (repository *CategoryRepositoryImpl) GetAll(ctx *fiber.Ctx, tx *sqlx.Tx) ([]entity.Category, error) {
	var categories []entity.Category
	query := "SELECT * FROM categories"
	errGetCategories := tx.SelectContext(ctx.Context(), &categories, query)
	if errGetCategories != nil {
		log.Error("Error GetAllCategories: ", errGetCategories)
		return []entity.Category{}, errGetCategories
	}

	return categories, nil
}

func (repository *CategoryRepositoryImpl) GetByID(ctx *fiber.Ctx, tx *sqlx.Tx, categoryID uint) (entity.Category, error) {
	category := entity.Category{}
	query := "SELECT * FROM categories WHERE id = ?"
	errGetCategory := tx.GetContext(ctx.Context(), &category, query, categoryID)
	if errors.Is(errGetCategory, sql.ErrNoRows) {
		return entity.Category{}, errGetCategory
	}
	if errGetCategory != nil {
		log.Error("Error GetCategoryByID: ", errGetCategory)
		return entity.Category{}, errGetCategory
	}

	return category, nil
}

func (repository *CategoryRepositoryImpl) Create(ctx *fiber.Ctx, tx *sqlx.Tx, category entity.Category) (entity.Category, error) {
	query := "INSERT INTO categories (name) VALUES (?)"
	_, errCreateCategory := tx.ExecContext(ctx.Context(), query, category.Name)
	if errCreateCategory != nil {
		log.Error("Error CreateCategory: ", errCreateCategory)
		return entity.Category{}, errCreateCategory
	}

	newCategory := entity.Category{}
	errGetCategory := tx.GetContext(ctx.Context(), &newCategory, "SELECT * FROM categories WHERE name = ?", category.Name)
	if errGetCategory != nil {
		log.Error("Error GetCategory: ", errGetCategory)
		return entity.Category{}, errGetCategory
	}

	_ = tx.Commit()
	return newCategory, nil
}

func (repository *CategoryRepositoryImpl) Update(ctx *fiber.Ctx, tx *sqlx.Tx, category entity.Category, categoryID uint) (entity.Category, error) {
	query := "UPDATE categories SET name = ? WHERE id = ?"
	_, errUpdateCategory := tx.ExecContext(ctx.Context(), query, category.Name, categoryID)
	if errUpdateCategory != nil {
		log.Error("Error UpdateCategory: ", errUpdateCategory)
		return entity.Category{}, errUpdateCategory
	}

	updatedCategory := entity.Category{}
	errGetCategory := tx.GetContext(ctx.Context(), &updatedCategory, "SELECT * FROM categories WHERE id = ?", categoryID)
	if errGetCategory != nil {
		log.Error("Error GetCategory: ", errGetCategory)
		return entity.Category{}, errGetCategory
	}

	_ = tx.Commit()
	return updatedCategory, nil
}
