package routes

import (
	"spk-be/internal/controllers"
	"spk-be/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")

	api.GET("/users", controllers.UserIndex)
	api.GET("/user/:id", controllers.GetUserByID)
	api.POST("/auth/login", controllers.Login)
	api.POST("/users", controllers.CreateUser)

	// Material
	api.GET("/material", controllers.MaterialIndex)
	api.GET("/material/:id", controllers.GetMaterialByID)
	api.POST("/material", controllers.CreateMaterial)
	api.PATCH("/material/:id", controllers.UpdateMaterial)
	api.DELETE("/material/:id", controllers.DeleteMaterial)

	auth := api.Group("/")
	auth.Use(middlewares.AuhtMiddleware())
	{
		auth.GET("/user/name", controllers.GetNameUser)
	}
}
