package entity

type User struct {
	Id       uint   `gorm:"AUTO_INCREMENT;PRIMARY_KEY;not null"`
	Email    string `gorm:"not null"`
	Password []byte
}
