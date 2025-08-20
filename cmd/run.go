/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
	"github.com/bilal-os/t-hex-Runner/utils"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Execute a Selenium test file through the T-Hex proxy",
	Long: `The run command executes a Selenium test file on the T-Hex cloud environment
	via a local proxy. It automatically starts the proxy, waits for it to initialize,
	and then runs the specified test file.

	You must provide exactly one argument: the path to the test file.

	Examples:
  	# Run a test file
  	t-hex-runner run tests/sample_test.py

  	# Run another test file
  	t-hex-runner run tests/login_test.py`,
	Args: cobra.ExactArgs(1), // expects exactly one argument: test file path
	Run: func(cmd *cobra.Command, args []string) {
		testFile := args[0]

		// create a channel to wait for proxy readiness
		ready := make(chan struct{})

		// Start proxy
		proxy, err := utils.StartProxy(ready)
		if err != nil {
			log.Fatalf("Failed to run test case(s): %v", err)
		}

		// wait until proxy signals readiness
		<-ready

		// Step 2: Run Selenium test
		log.Printf("Running Selenium test: %s", testFile)
		// Suppose your proxy runs at :8888
		executorURL := "http://localhost:8888/"
		cmdPython := exec.Command("python3", testFile, executorURL)
		cmdPython.Stdout = os.Stdout
		cmdPython.Stderr = os.Stderr
		if err := cmdPython.Run(); err != nil {
			log.Fatalf("Test failed: %v", err)
		}

		// Stop proxy after test
		if err := proxy.Stop(); err != nil {
			log.Fatalf("Failed to stop proxy: %v", err)
		}

		log.Println("Test completed successfully!")
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
