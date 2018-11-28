package cmd

import (
	"bufio"
	"fmt"
	. "os"
	"strings"

	"github.com/spf13/cobra"

	"packet-api-cli/util"
)

var addDeviceCmd = &cobra.Command{
	Use:   "add",
	Short: "Wizard to add a new device to project",

	Run: func(cmd *cobra.Command, args []string) {
		addDevice()
	},
}

func init() {
	devicesCmd.AddCommand(addDeviceCmd)
}


func addDevice() {
	preFetch()

	fmt.Println("Please choose datacenter:")

	// Choose facility
	facilities := util.GetFacilities()
	for i, facility := range *facilities {
		fmt.Printf("[%d] %s - %s\n", i, strings.ToUpper(facility.Code), facility.Name)
	}
	var facility = (*facilities)[util.NextInt(len(*facilities))]

	// Choose instance type
	fmt.Println("\nPlease choose device type:")
	for i, plan := range *facility.Plans {
		fmt.Printf("[%d] %s\n", i, plan.Name)
	}
	var plan = (*facility.Plans)[util.NextInt(len(*facility.Plans))]

	// Choose OS
	fmt.Println("\nPlease choose Operating System:")
	for i, os := range *plan.Os {
		fmt.Printf("[%d] %s\n", i, os.Name)
	}
	var os = (*plan.Os)[util.NextInt(len(*plan.Os))]

	// Prompt for hostname
	fmt.Print("Input hostname: ")
	scanner := bufio.NewScanner(Stdin)
	scanner.Scan()
	hostname := scanner.Text()

	// Create
	err := util.CreateDevice(facility.Id, plan.Id, os.Id, hostname)
	if err != nil {
		fmt.Printf("Error encountered creating device: %v", err)
		Exit(1)
	}
	fmt.Println("Device slated for creation")
}

func preFetch() {
	c := make(chan error, 3)
	defer close(c)

	go func() {
		c <- util.LoadFacilities()
	}()

	go func() {
		c <- util.LoadPlans()
	}()

	go func() {
		c <- util.LoadOs()
	}()

	if util.HandleErrs(<-c, <-c, <-c) {
		Exit(1)
	}
}