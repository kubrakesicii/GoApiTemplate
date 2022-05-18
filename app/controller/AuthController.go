package controller

import (
	"goapitemplate/app/service"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	AuthService service.IAuthService
}

func (controller *AuthController) Login(c *gin.Context) {

}
