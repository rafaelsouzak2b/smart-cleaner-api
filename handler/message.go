package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guicazaroto/learning-go/model"
	"github.com/guicazaroto/learning-go/schemas"
	"github.com/guicazaroto/learning-go/util"
)

func GetMessageHandler(c *gin.Context) {
	var messages []schemas.Message

	cleanerID := c.MustGet("id").(string)

	db.Where("cleaner_id = ?", cleanerID).Find(&messages)

	messagesResponse := []model.MessageResponse{}
	for _, message := range messages {
		messagesResponse = append(messagesResponse, message.ToResponse())
	}
	util.SendSuccess(c, "get-all-messages", messagesResponse)
}

func CreateMessageHandler(ctx *gin.Context) {
	request := model.MessageRequest{}
	cleanerID := ctx.Param("id")
	if err := ctx.BindJSON(&request); err != nil {
		logger.Errorf("body error: %v", err.Error())
		util.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		util.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	var cleaner schemas.Cleaner
	if result := db.Preload("UserInfos").First(&cleaner, cleanerID); result.RowsAffected == 0 {
		util.SendError(ctx, http.StatusNotFound, "cleaner not found")
		return
	}

	message := schemas.Message{
		CleanerId: cleaner.Id,
		Message:   request.Message,
		Telefone:  request.Telefone,
	}

	if err := db.Create(&message).Error; err != nil {
		logger.Errorf("error creating opening: %v", err.Error())
		util.SendError(ctx, http.StatusInternalServerError, "error creating opening on database")
		return
	}

	util.SendCreated(ctx, "create-cleaner-message", message.ToResponse())

}
