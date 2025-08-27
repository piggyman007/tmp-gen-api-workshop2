package handler

import (
	"time"
	"workshop2/model"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var TransferDB *gorm.DB

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
	var receiver model.User
	TransferDB.Where("email = ? OR id = ?", req.ReceiverCode, req.ReceiverCode).First(&receiver)
	if receiver.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "Receiver not found"})
	}
	transfer := model.Transfer{
		SenderID:     userID,
		ReceiverID:   receiver.ID,
		ReceiverCode: req.ReceiverCode,
		Points:       req.Points,
		CreatedAt:    time.Now().Format("2006-01-02"),
	}
	if err := TransferDB.Create(&transfer).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Transfer failed"})
	}
	return c.JSON(fiber.Map{"success": true})
}

func PointHistoriesHandler(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(int)
	var transfers []model.Transfer
	TransferDB.Where("sender_id = ? OR receiver_id = ?", userID, userID).Order("created_at desc").Limit(10).Find(&transfers)
	var histories []fiber.Map
	for _, t := range transfers {
		var sender model.User
		var receiver model.User
		TransferDB.First(&sender, t.SenderID)
		TransferDB.First(&receiver, t.ReceiverID)
		histories = append(histories, fiber.Map{
			"from":   sender.FirstName + " " + sender.LastName,
			"to":     receiver.FirstName + " " + receiver.LastName,
			"points": t.Points,
			"date":   t.CreatedAt,
		})
	}
	return c.JSON(histories)
}
