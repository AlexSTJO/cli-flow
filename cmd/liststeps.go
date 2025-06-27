package cmd

import (
  "os"
  "fmt"
  "path/filepath"
  "encoding/json"

  "github.com/spf13/cobra"
  "github.com/AlexSTJO/cli-flow/internal/structures"
)


var liststepsCmd = &cobra.Command {
  Use: "liststeps [workflow_name]",
  Short: "Lists all steps by workflow name",
  Args: cobra.ExactArgs(1),
  Run: func(cmd *cobra.Command, args[]string) {
    workflow_name := args[0]

    home, _ := os.UserHomeDir()
    path := filepath.Join(home, ".cli_flow", "workflows", workflow_name+".json")

    data, err := os.ReadFile(path)

    if err != nil {
      fmt.Printf("[Error] Could not read file: %v", err)
      return
    }

    var wf structures.Workflow
    err = json.Unmarshal(data, &wf)

    if err != nil {
      fmt.Printf("[Error] Error parsing json: %v\n", err)
      return
   	} 
    
    fmt.Println("Workflow Steps: ")
		
    for _, step := range wf.Steps {
			fmt.Printf("  - %s\n", step.Name)
			if (step.Service) == "loop" {
				inners, ok := step.Config["steps"].([]structures.Step)
				if !ok {
					fmt.Printf("[Error] Error converting type\n")
      		return
				}

				for _, inner := range inners {
					fmt.Printf("  -- %s\n", inner.Name)
				}
			} else if (step.Service) == "if" {
				fmt.Println("    <True> ")
				inners, ok := step.Config["true_steps"].([]structures.Step)
				if !ok {
					fmt.Println("    **Empty**")
				}
				for _, inner := range inners {
					fmt.Printf(" -- %s\n", inner.Name)
				}
				
				fmt.Println("    <False> ")
				falseInners, ok := step.Config["false_steps"].([]structures.Step)
				if !ok {
      		fmt.Println("    **Empty**")
				}

				for _, inner := range falseInners {
					fmt.Printf(" -- %s\n", inner.Name)
				}
			}
		}

  },
}
