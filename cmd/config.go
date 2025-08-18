/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage proxce_cli configuration",
	Long: `The config command allows you to view or update configuration settings
	for proxce_cli. It does not perform any action by itself and should be
	used with one of its subcommands:

  	# Set configuration values
  	proxce_cli config set --api-key YOUR_KEY --project-name YOUR_PROJECT

  	# Show current configuration values
  		proxce_cli config show

	Subcommands:
  	set   Set API key or project name
  	show  Display current configuration values`,
	
}

func init() {
	rootCmd.AddCommand(configCmd)
}
