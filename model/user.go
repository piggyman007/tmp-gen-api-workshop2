package model

type User struct {
	ID        int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Email     string `json:"email" gorm:"uniqueIndex"`
	Password  string `json:"password"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Phone     string `json:"phone_number"`
	Birthday  string `json:"birthday"`
}
