package main

import (
	"github.com/whitestudios/TaskFlowAPI/config"
	"github.com/whitestudios/TaskFlowAPI/router"
)

func main() {
	logger := config.GetLogger("main")

	if err := config.Initialize(); err != nil {
		logger.Errf("Configuration inicializatiion error: %v", err)
		return
	}

	router.Initialize()
}
