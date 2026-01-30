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

	api.GET("/home", controllers.HomeIndex)

	auth := api.Group("/")
	auth.Use(middlewares.AuthMiddleware())
	{
		auth.GET("/me", controllers.Me)
		auth.POST("/auth/logout", controllers.Logout)

		/** User */
		auth.GET("/user/name", controllers.GetNameUser)
		auth.GET("/user/:id", controllers.GetUserByID)
		auth.POST("/users", controllers.CreateUser)
		auth.PUT("/user/:id", controllers.UpdateUser)
		auth.DELETE("/user/:id", controllers.DeleteUser)
		auth.GET("/users/xlsx", controllers.ExportExcelUser)
		auth.GET("/users/pdf", controllers.ExportPDFUser)

		/** Material */
		auth.GET("/materials", controllers.MaterialPaginate)
		auth.GET("/material/:id", controllers.GetMaterialByID)
		auth.POST("/materials", controllers.CreateMaterial)
		auth.PUT("/material/:id", controllers.UpdateMaterial)
		auth.DELETE("/material/:id", controllers.DeleteMaterial)
	}
}
