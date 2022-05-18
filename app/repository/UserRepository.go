package repository

import (
	"gorm.io/gorm"
)

type IUserRepository interface {
}

type UserRepository struct {
	//repository IUserRepository
	db *gorm.DB
}
