package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "ginbuilder",
	Short: "A CLI tool to generate Gin web project",
	Long: `
ginbuilder is a CLI tool that helps you quickly create a new Gin web project with best practice project structure and common functionalities.
Usage: 
	ginbuilder new --project webserver --package github.com/scliangx/webserver --directory .`,
}

// RegisterCommand 注册子命令到根命令
func RegisterCommand(cmd *cobra.Command) {
	rootCmd.AddCommand(cmd)
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
