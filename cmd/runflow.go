package cmd

import (
  "fmt"
  "os"
  "path/filepath"


  "github.com/AlexSTJO/cli-flow/internal/services"
	"github.com/AlexSTJO/cli-flow/internal/structures"
	"github.com/spf13/cobra"
)


var runflow = &cobra.Command{
  Use:    "run [workflow_name]",
  Short:  "Runs workflow by name", 
  Args:   cobra.ExactArgs(1),
  Run:    func(cmd *cobra.Command, args []string) {
    workflow_name := args[0]


    home,_ := os.UserHomeDir()

    path := filepath.join(home, ".cli_flow", "workflows", workflow_name+".json")

    data, err = os.ReadFile(path)

    if (err != nil) {
      fmt.Printf("[Error] Error reading file: %v\n", err)
      return
    }

    var wf structures.Workflow

    err := json.Unmarshal(data, &wf)

    if (err != nil) {
      fmt.Printf("[Error] Error unparsing json: %v", err)
      return
    }

    fmt.Printf("[Workflow] Running workflow: %s\n", wf.Name)

    for _, step := range wf.Steps {
      fmt.Printf("[Workflow] Executing step: %s\n", step.Name)
      fmt.Printf("[Workflow] Utilizing service: %s\n", step.Service)

      svc, ok := services.Registry[step.Service]

      if (!ok) {
        fmt.Printf("[Error] Unknown Service: %s\n", step.Service)
        return
      }

      err := svc.Run(step)

      if (err != nil) {
        fmt.Printf("[Error] Step Failed: %v\n", err)
        return
      }
      
      fmt.Println("[Workflow] Step Executed Succesfully")

    }

    fmt.Println("[Success] Flow Succesful")
  }
}
