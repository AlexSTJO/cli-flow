package cmd

import (
  "fmt"
  "os"
  "path/filepath"
  "encoding/json"


  //"github.com/AlexSTJO/cli-flow/internal/services"
  "github.com/AlexSTJO/cli-flow/internal/runner"
  "github.com/AlexSTJO/cli-flow/internal/structures"
  "github.com/AlexSTJO/cli-flow/internal/config"
  "github.com/spf13/cobra"
)


var runflowCmd = &cobra.Command{
  Use:    "runflow [workflow_name]",
  Short:  "Runs workflow by name", 
  Args:   cobra.ExactArgs(1),
  Run:    func(cmd *cobra.Command, args []string) {
    workflow_name := args[0]


  
    cfg, err := config.LoadAWSConfig() 

    if err != nil {
      fmt.Printf("[Error] Error loading config: %v", err)
    }
 

    config.SetAWSEnvVars(cfg)
    defer config.UnsetAWSEnvVars()

    

    home,_ := os.UserHomeDir()

    path := filepath.Join(home, ".cli_flow", "workflows", workflow_name+".json")

    data, err := os.ReadFile(path)

    if (err != nil) {
      fmt.Printf("[Error] Error reading file: %v\n", err)
      return
    }

    var wf structures.Workflow

	err = json.Unmarshal(data, &wf)

    if (err != nil) {
      fmt.Printf("[Error] Error unparsing json: %v", err)
      return
    }
 	
	err = runner.RunWorkflow(wf)

	if err != nil{
		fmt.Printf("[Error] Runtime Error: %v", err)
	}
  },
}
