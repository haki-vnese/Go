package routes

import (
	"fmt"
	"go-rest-api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		api.POST("/login", func(c *gin.Context) {
			fmt.Println("Login route hit")
			controllers.Login(c)
		})
		api.POST("/employees", controllers.CreateEmployee)
		api.GET("/employees", controllers.GetEmployees)
		api.PUT("/employees/:id", controllers.UpdateEmployee)
		api.DELETE("/employees/:id", controllers.DeleteEmployee)
	}

	return r
}
