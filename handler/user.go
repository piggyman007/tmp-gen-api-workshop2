package handler

import (
	"workshop2/model"
	"workshop2/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var DB *gorm.DB

func RegisterHandler(c *fiber.Ctx) error {
	var req model.User
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	err := service.RegisterUser(DB, &req)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(true)
}

func LoginHandler(c *fiber.Ctx) error {
	type LoginReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var req LoginReq
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	token, err := service.LoginUser(DB, req.Email, req.Password)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"token": token})
}

func MeHandler(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(int)
	user, err := service.GetUserByID(DB, userID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(user)
}
