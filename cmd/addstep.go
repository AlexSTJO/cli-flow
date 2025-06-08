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
  Use: "addstep [workflow_name] [loop_name (optional)]",
  Short: "Adds a step to workflow by name, if loop name added step will go into loop",
  Args: cobra.RangeArgs(1,2),
  Run: func(cmd *cobra.Command, args []string) {
    workflow_name := args[0]
    loop_name := ""
    if len(args) == 2 {
        loop_name = args[1]
    }

    home, _ := os.UserHomeDir()
    path := filepath.Join(home, ".cli_flow", "workflows", workflow_name+".json")

    data, err := os.ReadFile(path)
    if err != nil {
        fmt.Printf("[Error] Error reading workflow: %v\n", err)
        return
    }

    var wf structures.Workflow
    if err := json.Unmarshal(data, &wf); err != nil {
        fmt.Printf("[Error] Error parsing json: %v\n", err)
        return
    }

    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Please enter step name/id: ")
    name, _ := reader.ReadString('\n')
    name = strings.TrimSpace(name)

    fmt.Print("Please enter service: ")
    service, _ := reader.ReadString('\n')
    service = strings.TrimSpace(service)

    svc, ok := services.Registry[service]
    if !ok {
        fmt.Printf("[Error] Unknown Service: %s\n", service)
        return
    }

    var config map[string]any

    if customSvc, ok := svc.(interface {
        PromptForConfig() (map[string]any, error)
    }); ok {
        config, err = customSvc.PromptForConfig()
        if err != nil {
            fmt.Printf("[Error] Failed to collect custom config: %v\n", err)
            return
        }
    } else {
        config = make(map[string]any)
        for _, key := range svc.ConfigSpec() {
            fmt.Printf("%s: ", strings.Title(key))
            value, _ := reader.ReadString('\n')
            value = strings.TrimSpace(value)

            if value == "" {
                fmt.Printf("[Error] %s cannot be empty\n", strings.Title(key))
                return
            }

            config[key] = value
        }
    }

    step := structures.Step{
        Name:    name,
        Service: service,
        Config:  config,
    }

    if loop_name != "" {
        found := false
        for i, outerStep := range wf.Steps {
            if outerStep.Name == loop_name && outerStep.Service == "loop" {
                // Marshal step into map[string]any
                stepBytes, _ := json.Marshal(step)
                var stepMap map[string]any
                _ = json.Unmarshal(stepBytes, &stepMap)

                rawLoopSteps, ok := outerStep.Config["steps"].([]any)
                if !ok {
                    rawLoopSteps = []any{}
                }

                rawLoopSteps = append(rawLoopSteps, stepMap)
                outerStep.Config["steps"] = rawLoopSteps
                wf.Steps[i] = outerStep
                found = true
                break
            }
        }
        if !found {
            fmt.Printf("[Error] Loop step '%s' not found in workflow\n", loop_name)
            return
        }
    } else {
        wf.Steps = append(wf.Steps, step)
    }

    file, err := os.Create(path)
    if err != nil {
        fmt.Printf("[Error] Could not save step %v\n", err)
        return
    }
    defer file.Close()

    encoder := json.NewEncoder(file)
    encoder.SetIndent("", "  ")
    encoder.Encode(wf)

    fmt.Println("[Success] Yipee your step has been added :)")
  },
}
