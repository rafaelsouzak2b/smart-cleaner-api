package router

import (
	"github.com/gin-gonic/gin"
	"github.com/guicazaroto/learning-go/handler"
)

func initializeRoutes(router *gin.Engine) {
	handler.InitializeHandler()

	v1 := router.Group("/api/v1")
	{
		v1.GET("/ping", handler.PingHandler)

		v1.GET("/cleaner", handler.GetCleanerHandler)

		v1.GET("/cleaner/:id", handler.GetCleanerByIdHandler)

		v1.POST("/cleaner", handler.CreateCleanerHandler)

		v1.PUT("/cleaner/:id", handler.UpdateCleanerHandler)

		// v1.DELETE("/cleaners/:id", handler.DeleteCleanerHandler)
		v1.GET("/user", handler.GetUserHandler)
		v1.POST("/user", handler.CreateUserHandler)
	}
}
