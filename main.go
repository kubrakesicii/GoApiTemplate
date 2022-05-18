package main

import (
	//"goapitemplate/pkg/middleware"
	"fmt"
	"goapitemplate/pkg/middleware"
	"goapitemplate/pkg/route"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.New()
	app.Use(middleware.ResponseReturn())

	router := app.Group("/")
	route.LoadRoutes(router)

	app.GET("/test", func(c *gin.Context) {
		fmt.Println("İstek geldi")
		//return helper.Response{StatusCode: 400,Message: "DENEME",Success: true,Data: nil,Count: 0}
	})
	app.Run(":5005")
}
