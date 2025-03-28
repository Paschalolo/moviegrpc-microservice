package main

import (
	"time"

	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	logger.Info("started the service ", zap.String("serviceName", "metadata"))
	logger.Info("Request timeout ", zap.Duration("timeout", time.Second*10))
}
