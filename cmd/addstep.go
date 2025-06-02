package cmd

import (
  "bufio"
  "encoding/json"
  "fmt"
  "os"
  "path/filepath"
  "strings"

  "github.com/AlexSTJO/cli-flow/internal/services"
	"github.com/AlexSTJO/cli-flow/internal/structures"
	"github.com/spf13/cobra"
)


var addstepCmd = &cobra.Command{
  Use: "addstep [workflow_name]",
  Short: "Adds a step to workflow by name",
  Args: cobra.ExactArgs(1),
  Run: func(cmd *cobra.Command, args[]string) {
    workflow_name := args[0]

    home, _ := os.UserHomeDir()
    path := filepath.Join(home, ".cli_flow", "workflows", workflow_name+".json")

    data, err := os.ReadFile(path)

    if (err != nil) {
      fmt.Printf("[Error] Error reading workflow: %v\n", err)
      return
    }

    var wf structures.Workflow
    err = json.Unmarshal(data, &wf)

    if (err != nil) {
      fmt.Printf("[Error] Error parsing json: %v\n", err)
      return
    }
  

    reader := bufio.NewReader(os.Stdin)
    
    fmt.Print("Please enter step name/id: ")
    name, _ := reader.ReadString('\n')

    fmt.Print("Please enter service: ")
    service, _ := reader.ReadString('\n')
    service = strings.TrimSpace(service)

    svc, ok := services.Registry[service]

    if (!ok) {
      fmt.Printf("[Error] Unknown Service: %s\n", service)
      return
    }

    config := make(map[string]interface{})

    for _, key := range svc.ConfigSpec(){
      fmt.Printf("%s: ", strings.Title(key))
      value, _ := reader.ReadString('\n')
      value = strings.TrimSpace(value)

      if value == "" {
        fmt.Printf("[Error] %s cannot be empty\n", strings.Title(key))
        return
      }

      config[key] = value
    }


    step := structures.Step{
      Name: strings.TrimSpace(name),
      Service: service,
      Config: config,
    }

    wf.Steps = append(wf.Steps, step)
    
    file, err := os.Create(path)

    if err != nil {
      fmt.Printf("[Error] Could not save step %v\n", err)
      return
    }

    defer file.Close()

    encoder := json.NewEncoder(file)
    encoder.SetIndent("", "  ")
    encoder.Encode(wf)

    fmt.Println("[Success] Yipee your step has been added :) ")

  },
}
