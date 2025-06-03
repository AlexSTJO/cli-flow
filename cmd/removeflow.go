package cmd

import (
  "fmt"
  "os"
  "path/filepath"
  "bufio"
  "strings"

  "github.com/spf13/cobra"
)

var removeflowCmd = &cobra.Command {
  Use: "removeflow [workflow_name]",
  Short: "Removes saved workflow name",
  Args: cobra.ExactArgs(1),
  Run: func(cmd *cobra.Command, args[]string) {
    workflow_name := args[0]

    home, err := os.UserHomeDir()
    
    if err != nil {
      fmt.Println("[Error] Error getting home directory: ", err)
    }
    path := filepath.Join(home, ".cli_flow", "workflows", workflow_name+".json")

    reader := bufio.NewReader(os.Stdin)

    fmt.Printf("Will delete '%s' || Please confirm by reentering workflow name: ", workflow_name)
    confirm, _ := reader.ReadString('\n')
    confirm = strings.TrimSpace(confirm)


    if confirm != workflow_name {
      fmt.Println("[Error] Input does not match with workflow name")
      return
    }

    err = os.Remove(path)
    
    if err != nil {
      fmt.Println("[Error] Error deleting workflow, maybe check workflow name?")
      return
    }

    fmt.Println("[Sucess] Workflow deleted")
    return

  },
}
