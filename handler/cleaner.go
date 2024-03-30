package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guicazaroto/learning-go/schemas"
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

func CreateCleanerHandler(ctx *gin.Context) {
	request := CreateUserRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	user := schemas.User{
		Role:     request.Role,
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
		Active:   request.Active,
	}

	if err := db.Create(&user).Error; err != nil {
		logger.Errorf("error creating opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating opening on database")
		return
	}

	sendSuccess(ctx, "create-user", user)

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