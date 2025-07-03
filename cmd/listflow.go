package cmd


import (
  "fmt"
  "os"
  "path/filepath"
  "strings"
 	"github.com/fatih/color"
  "github.com/spf13/cobra"
)

var (
	cyan   = color.New(color.FgCyan).SprintFunc()
	green  = color.New(color.FgGreen).SprintFunc()
	yellow = color.New(color.FgYellow).SprintFunc()
	red    = color.New(color.FgRed).SprintFunc()
	bold   = color.New(color.Bold).SprintFunc()
	faint  = color.New(color.Faint).SprintFunc()
)


var listflowCmd = &cobra.Command {
  Use: "listflow",
  Short: "List all saved workflows",
  Run: func(cmd *cobra.Command, args []string) {
    home_directory, err := os.UserHomeDir()

    if err != nil {
      fmt.Println(red("[Error] Error reading workflow directory: "), err)
      return
    }
    dir := filepath.Join(home_directory, ".cli_flow", "workflows")

    files, err := os.ReadDir(dir)
    if err != nil {
      fmt.Println(red("[Error] Error reading workflow directory: "), err)
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
