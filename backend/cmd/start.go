package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/yenyoong99/goProjects_yyblog/conf"
	"github.com/yenyoong99/goProjects_yyblog/ioc"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start server",
	Run: func(cmd *cobra.Command, args []string) {
		// 1. init program conf
		cobra.CheckErr(conf.LoadFromEnv())

		// 2. program object manage
		cobra.CheckErr(ioc.Init())

		// Protocol
		engine := gin.Default()

		rr := engine.Group("/yyblog/api/v1")
		ioc.RegistryGinApi(rr)

		// start http service
		cobra.CheckErr(engine.Run(":8080"))
	},
}
