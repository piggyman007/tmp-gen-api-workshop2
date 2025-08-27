package handler

import (
	"workshop2/model"
	"workshop2/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var TransferService *service.TransferService

type TransferRequest struct {
	ReceiverCode string `json:"receiver_code"`
	Points       int    `json:"points"`
}

func TransferPointHandler(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(int)
	var req TransferRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	serviceReq := service.TransferRequest{
		ReceiverCode: req.ReceiverCode,
		Points:       req.Points,
	}
	err := TransferService.TransferPoint(userID, serviceReq)
	if err == gorm.ErrRecordNotFound {
		return c.Status(404).JSON(fiber.Map{"error": "Receiver not found"})
	}
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Transfer failed"})
	}
	return c.JSON(fiber.Map{"success": true})
}

func PointHistoriesHandler(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(int)
	transfers := TransferService.GetPointHistories(userID, 10)
	var histories []fiber.Map
	for _, t := range transfers {
		var sender model.User
		var receiver model.User
		TransferService.DB.First(&sender, t.SenderID)
		TransferService.DB.First(&receiver, t.ReceiverID)
		histories = append(histories, fiber.Map{
			"from":        sender.FirstName + " " + sender.LastName,
			"to":          receiver.FirstName + " " + receiver.LastName,
			"points":      t.Points,
			"date":        t.CreatedAt,
			"sender_code": t.SenderCode,
		})
	}
	return c.JSON(histories)
}
