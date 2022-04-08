package main

import (
	"goapitemplate/pkg/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.New()
	app.Use(middlewares.Logger())
	app.Run(":5005")
}
