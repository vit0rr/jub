package main

import (
	"github.com/vit0rr/jub/config"
	"github.com/vit0rr/jub/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	err := config.Init()
	if err != nil {
		logger.ErrorF("config initialization error: %v", err)
		return
	}

	router.Initialize()
}
