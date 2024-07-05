package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guicazaroto/learning-go/model"
	"github.com/guicazaroto/learning-go/repository"
	"github.com/guicazaroto/learning-go/schemas"
	"github.com/guicazaroto/learning-go/util"
)

func GetMessageHandler(repository repository.IMessageRepositoryport) gin.HandlerFunc {
	return func(c *gin.Context) {
		var response gin.H
		defer util.CaptureResponse(c, "GetMessageHandler", response)
		cleanerID := c.GetString("id")

		messages := repository.GetMessagesByCleanerId(cleanerID)

		messagesResponse := []model.MessageResponse{}
		for _, message := range messages {
			messagesResponse = append(messagesResponse, message.ToResponse())
		}
		response = util.SendSuccess(c, "get-all-messages", messagesResponse)
	}
}

func CreateMessageHandler(repository repository.IMessageRepositoryport, repositoryCleaner repository.ICleanerRepositoryport) gin.HandlerFunc {
	return func(c *gin.Context) {
		var response gin.H
		defer util.CaptureResponse(c, "CreateMessageHandler", response)
		request := model.MessageRequest{}
		cleanerID := c.Param("id")
		if err := c.BindJSON(&request); err != nil {
			logger.Errorf("body error: %v", err.Error())
			response = util.SendError(c, http.StatusBadRequest, err.Error())
			return
		}

		if err := request.Validate(); err != nil {
			logger.Errorf("validation error: %v", err.Error())
			response = util.SendError(c, http.StatusBadRequest, err.Error())
			return
		}

		cleaner := repositoryCleaner.GetCleanerById(cleanerID)

		if cleaner == nil {
			response = util.SendError(c, http.StatusNotFound, "cleaner not found")
			return
		}

		message := schemas.Message{
			CleanerId: cleaner.Id,
			Message:   request.Message,
			Telefone:  request.Telefone,
			Nome:      request.Nome,
			Email:     request.Email,
		}

		if err := repository.CreateMessage(&message); err != nil {
			logger.Errorf("error creating opening: %v", err.Error())
			response = util.SendError(c, http.StatusInternalServerError, "error creating opening on database")
			return
		}

		response = util.SendCreated(c, "create-cleaner-message", message.ToResponse())
	}
}
