// TransferService now uses db.TransferDBInterface for DB operations
// Easily swap between real DB and mock for testing
// Last update: 2025-08-27
// Author: alongkorn.c

package service

import (
	"time"
	"workshop2/db"
	"workshop2/model"

	"gorm.io/gorm"
)

type TransferService struct {
	DB      *gorm.DB
	DBLayer db.TransferDB
}

type TransferRequest struct {
	ReceiverCode string
	Points       int
}

func NewTransferService(dbConn *gorm.DB) *TransferService {
	return &TransferService{
		DB:      dbConn,
		DBLayer: db.NewTransferDB(dbConn),
	}
}

func (s *TransferService) TransferPoint(userID int, req TransferRequest) error {
	receiver, err := s.DBLayer.FindUserByCode(req.ReceiverCode)
	if err != nil {
		return err
	}
	sender, err := s.DBLayer.FindUserByID(userID)
	if err != nil {
		return err
	}
	transfer := model.Transfer{
		SenderID:     userID,
		ReceiverID:   receiver.ID,
		SenderCode:   sender.Email,
		ReceiverCode: req.ReceiverCode,
		Points:       req.Points,
		CreatedAt:    time.Now().Format("2006-01-02"),
	}
	if err := s.DBLayer.CreateTransfer(&transfer); err != nil {
		return err
	}
	return nil
}

func (s *TransferService) GetPointHistories(userID int, limit int) []model.Transfer {
	return s.DBLayer.FindTransfersByUser(userID, limit)
}
