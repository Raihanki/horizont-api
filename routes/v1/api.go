package v1

import "github.com/gofiber/fiber/v2"

func Routes(Api *fiber.App) {
	app := Api.Group("/api/v1")

	// user routes
	userRoutes := app.Group("/users")
	userRoutes.Get("/")
	userRoutes.Post("/login")
	userRoutes.Post("/register")

	// category routes
	categoryRoutes := app.Group("/categories")
	categoryRoutes.Get("/")

	// product routes
	productRoutes := app.Group("/products")
	productRoutes.Get("/")
	productRoutes.Get("/:product_id")

	// transaction routes
	transactionRoutes := app.Group("/transactions")
	transactionRoutes.Get("/")
	transactionRoutes.Get("/:transaction_id")
	transactionRoutes.Post("/")

	//-----------------------------------

	// admin routes
	adminRoutes := app.Group("/admin")
	// category
	adminRoutes.Post("/categories/")
	adminRoutes.Patch("/categories/:category_id")
	// product
	adminRoutes.Post("/products")
	adminRoutes.Put("/products/:product_id")
	adminRoutes.Patch("/products/:product_id/update-status")
	// transaction
	adminRoutes.Get("/transactions")
	adminRoutes.Get("/transactions/:transaction_id")
	adminRoutes.Patch("/transactions/:transaction_id/update-shipment-number")
}
