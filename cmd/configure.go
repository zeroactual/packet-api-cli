package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"packet-api-cli/util"
)

var token string

var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Setup Packet CLI to connect to Packet API",

	Run: func(cmd *cobra.Command, args []string) {
		configure(false)
	},
}

func init() {
	configureCmd.Flags().StringVarP(&token, "token", "t", "", "Packet API Token")
	rootCmd.AddCommand(configureCmd)
}

func configure(init bool) {
	if init {
		viper.Set("project.token", "")
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		err = viper.WriteConfigAs(fmt.Sprintf("%s/packetconfig.yaml", home))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)		}
	}

	// Prompt for input use current config val as default
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Packet API Token for a project [%s]: ", viper.Get("project.token"))
	scanner.Scan()
	token = scanner.Text()
	if len(token) != 0 {
		viper.Set("project.token", token)
	}

	// Get project info
	project, err := util.GetProject()
	if err != nil {
		fmt.Printf("Encountered an error updating project: %v\n", err)
	}

	// Persist changes
	viper.Set("project.id", project.Id)
	viper.Set("project.name", project.Name)
	err = viper.WriteConfig()
	if err != nil {
		fmt.Printf("Encountered an persisting config: %v\n", err)
		return
	}

	fmt.Println("Configuration updated")
}