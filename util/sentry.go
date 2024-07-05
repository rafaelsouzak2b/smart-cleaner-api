package util

import (
	"bytes"
	"fmt"
	"io"

	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
)

func CaptureRequestBodyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var bodyBytes []byte
		if c.Request.Body != nil {
			bodyBytes, _ = io.ReadAll(c.Request.Body)
		}

		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		hub := sentry.CurrentHub().Clone()
		c.Set("sentryHub", hub)
		fmt.Println(string(bodyBytes))
		hub.ConfigureScope(func(scope *sentry.Scope) {
			scope.SetTag("path", c.FullPath())
			scope.SetTag("method", c.Request.Method)
			scope.SetTag("url", c.Request.URL.String())
			scope.SetContext("request", map[string]interface{}{
				"url":     c.Request.URL.String(),
				"method":  c.Request.Method,
				"headers": c.Request.Header,
				"body":    string(bodyBytes),
				"params":  c.Request.URL.Query(),
			})
		})

		c.Next()
	}
}

func CaptureResponse(c *gin.Context, handler string, response gin.H) {
	hub, _ := c.Get("sentryHub")
	sentryHub := hub.(*sentry.Hub)

	sentryHub.WithScope(func(scope *sentry.Scope) {
		scope.SetTag("handler", handler)
		scope.SetLevel(sentry.LevelInfo)
		scope.SetTag("status_code", fmt.Sprint(c.Writer.Status()))
		scope.SetContext("response", map[string]interface{}{
			"body": response,
		})
		if c.Writer.Status() >= 500 {
			sentryHub.CaptureException(fmt.Errorf("%s request error", handler))
			return
		}
		sentryHub.CaptureMessage(fmt.Sprintf("%s request", handler))
	})
}
