package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yenyoong99/goProjects_yyblog/conf"
	"github.com/yenyoong99/goProjects_yyblog/ioc"

	_ "github.com/yenyoong99/goProjects_yyblog/apps"
)

func main() {
	//// user service impl
	//usvc := userimpl.NewUserServiceImpl()
	//
	//// token service impl
	//tsvc := tokenimpl.NewTokenServiceImpl(usvc)
	//
	//// api
	//TokenApiHander := api.NewTokenApiHandler(tsvc)
	//
	//// Protocol
	//engine := gin.Default()

	// 1. init program conf
	if err := conf.LoadFromEnv(); err != nil {
		panic(err)
	}

	// 2. program object manage
	if err := ioc.Init(); err != nil {
		panic(err)
	}

	// Protocol
	engine := gin.Default()

	rr := engine.Group("/yyblog/api/v1")
	//TokenApiHander.Registry(rr)
	ioc.RegistryGinApi(rr)

	// start http service
	if err := engine.Run(":8080"); err != nil {
		panic(err)
	}
}
