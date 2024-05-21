package handler

// manter somente coisas relacionadas aos cleaners, ou seja, usuários com Role = cleaner

import (
	"context"
	"fmt"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	aws_config "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/guicazaroto/learning-go/config"
	"github.com/guicazaroto/learning-go/model"
	"github.com/guicazaroto/learning-go/schemas"
	"github.com/guicazaroto/learning-go/util"
)

func GetCleanerHandler(c *gin.Context) {
	var cleaners []schemas.Cleaner
	db.Preload("UserInfos").Find(&cleaners)
	if len(cleaners) == 0 {
		util.SendError(c, http.StatusNotFound, "cleaners not found")
		return
	}
	cleanersResponse := []model.CleanerResponse{}
	for _, cleaner := range cleaners {
		cleanersResponse = append(cleanersResponse, cleaner.ToResponse())
	}
	util.SendSuccess(c, "get-all-cleaner", cleanersResponse)
}

func GetCleanerByIdHandler(c *gin.Context) {
	cleanerID := c.Param("id")
	var cleaner *schemas.Cleaner
	result := db.Preload("UserInfos").First(&cleaner, cleanerID)
	if result.RowsAffected == 0 {
		util.SendError(c, http.StatusNotFound, "cleaner not found")
		return
	}
	util.SendSuccess(c, "get-cleaner-by-id", cleaner.ToResponse())
}

func CreateCleanerHandler(ctx *gin.Context) {
	request := model.CleanerRequest{}

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
		util.SendError(ctx, http.StatusInternalServerError, "error creating opening on database")
		return
	}

	util.SendCreated(ctx, "create-cleaner", cleaner.ToResponse())

}

func UpdateCleanerHandler(c *gin.Context) {
	request := model.CleanerRequest{}
	cleanerID := c.Param("id")
	var cleaner schemas.Cleaner
	if result := db.Preload("UserInfos").First(&cleaner, cleanerID); result.RowsAffected == 0 {
		util.SendError(c, http.StatusNotFound, "cleaner not found")
		return
	}

	if err := c.BindJSON(&request); err != nil {
		util.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	cleaner.UserInfos = schemas.User{
		Name:      request.Name,
		Email:     request.Email,
		Active:    request.Active,
		Role:      cleaner.UserInfos.Role,
		ImagemUrl: cleaner.UserInfos.ImagemUrl,
		Password:  cleaner.UserInfos.Password,
	}
	cleaner.CPF = request.CPF
	cleaner.Cep = request.Cep
	cleaner.Cidade = request.Cidade
	cleaner.DataNascimento = request.DataNascimento
	cleaner.Descricao = request.Descricao
	cleaner.Logradouro = request.Logradouro
	cleaner.Numero = request.Numero
	cleaner.Telefone = request.Telefone
	cleaner.Uf = request.Uf

	result := db.Save(&cleaner)
	if result.Error != nil {
		util.SendError(c, http.StatusInternalServerError, result.Error.Error())
		return
	}
	util.SendSuccess(c, "update-cleaner", cleaner.ToResponse())
}

func DeleteCleanerHandler(c *gin.Context) {
	cleanerID := c.Param("id")
	var cleaner schemas.Cleaner
	result := db.Unscoped().Delete(&cleaner, cleanerID)
	if result.Error != nil {
		util.SendError(c, http.StatusInternalServerError, result.Error.Error())
		return
	}
	util.SendSuccess(c, "delete-cleaner", "Cleaner deleted successful")
}

func SendImgProfileHandler(c *gin.Context) {
	cleanerID := c.Param("id")
	var cleaner schemas.Cleaner
	if result := db.Preload("UserInfos").First(&cleaner, cleanerID); result.RowsAffected == 0 {
		util.SendError(c, http.StatusNotFound, "cleaner not found")
		return
	}
	file, err := c.FormFile("image")
	if err != nil {
		util.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))

	if !config.AllowedExtensions[ext] {
		util.SendError(c, http.StatusBadRequest, fmt.Sprintf("file type not allowed: %s", ext))
		return
	}

	cfg, err := aws_config.LoadDefaultConfig(context.TODO(), aws_config.WithRegion(config.Environment.AwsRegion))
	if err != nil {
		util.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	client := s3.NewFromConfig(cfg)
	openedFile, err := file.Open()
	if err != nil {
		util.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}
	defer openedFile.Close()

	uploader := manager.NewUploader(client)
	result, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(config.Environment.AwsImgProfileBucket),
		Key:    aws.String(file.Filename),
		Body:   openedFile,
	})

	if err != nil {
		util.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}
	resultUpdate := db.Model(&cleaner.UserInfos).Update("ImagemUrl", result.Location)
	if resultUpdate.Error != nil {
		util.SendError(c, http.StatusInternalServerError, resultUpdate.Error.Error())
		return
	}

	util.SendSuccess(c, "send-img-cleaner", gin.H{"message": "File uploaded successfully", "location": result.Location})
}
