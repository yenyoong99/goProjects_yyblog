package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "yyblog",
	Short: "yyblog api server",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	// register root subcommand
	// Flags return all subcommand flag
	// StringVarP get command line param，convert to string and assign value to specified variable
	// --username -u
	initCmd.Flags().StringVarP(&username, "username", "u", "admin", "administrator username")
	rootCmd.AddCommand(initCmd, startCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
