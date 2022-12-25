package main

import (
	"flag"
)

var (
	path = flag.String("c", "./configs/cronjob/block/local.yaml", "set config file path")

	initHeight = flag.Uint64("init-height", 16258466, "set init height for start")
)

func init() {
	flag.Parse()
}

func main() {
	svc, err := CreateService(*path, *initHeight)
	if err != nil {
		panic(err)
	}

	err = svc.Start()
	if err != nil {
		panic(err)
	}

	err = svc.AwaitSignal()
	if err != nil {
		panic(err)
	}
}
