package main

import (
	"flag"
)

var path = flag.String("c", "./configs/activity/grpc/local.yaml", "set config file path")

func init() {
	flag.Parse()
}

func main() {
	svc, err := CreateService(*path, 1)
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
