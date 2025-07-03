package cmd

import (
  "fmt"
  "os"
  "path/filepath"
  "encoding/json"

  "github.com/AlexSTJO/cli-flow/internal/runner"
  "github.com/AlexSTJO/cli-flow/internal/structures"
  "github.com/AlexSTJO/cli-flow/internal/config"
  "github.com/spf13/cobra"
)



var runCmd = &cobra.Command{
	Use:   "run [workflow_name]",
	Short: "Runs workflow by name",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		workflowName := args[0]
		cfg, err := config.LoadAWSConfig()
		if err == nil {
			config.SetAWSEnvVars(cfg)
			defer config.UnsetAWSEnvVars()
		}

		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Printf(red("[Error] Failed to get user home directory: %v\n"), err)
			return
		}

		path := filepath.Join(home, ".cli_flow", "workflows", workflowName+".json")
		data, err := os.ReadFile(path)
		if err != nil {
			fmt.Printf(red("[Error] Could not read workflow file: %v\n"), err)
			return
		}

		var wf structures.Workflow
		if err := json.Unmarshal(data, &wf); err != nil {
			fmt.Printf(red("[Error] Failed to parse workflow: %v\n"), err)
			return
		}

		if err := runner.RunWorkflow(wf); err != nil {
			fmt.Printf(red("[Error] Workflow failed: %v\n"), err)
		} else {
			fmt.Println(green("[Success] Workflow ran successfully."))
		}
	},
}

