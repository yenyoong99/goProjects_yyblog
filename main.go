package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yenyoong99/goProjects_yyblog/apps/token/api"
	tokenimpl "github.com/yenyoong99/goProjects_yyblog/apps/token/impl"
	userimpl "github.com/yenyoong99/goProjects_yyblog/apps/user/impl"
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
