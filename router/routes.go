package router

import (
	"github.com/gin-gonic/gin"
	"github.com/guicazaroto/learning-go/config"
	"github.com/guicazaroto/learning-go/handler"
	"github.com/guicazaroto/learning-go/repository"
)

func initializeRoutes(router *gin.Engine) {
	handler.InitializeHandler()

	cleanerRepository := repository.NewCleanerRepository(config.GetDb())
	messageRepository := repository.NewMessageRepository(config.GetDb())

	v1 := router.Group("/api/v1")
	{
		v1.GET("/ping", handler.PingHandler)
		v1.GET("/cleaner/search", config.AuthTokenMiddleware(), handler.GetCleanerHandler(cleanerRepository))
		v1.GET("/cleaner/:id", config.AuthTokenMiddleware(), handler.GetCleanerByIdHandler(cleanerRepository))
		v1.GET("/cleaner/me", config.JWTMiddleware("cleaner"), handler.GetCleanerMeByIdHandler(cleanerRepository))
		v1.POST("/cleaner", config.AuthTokenMiddleware(), handler.CreateCleanerHandler(cleanerRepository))
		v1.PUT("/cleaner", config.JWTMiddleware("cleaner"), handler.UpdateCleanerHandler(cleanerRepository))
		v1.DELETE("/cleaner", config.JWTMiddleware("cleaner"), handler.DeleteCleanerHandler(cleanerRepository))
		v1.POST("/cleaner/login", config.AuthTokenMiddleware(), handler.LoginCleanerHandler(cleanerRepository))
		v1.POST("/cleaner/:id/img", config.AuthTokenMiddleware(), handler.SendImgProfileHandler(cleanerRepository))
		v1.PATCH("/cleaner/img", config.JWTMiddleware("cleaner"), handler.UpdateImgProfileHandler(cleanerRepository))

		v1.GET("/cleaner/message", config.JWTMiddleware("cleaner"), handler.GetMessageHandler(messageRepository))
		v1.POST("/cleaner/message/:id", config.AuthTokenMiddleware(), handler.CreateMessageHandler(messageRepository, cleanerRepository))

	}
}
