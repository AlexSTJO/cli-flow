package cmd

import (
  "encoding/json"
  "os"
  "fmt"
  "path/filepath"
  "bufio"
  "strings"

  "github.com/spf13/cobra"
  "github.com/AlexSTJO/cli-flow/internal/structures"
)

var removestepCmd = &cobra.Command {
  Use: "removestep [workflow_name] [step_name]",
  Short: "Removes a step by workflow name and step name",
  Args: cobra.ExactArgs(2),
  Run: func(cmd *cobra.Command, args[]string) {
    workflow_name := args[0]
    step_name := args[1]

    home, _ := os.UserHomeDir()
    path := filepath.Join(home, ".cli_flow", "workflows", workflow_name+".json")

    data, err := os.ReadFile(path)

    if err != nil {
      fmt.Printf("[Error] Error reading configuration: %v\n", err)
      return
    }

    var wf structures.Workflow
    err = json.Unmarshal(data, &wf)

    if (err != nil) {
      fmt.Printf("[Error] Error parsing json: %v\n", err)
      return
    }

    reader := bufio.NewReader(os.Stdin)

    fmt.Printf("Deleting '%s' Please reenter step name to confirm: ", step_name)
    confirm, _ := reader.ReadString('\n')
    confirm = strings.TrimSpace(confirm)
    if confirm != step_name {
      fmt.Printf("[Error] Confirmation string did not match")
      return
    }
    
    for i, step := range wf.Steps {
      if step_name == step.Name {
        wf.Steps = append(wf.Steps[:i], wf.Steps[i+1:]...)
        break
      }
    }

    file, err := os.Create(path)

    if err != nil {
      fmt.Printf("[Error] Could not save step: %v\n", err)
      return
    }

    defer file.Close()

    encoder := json.NewEncoder(file)
    encoder.SetIndent("", "  ")
    encoder.Encode(wf)

    fmt.Println("[Success] Step succesfully removed")



  },
}
