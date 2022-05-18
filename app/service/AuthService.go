package service

import (
	"goapitemplate/app/repository"
)

type IAuthService interface {
}

type AuthService struct {
	AuthRepository repository.IAuthRepository
}
