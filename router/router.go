package router

import (
	"github.com/gin-gonic/gin"
	"github.com/guicazaroto/learning-go/util"
)

func Initialize() {
	r := gin.Default()
	r.Use(util.CORSMiddleware())

	initializeRoutes(r)
	r.Run()
}
