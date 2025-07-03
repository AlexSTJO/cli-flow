package cmd

import (
  "os"
  "fmt"
  "path/filepath"
  "encoding/json"

  "github.com/spf13/cobra"
  "github.com/AlexSTJO/cli-flow/internal/structures"
	"github.com/AlexSTJO/cli-flow/internal/formatter"
)


var mapCmd = &cobra.Command{
	Use:   "map [workflow_name]",
	Short: "Map all steps by workflow name",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		workflowName := args[0]

		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Printf("[Error] Could not determine home directory: %v\n", err)
			return
		}

		path := filepath.Join(home, ".cli_flow", "workflows", workflowName+".json")
		data, err := os.ReadFile(path)
		if err != nil {
			fmt.Printf("[Error] Could not read workflow file at %s: %v\n", path, err)
			return
		}

		var wf structures.Workflow
		if err := json.Unmarshal(data, &wf); err != nil {
			fmt.Printf("[Error] Failed to parse workflow JSON: %v\n", err)
			return
		}

		fmt.Println("Workflow Steps:")
		fmt.Println("----------------")

		for i, step := range wf.Steps {
			formatter.PrintStepTree("", step, i == len(wf.Steps)-1)
		}
	},
}

