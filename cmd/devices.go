package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"packet-api-cli/util"
)

var devicesCmd = &cobra.Command{
	Use:   "devices",
	Short: "List devices",

	Run: func(cmd *cobra.Command, args []string) {
		devices, err := util.GetDevices()
		if err != nil {
			fmt.Printf("Error retrieving device list: %v\n", err)
		}

		if len(*devices) == 0 {
			fmt.Println("No available devices")
			os.Exit(0)
		}
	},
}

func init() {
	rootCmd.AddCommand(devicesCmd)
}
