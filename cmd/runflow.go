package cmd

import (
  "fmt"
  "os"
  "path/filepath"
  "encoding/json"


  "github.com/AlexSTJO/cli-flow/internal/services"
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

    fmt.Printf("[Workflow] Running workflow: %s\n", wf.Name)

    var ctx structures.Context = map[string]any{}

    for _, step := range wf.Steps {
      fmt.Printf("[Workflow] Executing step: %s\n", step.Name)
      fmt.Printf("[Workflow] Utilizing service: %s\n", step.Service)

      svc, ok := services.Registry[step.Service]

      if (!ok) {
        fmt.Printf("[Error] Unknown Service: %s\n", step.Service)
        return
      }

      step.Config["__context"] = ctx

      stepCtx, err := svc.Run(step)

      if (err != nil) {
        fmt.Printf("[Error] Step Failed: %v\n", err)
        return
      }
      

      ctx[step.Name] = map[string]any{}
      
      if stepCtxMap, ok := ctx[step.Name].(map[string]any); ok {
        for k,v := range stepCtx {
          stepCtxMap[k] = v
        }

        ctx[step.Name] = stepCtxMap
      } else {
        fmt.Printf("[Error] Low-key I do not know that you are doing if you hit this error, you somehow messed up a cast of a string to a map, cringe")
        return
      }
      
      fmt.Println("[Workflow] Step Executed Succesfully")

    }

    fmt.Println("[Success] Flow Succesful")

    
  },
}
