package service

import (
	"goapitemplate/app/repository"
)

type IUserService interface {
}

type UserService struct {
	userRepo repository.IUserRepository
}
