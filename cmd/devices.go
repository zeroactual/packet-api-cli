package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"packet-api-cli/util"
)

// deviceCmd represents the device command
var devicesCmd = &cobra.Command{
	Use:   "devices",
	Short: "List devices",

	Run: func(cmd *cobra.Command, args []string) {
		_, err := util.GetDevices()
		if err != nil {
			fmt.Printf("Error retrieving device list: %v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(devicesCmd)
}
