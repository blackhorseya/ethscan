package main

import (
	"flag"
)

var path = flag.String("c", "./configs/activity/restful/local.yaml", "set config file path")

func init() {
	flag.Parse()
}

// @title ethscan API
// @version 0.0.1
// @description API for ethscan
//
// @contact.name sean.zheng
// @contact.email blackhorseya@gmail.com
// @contact.url https://blog.seancheng.space
//
// @license.name GPL-3.0
// @license.url https://spdx.org/licenses/GPL-3.0-only.html
//
// @BasePath /api
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
