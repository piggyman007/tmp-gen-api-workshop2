package model

type Transfer struct {
	ID           int    `json:"id" gorm:"primaryKey;autoIncrement"`
	SenderID     int    `json:"sender_id"`
	SenderCode   string `json:"sender_code"`
	ReceiverID   int    `json:"receiver_id"`
	ReceiverCode string `json:"receiver_code"`
	Points       int    `json:"points"`
	CreatedAt    string `json:"created_at"`
}
