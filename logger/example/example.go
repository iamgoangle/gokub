package main

import (
	"github.com/iamgoangle/gokub/logger"
)

func main() {
	logger.Debug("Debug -- hello, world")
	logger.Info("Info -- hello, world")
	logger.Infof("Hello %s", "Golf")
}
