/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Display the current proxce_cli configuration",
	Long: `The show command prints all configuration values currently stored
	in the proxce_cli config file. It helps you verify what API key,
	project name, or other settings are active.

  	Examples:
  	# Show all current configuration values
  	proxce_cli config show

  	# Check if the API key or project name is set
  	proxce_cli config show`,
	RunE: func(cmd *cobra.Command, args []string) error {

		// Fetch all config values
		settings := viper.AllSettings()
		if len(settings) == 0 {
			fmt.Println("⚠️  No configuration found")
			return nil
		}

		fmt.Println("📌 Current Configuration:")
		for k, v := range settings {
			fmt.Printf("  %s: %v\n", k, v)
		}

		return nil
	},
}

func init() {
	configCmd.AddCommand(showCmd)
	// Here you will define your flags and configuration settings.
}
