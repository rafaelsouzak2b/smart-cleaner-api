package handler

import (
	"context"
	"fmt"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	aws_config "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/guicazaroto/learning-go/config"
	"github.com/guicazaroto/learning-go/model"
	"github.com/guicazaroto/learning-go/repository"
	"github.com/guicazaroto/learning-go/schemas"
	"github.com/guicazaroto/learning-go/util"
)

func GetCleanerHandler(repository repository.ICleanerRepositoryport) gin.HandlerFunc {
	return func(c *gin.Context) {
		var response gin.H
		defer util.CaptureResponse(c, "GetCleanerHandler", response)

		city := c.Query("city")

		cleaners := repository.GetCleaners(city)

		if len(cleaners) == 0 {
			response = util.SendError(c, http.StatusNotFound, "cleaners not found")
			return
		}
		cleanersResponse := []model.CleanerResponse{}
		for _, cleaner := range cleaners {
			cleanersResponse = append(cleanersResponse, cleaner.ToResponse())
		}
		response = util.SendSuccess(c, "get-all-cleaner", cleanersResponse)
	}
}

func GetCleanerByIdHandler(repository repository.ICleanerRepositoryport) gin.HandlerFunc {
	return func(c *gin.Context) {
		var response gin.H
		defer util.CaptureResponse(c, "GetCleanerByIdHandler", response)

		cleanerID := c.Param("id")
		cleaner := repository.GetCleanerById(cleanerID)
		if cleaner == nil {
			response = util.SendError(c, http.StatusNotFound, "cleaners not found")
			return
		}
		response = util.SendSuccess(c, "get-cleaner-by-id", cleaner.ToResponse())
	}
}

func GetCleanerMeByIdHandler(repository repository.ICleanerRepositoryport) gin.HandlerFunc {
	return func(c *gin.Context) {
		var response gin.H
		defer util.CaptureResponse(c, "GetCleanerMeByIdHandler", response)

		cleanerID := c.GetString("id")
		cleaner := repository.GetCleanerById(cleanerID)
		if cleaner == nil {
			response = util.SendError(c, http.StatusNotFound, "cleaners not found")
			return
		}
		response = util.SendSuccess(c, "get-cleaner-me-by-id", cleaner.ToResponseMe())
	}
}

func CreateCleanerHandler(repository repository.ICleanerRepositoryport) gin.HandlerFunc {
	return func(c *gin.Context) {
		var response gin.H
		defer util.CaptureResponse(c, "CreateCleanerHandler", response)

		request := model.CleanerRequest{}

		if err := c.BindJSON(&request); err != nil {
			response = util.SendError(c, http.StatusBadRequest, err.Error())
			return
		}

		if err := request.Validate(); err != nil {
			response = util.SendError(c, http.StatusBadRequest, err.Error())
			return
		}

		count := repository.GetCleanerByEmailAndCpf(request.Email, request.CPF)
		if count > 0 {
			response = util.SendError(c, http.StatusConflict, "already registered cleaner")
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

		if err := repository.CreateCleaner(&cleaner); err != nil {
			response = util.SendError(c, http.StatusInternalServerError, "error creating opening on database")
			return
		}

		response = util.SendCreated(c, "create-cleaner", cleaner.ToResponse())
	}
}

func UpdateCleanerHandler(repository repository.ICleanerRepositoryport) gin.HandlerFunc {
	return func(c *gin.Context) {
		var response gin.H
		defer util.CaptureResponse(c, "UpdateCleanerHandler", response)

		request := model.CleanerRequest{}
		cleanerID := c.GetString("id")

		cleaner := repository.GetCleanerById(cleanerID)
		if cleaner == nil {
			response = util.SendError(c, http.StatusNotFound, "cleaners not found")
			return
		}

		if err := c.BindJSON(&request); err != nil {
			response = util.SendError(c, http.StatusBadRequest, err.Error())
			return
		}

		cleaner.UserInfos = schemas.User{
			Name:      request.Name,
			Active:    cleaner.UserInfos.Active,
			Role:      cleaner.UserInfos.Role,
			ImagemUrl: cleaner.UserInfos.ImagemUrl,
			Password:  cleaner.UserInfos.Password,
			Email:     cleaner.UserInfos.Email,
			Model:     cleaner.Model,
		}
		cleaner.Cep = request.Cep
		cleaner.Cidade = request.Cidade
		cleaner.Descricao = request.Descricao
		cleaner.Logradouro = request.Logradouro
		cleaner.Numero = request.Numero
		cleaner.Telefone = request.Telefone
		cleaner.Uf = request.Uf

		if err := repository.SaveCleaner(cleaner); err != nil {
			response = util.SendError(c, http.StatusInternalServerError, err.Error())
			return
		}
		response = util.SendSuccess(c, "update-cleaner", cleaner.ToResponseMe())
	}
}

func DeleteCleanerHandler(repository repository.ICleanerRepositoryport) gin.HandlerFunc {
	return func(c *gin.Context) {
		var response gin.H
		defer util.CaptureResponse(c, "DeleteCleanerHandler", response)

		cleanerID := c.GetString("id")
		if err := repository.DeleteCleaner(cleanerID); err != nil {
			response = util.SendError(c, http.StatusInternalServerError, err.Error())
			return
		}
		response = util.SendSuccess(c, "delete-cleaner", "Cleaner deleted successful")
	}
}

func SendImgProfileHandler(repository repository.ICleanerRepositoryport) gin.HandlerFunc {
	return func(c *gin.Context) {
		var response gin.H
		defer util.CaptureResponse(c, "SendImgProfileHandler", response)

		cleanerID := c.Param("id")
		cleaner := repository.GetCleanerById(cleanerID)
		if cleaner == nil {
			response = util.SendError(c, http.StatusNotFound, "cleaners not found")
			return
		}
		if cleaner.UserInfos.ImagemUrl != "" {
			response = util.SendError(c, http.StatusBadRequest, "image already sent")
			return
		}
		file, err := c.FormFile("file")
		if err != nil {
			response = util.SendError(c, http.StatusBadRequest, err.Error())
			return
		}

		ext := strings.ToLower(filepath.Ext(file.Filename))

		if !config.AllowedExtensions[ext] {
			response = util.SendError(c, http.StatusBadRequest, fmt.Sprintf("file type not allowed: %s", ext))
			return
		}

		cfg, err := aws_config.LoadDefaultConfig(context.TODO(), aws_config.WithRegion(config.Environment.AwsRegion))
		if err != nil {
			response = util.SendError(c, http.StatusInternalServerError, err.Error())
			return
		}

		openedFile, err := file.Open()
		if err != nil {
			response = util.SendError(c, http.StatusInternalServerError, err.Error())
			return
		}
		defer openedFile.Close()

		client := s3.NewFromConfig(cfg)

		fileName := fmt.Sprintf("%s%s", uuid.New().String(), ext)

		uploader := manager.NewUploader(client)
		result, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
			Bucket: aws.String(config.Environment.AwsImgProfileBucket),
			Key:    aws.String(fileName),
			Body:   openedFile,
			ACL:    types.ObjectCannedACLPublicRead,
		})

		if err != nil {
			response = util.SendError(c, http.StatusInternalServerError, err.Error())
			return
		}
		if err := repository.UpdateImgUrlCleaner(cleaner, result.Location); err != nil {
			response = util.SendError(c, http.StatusInternalServerError, err.Error())
			return
		}

		response = util.SendSuccess(c, "send-img-cleaner", gin.H{"message": "File uploaded successfully", "location": result.Location})
	}
}

func UpdateImgProfileHandler(repository repository.ICleanerRepositoryport) gin.HandlerFunc {
	return func(c *gin.Context) {
		var response gin.H
		defer util.CaptureResponse(c, "UpdateImgProfileHandler", response)

		cleanerID := c.GetString("id")
		cleaner := repository.GetCleanerById(cleanerID)
		if cleaner == nil {
			response = util.SendError(c, http.StatusNotFound, "cleaners not found")
			return
		}
		file, err := c.FormFile("file")
		if err != nil {
			response = util.SendError(c, http.StatusBadRequest, err.Error())
			return
		}

		ext := strings.ToLower(filepath.Ext(file.Filename))

		if !config.AllowedExtensions[ext] {
			response = util.SendError(c, http.StatusBadRequest, fmt.Sprintf("file type not allowed: %s", ext))
			return
		}

		cfg, err := aws_config.LoadDefaultConfig(context.TODO(), aws_config.WithRegion(config.Environment.AwsRegion))
		if err != nil {
			response = util.SendError(c, http.StatusInternalServerError, err.Error())
			return
		}

		openedFile, err := file.Open()
		if err != nil {
			response = util.SendError(c, http.StatusInternalServerError, err.Error())
			return
		}
		defer openedFile.Close()

		client := s3.NewFromConfig(cfg)

		go func() {
			fileUrl := strings.Split(cleaner.UserInfos.ImagemUrl, "/")
			fileName := fileUrl[len(fileUrl)-1]

			_, err = client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
				Bucket: aws.String(config.Environment.AwsImgProfileBucket),
				Key:    aws.String(fileName),
			})
			if err != nil {
				response = util.SendError(c, http.StatusInternalServerError, err.Error())
				return
			}

			waiter := s3.NewObjectNotExistsWaiter(client)
			err = waiter.Wait(context.TODO(), &s3.HeadObjectInput{
				Bucket: aws.String(config.Environment.AwsImgProfileBucket),
				Key:    aws.String(fileName),
			}, *aws.Duration(time.Minute * 1))
			if err != nil {
				response = util.SendError(c, http.StatusInternalServerError, err.Error())
				return
			}

		}()

		fileName := fmt.Sprintf("%s%s", uuid.New().String(), ext)

		uploader := manager.NewUploader(client)
		result, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
			Bucket: aws.String(config.Environment.AwsImgProfileBucket),
			Key:    aws.String(fileName),
			Body:   openedFile,
			ACL:    types.ObjectCannedACLPublicRead,
		})

		if err != nil {
			response = util.SendError(c, http.StatusInternalServerError, err.Error())
			return
		}

		if err := repository.UpdateImgUrlCleaner(cleaner, result.Location); err != nil {
			response = util.SendError(c, http.StatusInternalServerError, err.Error())
			return
		}

		response = util.SendSuccess(c, "update-img-cleaner", gin.H{"message": "File uploaded successfully", "location": result.Location})
	}
}

func LoginCleanerHandler(repository repository.ICleanerRepositoryport) gin.HandlerFunc {
	return func(c *gin.Context) {
		var response gin.H
		defer util.CaptureResponse(c, "LoginCleanerHandler", response)

		var creds struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		if err := c.BindJSON(&creds); err != nil {
			response = util.SendError(c, http.StatusBadRequest, "Invalid request")
			return
		}
		cleaner := repository.GetCleanerByEmailAndPassword(creds.Email, creds.Password)
		if cleaner == nil {
			response = util.SendError(c, http.StatusUnauthorized, "Unauthorized")
			return
		}

		tokenString, err := config.GenerateJWT(fmt.Sprint(cleaner.Id), "cleaner")
		if err != nil {
			response = util.SendError(c, http.StatusInternalServerError, "Could not generate token")
			return
		}
		response = gin.H{"token": tokenString}
		c.JSON(http.StatusOK, response)
	}
}
