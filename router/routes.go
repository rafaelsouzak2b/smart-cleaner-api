package router

import (
	"github.com/gin-gonic/gin"
	"github.com/guicazaroto/learning-go/handler"
)

func initializeRoutes (router *gin.Engine) {
	handler.InitializeHandler()

	v1 := router.Group("/api/v1")
	{
		v1.GET("/ping", handler.PingHandler)
		
		v1.GET("/cleaners", handler.GetCleanerHandler)

		v1.GET("/cleaners/:id", handler.GetCleanerByIdHandler)

		v1.POST("/cleaners", handler.CreateCleanerHandler)

		v1.PUT("/cleaners/:id", handler.UpdateCleanerHandler)

		v1.DELETE("/cleaners/:id", handler.DeleteCleanerHandler)
	}
}