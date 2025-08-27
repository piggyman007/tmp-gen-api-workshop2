package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var JwtSecret []byte

func AuthMiddleware(c *fiber.Ctx) error {
	auth := c.Get("Authorization")
	if len(auth) < 8 || auth[:7] != "Bearer " {
		return c.Status(401).JSON(fiber.Map{"error": "Missing or invalid token"})
	}
	tokenStr := auth[7:]
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return JwtSecret, nil
	})
	if err != nil || !token.Valid {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid token"})
	}
	claims := token.Claims.(jwt.MapClaims)
	c.Locals("user_id", int(claims["user_id"].(float64)))
	return c.Next()
}
