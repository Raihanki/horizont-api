package services

import (
	"github.com/Raihanki/horizont-api/entity"
	"github.com/Raihanki/horizont-api/repositories"
	"github.com/Raihanki/horizont-api/requests"
	"github.com/Raihanki/horizont-api/resources"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type CategoryService interface {
	GetAllCategories(ctx *fiber.Ctx) ([]resources.CategoryResource, error)
	CreateCategory(ctx *fiber.Ctx, categoryRequest requests.CategoryRequest) (resources.CategoryResource, error)
	UpdateCategory(ctx *fiber.Ctx, categoryRequest requests.CategoryRequest, categoryID uint) (resources.CategoryResource, error)
}

type CategoryServiceImpl struct {
	db                 *sqlx.DB
	categoryRepository repositories.CategoryRepository
}

func NewCategoryService(db *sqlx.DB, categoryRepository repositories.CategoryRepository) CategoryService {
	return &CategoryServiceImpl{db, categoryRepository}
}

func (service *CategoryServiceImpl) GetAllCategories(ctx *fiber.Ctx) ([]resources.CategoryResource, error) {
	tx, _ := service.db.Beginx()
	defer tx.Rollback()
	categories, errGetCategories := service.categoryRepository.GetAll(ctx, tx)
	if errGetCategories != nil {
		return []resources.CategoryResource{}, errGetCategories
	}

	var categoryResources []resources.CategoryResource
	for _, category := range categories {
		categoryResource := resources.CategoryResource{
			ID:   category.ID,
			Name: category.Name,
		}
		categoryResources = append(categoryResources, categoryResource)
	}

	return categoryResources, nil
}

func (service *CategoryServiceImpl) CreateCategory(ctx *fiber.Ctx, categoryRequest requests.CategoryRequest) (resources.CategoryResource, error) {
	tx, _ := service.db.Beginx()
	defer tx.Rollback()

	category, errCreateCategory := service.categoryRepository.Create(ctx, tx, entity.Category{
		Name: categoryRequest.Name,
	})
	if errCreateCategory != nil {
		return resources.CategoryResource{}, errCreateCategory
	}

	categoryResource := resources.CategoryResource{
		ID:   category.ID,
		Name: category.Name,
	}
	return categoryResource, nil
}

func (service *CategoryServiceImpl) UpdateCategory(ctx *fiber.Ctx, categoryRequest requests.CategoryRequest, categoryID uint) (resources.CategoryResource, error) {
	tx, _ := service.db.Beginx()
	defer tx.Rollback()

	category, errGetCategory := service.categoryRepository.GetByID(ctx, tx, categoryID)
	if errGetCategory != nil {
		return resources.CategoryResource{}, errGetCategory
	}

	updatedCategory, errUpdateCategory := service.categoryRepository.Update(ctx, tx, entity.Category{
		Name: categoryRequest.Name,
	}, category.ID)
	if errUpdateCategory != nil {
		return resources.CategoryResource{}, errUpdateCategory
	}

	categoryResource := resources.CategoryResource{
		ID:   updatedCategory.ID,
		Name: updatedCategory.Name,
	}
	return categoryResource, nil
}
