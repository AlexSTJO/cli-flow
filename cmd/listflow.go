package cmd


import (
  "fmt"
  "os"
  "path/filepath"
  "strings"
  
  "github.com/spf13/cobra"
)


var listflowCmd = &cobra.Command {
  Use: "listflow",
  Short: "List all saved workflows",
  Run: func(cmd *cobra.Command, args []string) {
    home_directory, err := os.UserHomeDir()

    if err != nil {
      fmt.Println("[Error] Error reading workflow directory: ", err)
      return
    }
    dir := filepath.Join(home_directory, ".cli_flow", "workflows")

    files, err := os.ReadDir(dir)
    if err != nil {
      fmt.Println("[Error] Error reading workflow directory: ", err)
      return
    }


    fmt.Println("Available workflows:")
    for _, file := range files {
      if !file.IsDir() && strings.HasSuffix(file.Name(), ".json") {
        name := strings.TrimSuffix(file.Name(), ".json")
        fmt.Println("- " + name)
      } 
    }
    

  },
}
