package v1

import (
	"github.com/Raihanki/horizont-api/controllers"
	"github.com/Raihanki/horizont-api/repositories"
	"github.com/Raihanki/horizont-api/services"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func Routes(Api *fiber.App, db *sqlx.DB) {
	// CONTROLLER INITIALIZE
	roleController := controllers.NewRoleController(services.NewRoleService(db, repositories.NewRoleRepository()))
	categoryController := controllers.NewCategoryController(services.NewCategoryService(db, repositories.NewCategoryRepository()))
	productController := controllers.NewProductController(services.NewProductService(db, repositories.NewProductRepository(), repositories.NewCategoryRepository()))

	Api.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(200).JSON(fiber.Map{
			"message": "Welcome to Horizont API",
		})
	})

	app := Api.Group("/api/v1")

	// user routes
	// userRoutes := app.Group("/users")
	// userRoutes.Get("/")
	// userRoutes.Post("/login")
	// userRoutes.Post("/register")

	// category routes
	categoryRoutes := app.Group("/categories")
	categoryRoutes.Get("/", categoryController.Index)
	categoryRoutes.Post("/", categoryController.Store)
	categoryRoutes.Patch("/:category_id", categoryController.Update)

	// product routes
	productRoutes := app.Group("/products")
	productRoutes.Get("/", productController.Index)
	productRoutes.Put("/:product_slug", productController.Update)
	productRoutes.Post("/", productController.Store)
	productRoutes.Get("/:product_slug", productController.Show)
	productRoutes.Patch("/:product_slug/update-status", productController.UpdateStatus)

	// transaction routes
	// transactionRoutes := app.Group("/transactions")
	// transactionRoutes.Get("/")
	// transactionRoutes.Get("/:transaction_id")
	// transactionRoutes.Post("/")

	// role
	roleRoutes := app.Group("/roles")
	roleRoutes.Get("/", roleController.Index)
	roleRoutes.Post("/", roleController.Store)
	roleRoutes.Patch("/:role_id", roleController.Update)

	//-----------------------------------
	// transaction
	// adminRoutes.Get("/transactions")
	// adminRoutes.Get("/transactions/:transaction_id")
	// adminRoutes.Patch("/transactions/:transaction_id/update-shipment-number")
}
