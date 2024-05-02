package main

import (
	"github.com/guicazaroto/learning-go/config"
	"github.com/guicazaroto/learning-go/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	config.InitEnvs()
	// err := config.Init()
	// if err != nil {
	// 	logger.Errorf("Error initializing config")
	// 	return
	// }

	router.Initialize()
}
