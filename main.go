package main

import (
	"github.com/Ammce/go-banking-core/app"
	"github.com/Ammce/go-banking-core/logger"
)

func main() {
	logger.Info("Server started")
	app.Start()
}
