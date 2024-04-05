package routes

import (
	"gin-todoapp/src/controllers"
	"github.com/gin-gonic/gin"
)

func SetupTodosRouter(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		v1.GET("todo", controllers.GetTodos)
		v1.POST("todo", controllers.CreateATodo)
		v1.GET("todo/:id", controllers.GetATodo)
		v1.PUT("todo/:id", controllers.UpdateATodo)
		v1.DELETE("todo/:id", controllers.DeleteATodo)
	}
}

func SetupUserRoutes(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		v1.GET("user", controllers.GetUsers)
		v1.POST("user", controllers.CreateAnUser)
		v1.GET("user/:id", controllers.GetAUser)
		v1.PUT("user/:id", controllers.UpdateAUser)
		v1.DELETE("user/:id", controllers.DeleteAUser)
	}
}
