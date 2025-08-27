package main

import (
	"workshop2/handler"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(app *fiber.App, db interface{}) {
	// Set DB and JWT secret for handlers
	handler.DB = db.(*gorm.DB)
	handler.JwtSecret = []byte("supersecretkey")

	app.Post("/register", handler.RegisterHandler)
	app.Post("/login", handler.LoginHandler)
	app.Get("/me", handler.AuthMiddleware, handler.MeHandler)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Hello World"})
	})
}
