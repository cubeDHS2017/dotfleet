
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the dotfleet configuration",
	Long:  `Sets up the dotfleet configuration (config.toml) to manage dotfiles for different machines.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Here, you can generate or check for the config.toml file
		fmt.Println("Initializing dotfleet configuration...")
		// For now, just create a simple config.toml if it doesn't exist
		configFile := "config.toml"
		if _, err := os.Stat(configFile); os.IsNotExist(err) {
			// Create a basic config.toml
			file, err := os.Create(configFile)
			if err != nil {
				fmt.Println("Error creating config.toml:", err)
				return
			}
			defer file.Close()

			// Example basic content for config.toml
			file.WriteString(`
[machine]
name = "my-laptop"
category = "laptop"
`)
			fmt.Println("config.toml created successfully.")
		} else {
			fmt.Println("config.toml already exists.")
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
