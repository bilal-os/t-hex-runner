package cmd

import (
	"github.com/spf13/cobra"
	"os"
	"github.com/bilal-os/t-hex-runner/utils"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "t-hex-runner",
	Short: "Run automated test cases on the T-Hex cloud environment",
    Long: `t-hex-runner is a CLI tool to execute test cases in the T-Hex cloud environment.
It allows you to manage configuration, set API keys, and run tests seamlessly from the command line.

Features:
  - Configure API keys and project settings
  - Run test cases on T-Hex cloud
  - View results and logs

Use "t-hex-runner [command] --help" for more information about a command.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return utils.InitConfig()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	//Define Flags and configuration settings here
}
