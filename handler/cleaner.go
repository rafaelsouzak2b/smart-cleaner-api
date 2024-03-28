package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCleanerHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "all cleaners",
	})
}

func GetCleanerByIdHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "cleaner",
	})
}

func CreateCleanerHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "cleaner created",
	})
}

func UpdateCleanerHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "cleaner updated",
	})
}

func DeleteCleanerHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "cleaner deleted",
	})
}