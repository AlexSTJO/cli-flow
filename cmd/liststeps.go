package cmd

import (
  "os"
  "fmt"
  "path/filepath"
  "encoding/json"

  "github.com/spf13/cobra"
  "github.com/AlexSTJO/cli-flow/internal/structures"
	//"github.com/AlexSTJO/cli-flow/internal/parser"
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
		steps := wf.Steps	
		queue := append([]structures.Step{}, steps...)

		for len(queue) > 0 {
			c := queue[0]
			queue = queue[1:]
			fmt.Println(c.Name)
		}
		
    

  },
}
