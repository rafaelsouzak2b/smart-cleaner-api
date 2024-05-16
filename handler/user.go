package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guicazaroto/learning-go/schemas"
)

func GetUserHandler(c *gin.Context) {
	var users []schemas.User
	db.Find(&users, "Role = ?", "user")
	sendSuccess(c, "user", users)
}

func GetUserByIdHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "user",
	})
}

func CreateUserHandler(ctx *gin.Context) {
	request := CreateUserRequest{}

	if err := ctx.BindJSON(&request); err != nil {
		logger.Errorf("body error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

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

	sendCreated(ctx, "user", user)
}

func UpdateUserHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "user updated",
	})
}

func DeleteUserHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "user deleted",
	})
}
