package route

import (
	"workshop2/handler"
	"workshop2/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(app *fiber.App, db interface{}) {
	// Set DB and JWT secret for handlers
	gdb := db.(*gorm.DB)
	jwtSecret := []byte("supersecretkey")
	handler.UserService = service.NewUserService(gdb, jwtSecret)
	handler.TransferService = service.NewTransferService(gdb)

	app.Post("/register", handler.RegisterHandler)
	app.Post("/login", handler.LoginHandler)
	app.Get("/me", handler.AuthMiddleware, handler.MeHandler)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Hello World"})
	})
	app.Post("/transfer", handler.AuthMiddleware, handler.TransferPointHandler)
	app.Get("/point-histories", handler.AuthMiddleware, handler.PointHistoriesHandler)
}
