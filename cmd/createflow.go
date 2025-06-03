package cmd

import (
  "bufio"
  "encoding/json"
  "fmt"
  "os"
  "path/filepath"
  "strings"

  "github.com/spf13/cobra"
  "github.com/AlexSTJO/cli-flow/internal/structures"
)




var createflowCmd = &cobra.Command {
  Use: "createflow",
  Short: "Create a new workflow",
  Run: func(cmd *cobra.Command, args []string) {
    reader := bufio.NewReader(os.Stdin)

    fmt.Print("Workflow Name: ")
    name, _ := reader.ReadString('\n')

    fmt.Print("Description: ")
    description, _ := reader.ReadString('\n')

    workflow_name := strings.TrimSpace(name)
    workflow_description := strings.TrimSpace(description)

    wf := structures.Workflow{
      Name: workflow_name,
      Description: workflow_description,
      Steps: []structures.Step{},
    }

    home_directory, _ := os.UserHomeDir()
    dir := filepath.Join(home_directory, ".cli_flow", "workflows")
    os.MkdirAll(dir, os.ModePerm)

    file_path := filepath.Join(dir, workflow_name+".json")
    file, err := os.Create(file_path)

    if err != nil {
      fmt.Println("Error creating workflow: ", err)
      return
    }

    defer file.Close()

    encoder := json.NewEncoder(file)
    encoder.SetIndent("", "  ")
    encoder.Encode(wf)

    fmt.Printf("Workflow '%s' created at %s\n", workflow_name, file_path)
  },
}
