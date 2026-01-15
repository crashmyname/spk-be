package routes

import (
	"spk-be/internal/controllers"
	"spk-be/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")

	api.GET("/users", controllers.UserIndex)
	api.POST("/auth/login", controllers.Login)

	// Material
	api.GET("/material", controllers.MaterialIndex)
	api.GET("/material/:id", controllers.GetMaterialByID)
	api.POST("/material", controllers.CreateMaterial)
	api.PATCH("/material/:id", controllers.UpdateMaterial)
	api.DELETE("/material/:id", controllers.DeleteMaterial)

	api.GET("/home", controllers.HomeIndex)

	auth := api.Group("/")
	auth.Use(middlewares.AuthMiddleware())
	{
		auth.GET("/me", controllers.Me)
		auth.POST("/auth/logout", controllers.Logout)

		auth.GET("/user/name", controllers.GetNameUser)
		auth.GET("/user/:id", controllers.GetUserByID)
		auth.POST("/users", controllers.CreateUser)
		auth.PUT("/user/:id", controllers.UpdateUser)
		auth.DELETE("/user/:id", controllers.DeleteUser)
	}
}
