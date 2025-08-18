package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var apiKey string
var projectName string

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set configuration values for proxce_cli",
	Long: `Use this command to configure your proxce_cli settings.
You can set your API key and/or project name. Existing values are preserved if a flag is not provided.

Examples:
  # Set only the API key
  proxce_cli config set --api-key YOUR_KEY

  # Set only the project name
  proxce_cli config set --project-name YOUR_PROJECT

  # Set both API key and project name
  proxce_cli config set --api-key YOUR_KEY --project-name YOUR_PROJECT`,
	RunE: func(cmd *cobra.Command, args []string) error {

		// Check if at least one of the flags is provided
		if apiKey == "" && projectName == "" {
			return fmt.Errorf("Provide at least one flag: --api-key or --project-name")
		}

		// Only set api_key if provided
		if apiKey != "" {
			viper.Set("api_key", apiKey)
			fmt.Printf("Config updated: api_key = %s\n", apiKey)
		}

		// Only set project_name if provided
		if projectName != "" {
			viper.Set("project_name", projectName)
			fmt.Printf("Config updated: project_name = %s\n", projectName)
		}

		// Save to config file
		if err := viper.WriteConfig(); err != nil {
			return fmt.Errorf("failed to save config: %w", err)
		}
		
		fmt.Println("Configuration saved successfully.")
		return nil
	},
}

func init() {
	configCmd.AddCommand(setCmd)

	// Define flags for the set command
	setCmd.Flags().StringVar(&apiKey, "api-key", "", "Set the T-Hex API key")
	setCmd.Flags().StringVar(&projectName, "project-name", "", "Set the project name for T-Hex")
}
