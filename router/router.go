package router

import (
	"log"

	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	"github.com/guicazaroto/learning-go/config"
	"github.com/guicazaroto/learning-go/util"
)

func Initialize() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn:           config.Environment.SentryDns,
		Environment:   config.Environment.Environment,
		EnableTracing: true,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	r := gin.Default()
	r.Use(util.CORSMiddleware())
	r.Use(sentrygin.New(sentrygin.Options{
		Repanic: true,
	}))

	r.Use(util.CaptureRequestBodyMiddleware())

	initializeRoutes(r)
	r.Run()
}
