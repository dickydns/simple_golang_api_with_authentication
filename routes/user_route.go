package routes

import (
	"simple_golang_api_with_authentication/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoute(r *gin.Engine) {
	authentication := r.Group("auth")
	{
		authentication.POST("/login", controllers.LoginUser)
		authentication.POST("/register", controllers.RegisterUser)
		authentication.POST("/refresh", controllers.RefreshToken)
	}
	// , middleware.JWTAuth()
	user := r.Group("/user")
	{
		user.GET("", controllers.GetUser)
		user.PUT("/:id", controllers.UpdateUser)
		user.DELETE("/:id", controllers.DeleteUser)
	}

}
