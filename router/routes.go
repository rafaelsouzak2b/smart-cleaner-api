package router

import (
	"github.com/gin-gonic/gin"
	"github.com/guicazaroto/learning-go/config"
	"github.com/guicazaroto/learning-go/handler"
)

func initializeRoutes(router *gin.Engine) {
	handler.InitializeHandler()

	v1 := router.Group("/api/v1")
	{
		v1.GET("/ping", handler.PingHandler)
		v1.GET("/cleaner/search", config.AuthTokenMiddleware(), handler.GetCleanerHandler)
		v1.GET("/cleaner/:id", config.AuthTokenMiddleware(), handler.GetCleanerByIdHandler)
		v1.GET("/cleaner/me", config.JWTMiddleware("cleaner"), handler.GetCleanerMeByIdHandler)
		v1.POST("/cleaner", config.AuthTokenMiddleware(), handler.CreateCleanerHandler)
		v1.PUT("/cleaner", config.JWTMiddleware("cleaner"), handler.UpdateCleanerHandler)
		v1.DELETE("/cleaner", config.JWTMiddleware("cleaner"), handler.DeleteCleanerHandler)
		v1.POST("/cleaner/login", config.AuthTokenMiddleware(), handler.LoginCleanerHandler)
		v1.POST("/cleaner/img", config.JWTMiddleware("cleaner"), handler.SendImgProfileHandler)
		v1.PATCH("/cleaner/img", config.JWTMiddleware("cleaner"), handler.UpdateImgProfileHandler)

	}
}
