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




var configureCmd = &cobra.Command{
  Use: "configure",
  Short: "Configuration command to store aws credentials",
  Run: func(cmd *cobra.Command, args []string) {
    reader := bufio.NewReader(os.Stdin)

    fmt.Print("Enter AWS Access Key ID: ")
    access_key, _ := reader.ReadString('\n')

    fmt.Print("Enter AWS Secret Key: ")
    secret_key, _ := reader.ReadString('\n')

    fmt.Print("Enter Region: ")
    region, _ := reader.ReadString('\n')
    

    config := structures.AWSConfig{
      AccessKey: strings.TrimSpace(access_key),
      SecretKey: strings.TrimSpace(secret_key),
      Region: strings.TrimSpace(region),
    }

    home_directory, _ := os.UserHomeDir()
    configPath := filepath.Join(home_directory, ".cli_flow")
    os.MkdirAll(configPath, os.ModePerm)

    file, err := os.Create(filepath.Join(configPath, "config.json"))

    if err != nil {
      fmt.Println("Error saving config: ", err)
      return
    }
    defer file.Close()

    encoder := json.NewEncoder(file)
    encoder.SetIndent("", "  ")
    encoder.Encode(config)

    fmt.Println("Config Saved Succesfully!")
  },
}
