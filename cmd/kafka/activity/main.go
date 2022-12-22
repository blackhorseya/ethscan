package main

import (
	"flag"

	"go.uber.org/zap"
)

var path = flag.String("c", "./configs/kafka/activity/local.yaml", "set config file path")

func init() {
	_ = flag.Parsed()
}

func main() {
	svc, err := CreateService(*path, 1)
	if err != nil {
		zap.S().Fatal(zap.Error(err))
	}

	err = svc.Start()
	if err != nil {
		zap.S().Fatal(zap.Error(err))
	}

	err = svc.AwaitSignal()
	if err != nil {
		zap.S().Fatal(zap.Error(err))
	}
}
