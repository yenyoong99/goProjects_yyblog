package main

import (
	"github.com/gin-gonic/gin"
	"goProjects/yyblog/apps/token/api"
	tokenimpl "goProjects/yyblog/apps/token/impl"
	userimpl "goProjects/yyblog/apps/user/impl"
)

func main() {
	// user service impl
	usvc := userimpl.NewUserServiceImpl()

	// token service impl
	tsvc := tokenimpl.NewTokenServiceImpl(usvc)

	// api
	TokenApiHander := api.NewTokenApiHandler(tsvc)

	// Protocol
	engine := gin.Default()

	rr := engine.Group("/yyblog/api/v1")
	TokenApiHander.Registry(rr)

	// start http service
	if err := engine.Run(":8080"); err != nil {
		panic(err)
	}
}
