package controller

import "goapitemplate/app/service"

type UserController struct {
	userService service.IUserService
}

func (controller *UserController) GetUser() {

}
