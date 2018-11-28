package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"packet-api-cli/util"
)

var deviceRemoveCmd = &cobra.Command{
	Use:   "delete",
	Short: "Wizard to delete a device from project",

	Run: func(cmd *cobra.Command, args []string) {
		devices, err := util.GetDevices()
		if err != nil {
			fmt.Printf("Error retrieving device list: %v\n", err)
		}

		if len(*devices) == 0 {
			fmt.Println("No available devices")
			os.Exit(0)
		}

		// Get user choice
		var device = (*devices)[util.NextInt(len(*devices))]

		err = util.DeleteDevice(device.Id)
		if err != nil {
			fmt.Printf("Error deleting device: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Device slated for deletion")
	},
}

func init() {
	devicesCmd.AddCommand(deviceRemoveCmd)
}
