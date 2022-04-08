package route

import (
	"fmt"
	"goapitemplate/platform/database"

	"github.com/gin-gonic/gin"
)

func LoadRoutes(superRoute *gin.RouterGroup) {
	db := database.GetDb()

	fmt.Println(db)
	/// Route Groups
	//UserRoutes(superRoute, db)
}
