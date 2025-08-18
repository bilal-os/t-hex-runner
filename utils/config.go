package utils

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

// InitConfig initializes the configuration for the CLI application.
func InitConfig() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get home directory: %w", err)
	}

	configDir := filepath.Join(home, ".proxce_cli")
	configFile := filepath.Join(configDir, "proxce_cli_config.yaml")

	// Ensure config dir exists
	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		if err := os.Mkdir(configDir, 0700); err != nil {
			return fmt.Errorf("failed to create config directory: %w", err)
		}
	}

	viper.SetConfigFile(configFile)
	viper.SetConfigType("yaml")

	// If file doesn't exist, create it with empty YAML content
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		if err := os.WriteFile(configFile, []byte("---\n"), 0600); err != nil {
			return fmt.Errorf("failed to create config file: %w", err)
		}
		fmt.Println("Created new config file:", configFile)
	}

	// Now read the config (empty YAML is valid)
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to read config: %w", err)
	}

	// Prompt for "T-Hex" key if not set
	if !viper.IsSet("api_key") {
		fmt.Print("Enter your T-Hex key: ")
		var key string
		if _, err := fmt.Scanln(&key); err != nil {
			return fmt.Errorf("failed to read T-Hex key: %w", err)
		}

		viper.Set("api_key", key)
		if err := viper.WriteConfig(); err != nil {
			return fmt.Errorf("failed to save T-Hex key: %w", err)
		}
	}

	//Prompt for "Project Name" if not set
	if !viper.IsSet("project_name") {
		fmt.Print("Enter your Project Name: ")
		var projectName string
		if _, err := fmt.Scanln(&projectName); err != nil {
			return fmt.Errorf("Project Name must have no space. Please try again.")
		}
		viper.Set("project_name", projectName)
		if err := viper.WriteConfig(); err != nil {
			return fmt.Errorf("failed to save Project Name: %w", err)
		}
	}

	return nil
}
