package main

import (
	"flag"

	"go.uber.org/zap"
)

var (
	path = flag.String("c", "./configs/cronjob/block/local.yaml", "set config file path")

	initHeight = flag.Uint64("init-height", 0, "set init height for start")
)

func init() {
	flag.Parse()
}

func main() {
	svc, err := CreateService(*path, *initHeight)
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
