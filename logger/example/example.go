package main

import "github.com/iamgoangle/gokub/logger"

func main() {
	//logger.Debug("Debug -- hello, world")
	//logger.Info("Info -- hello, world")
	//logger.Infof("Hello %s", "Golf")

	//logger.InfoWithFields("[test]: golf", logger.Fields{
	//	logger.Field{Key: "test", Val: "value"},
	//})

	logger.WithFields("test", "golf").Info("test test")
}
