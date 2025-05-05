package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dotfleet",
	Short: "DotFleet manages and syncs dotfiles across machines",
	Long:  `DotFleet is a Git-based dotfile manager that supports machine-specific setups and syncing.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Run `dotfleet init`, `dotfleet sync`, or `dotfleet apply`.")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
