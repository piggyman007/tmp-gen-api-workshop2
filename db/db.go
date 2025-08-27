package db

import (
	"workshop2/model"

	"gorm.io/gorm"
)

type TransferDB struct {
	DB *gorm.DB
}

func NewTransferDB(db *gorm.DB) TransferDBInterface {
	return &TransferDB{DB: db}
}

func (t *TransferDB) FindUserByCode(code interface{}) (model.User, error) {
	var user model.User
	t.DB.Where("email = ? OR id = ?", code, code).First(&user)
	if user.ID == 0 {
		return user, gorm.ErrRecordNotFound
	}
	return user, nil
}

func (t *TransferDB) FindUserByID(id int) (model.User, error) {
	var user model.User
	t.DB.First(&user, id)
	if user.ID == 0 {
		return user, gorm.ErrRecordNotFound
	}
	return user, nil
}

func (t *TransferDB) CreateTransfer(transfer *model.Transfer) error {
	return t.DB.Create(transfer).Error
}

func (t *TransferDB) FindTransfersByUser(userID int, limit int) []model.Transfer {
	var transfers []model.Transfer
	t.DB.Where("sender_id = ? OR receiver_id = ?", userID, userID).Order("created_at desc").Limit(limit).Find(&transfers)
	return transfers
}

type TransferDBInterface interface {
	FindUserByCode(code interface{}) (model.User, error)
	FindUserByID(id int) (model.User, error)
	CreateTransfer(transfer *model.Transfer) error
	FindTransfersByUser(userID int, limit int) []model.Transfer
}

var _ TransferDBInterface = (*TransferDB)(nil)
