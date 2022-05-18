package entity

import "goapitemplate/app/generic"

type User struct {
	generic.IEntity
	Id       uint   `gorm:"AUTO_INCREMENT;PRIMARY_KEY;not null"`
	Email    string `gorm:"not null"`
	Password []byte
	Deneme   string
}
