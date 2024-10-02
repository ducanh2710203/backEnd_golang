package routes

import (
	"go-backend/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	r.POST("/login", controllers.LoginUser)
	r.POST("/register", controllers.RegisterUser)
}
