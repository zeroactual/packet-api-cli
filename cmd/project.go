package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "Get info about the current Packet project",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Project info: Id: %s, Name: %s", viper.Get("project.id"), viper.Get("project.name"))
	},
}

func init() {
	rootCmd.AddCommand(projectCmd)
}
