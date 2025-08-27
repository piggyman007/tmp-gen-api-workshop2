package route

import (
	"workshop2/handler"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(app *fiber.App, db interface{}) {
	// Set DB and JWT secret for handlers
	gdb := db.(*gorm.DB)
	handler.DB = gdb
	handler.TransferDB = gdb
	handler.JwtSecret = []byte("supersecretkey")

	app.Post("/register", handler.RegisterHandler)
	app.Post("/login", handler.LoginHandler)
	app.Get("/me", handler.AuthMiddleware, handler.MeHandler)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Hello World"})
	})
	app.Post("/transfer", handler.AuthMiddleware, handler.TransferPointHandler)
	app.Get("/point-histories", handler.AuthMiddleware, handler.PointHistoriesHandler)
}
