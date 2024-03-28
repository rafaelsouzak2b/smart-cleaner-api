package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes (router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
		
		v1.GET("/cleaners", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "all cleaners",
			})
		})

		v1.GET("/cleaners/:id", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "cleaner",
			})
		})

		v1.POST("/cleaners", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "cleaner created",
			})
		})

		v1.PUT("/cleaners/:id", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "cleaner updated",
			})
		})

		v1.DELETE("/cleaners/:id", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "cleaner deleted",
			})
		})
	}
}