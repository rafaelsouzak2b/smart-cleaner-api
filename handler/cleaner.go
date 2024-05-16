package handler

// manter somente coisas relacionadas aos cleaners, ou seja, usuários com Role = cleaner

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guicazaroto/learning-go/schemas"
	"github.com/guicazaroto/learning-go/util"
)

func GetCleanerHandler(c *gin.Context) {
	var cleaners []schemas.Cleaner
	db.Preload("UserInfos").Find(&cleaners)
	if len(cleaners) == 0 {
		sendError(c, http.StatusNotFound, "cleaners not found")
		return
	}
	sendSuccess(c, "cleaner", cleaners)
}

func GetCleanerByIdHandler(c *gin.Context) {
	cleanerID := c.Param("id")
	var cleaner []schemas.Cleaner
	db.Preload("UserInfos").Find(&cleaner, cleanerID)
	if len(cleaner) == 0 {
		sendError(c, http.StatusNotFound, "cleaner not found")
		return
	}
	sendSuccess(c, "cleaner", cleaner)
}

func CreateCleanerHandler(ctx *gin.Context) {
	request := CleanerRequest{}

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

	cleaner := schemas.Cleaner{
		UserInfos: schemas.User{
			Name:     request.Name,
			Email:    request.Email,
			Password: util.HashString(request.Password),
			Role:     "cleaner",
			Active:   request.Active,
		},
		Telefone:       request.Telefone,
		CPF:            request.CPF,
		DataNascimento: request.DataNascimento,
		Cep:            request.Cep,
		Logradouro:     request.Logradouro,
		Numero:         request.Numero,
		Cidade:         request.Cidade,
		Uf:             request.Uf,
		Descricao:      request.Descricao,
	}

	if err := db.Create(&cleaner).Error; err != nil {
		logger.Errorf("error creating opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating opening on database")
		return
	}

	sendCreated(ctx, "create-cleaner", cleaner)

}

func UpdateCleanerHandler(c *gin.Context) {
	request := CleanerRequest{}
	cleanerID := c.Param("id")
	var cleaner []schemas.Cleaner
	db.Preload("UserInfos").Find(&cleaner, cleanerID)
	if len(cleaner) == 0 {
		sendError(c, http.StatusNotFound, "cleaner not found")
		return
	}

	if err := c.BindJSON(&request); err != nil {
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}
	result := db.Save(&cleaner)
	if result.Error != nil {
		sendError(c, http.StatusInternalServerError, result.Error.Error())
		return
	}
}

// func DeleteCleanerHandler(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "cleaner deleted",
// 	})
// }
