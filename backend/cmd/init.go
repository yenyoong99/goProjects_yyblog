package cmd

import (
	"fmt"
	"github.com/rs/xid"
	"github.com/spf13/cobra"
	"github.com/yenyoong99/goProjects_yyblog/apps/user"
	"github.com/yenyoong99/goProjects_yyblog/conf"
	"github.com/yenyoong99/goProjects_yyblog/ioc"
)

var (
	// command param, Ex: --username root
	username string
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Init Program",
	Run: func(cmd *cobra.Command, args []string) {

		// 1. init program conf
		cobra.CheckErr(conf.LoadFromEnv())

		// 2. project object manage
		cobra.CheckErr(ioc.Init())

		// 3. init admin user account
		// Create a request object using the constructor
		// user.CreateUserRequest{}
		req := user.NewCreateUserRequest()
		req.Username = username
		req.Password = xid.New().String()
		req.Role = user.RoleAdmin

		fmt.Println("Username: ", req.Username)
		fmt.Println("Password  : ", req.Password)

		// handle unit test exceptions
		u, err := ioc.Controller().Get(user.AppName).(user.Service).CreateUser(
			cmd.Context(),
			req,
		)
		// an error directly to interrupt the unit process and fail.
		cobra.CheckErr(err)

		// Normal printing object
		fmt.Println(u)
	},
}
