package util

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendError(ctx *gin.Context, code int, msg string) gin.H {
	body := gin.H{
		"message":   msg,
		"errorCode": code,
	}
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, body)
	return body
}

func SendSuccess(ctx *gin.Context, op string, data interface{}) gin.H {
	code := http.StatusOK
	body := gin.H{
		"message":     fmt.Sprintf("operation from handler: %s successfull", op),
		"data":        data,
		"status_code": http.StatusOK,
	}
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, body)
	return body
}

func SendCreated(ctx *gin.Context, op string, data interface{}) gin.H {
	body := gin.H{
		"message": fmt.Sprintf("operation from handler: %s successfull", op),
		"data":    data,
	}
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusCreated, body)
	return body
}

func ErrParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}
