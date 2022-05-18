package route

import (
	"goapitemplate/app/controller"
	"goapitemplate/app/repository"
	"goapitemplate/app/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthRoutes(route *gin.RouterGroup, db *gorm.DB) {
	authRepository := &repository.AuthRepository{DB: db}
	authService := &service.AuthService{AuthRepository: authRepository}
	authController := controller.AuthController{AuthService: authService}

	route.POST("/login", authController.Login)
}
