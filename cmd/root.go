/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
  "bufio"
  "encoding/json"
  "path/filepath"
  "strings"
	"os"
  "fmt"
  "github.com/AlexSTJO/cli-flow/internal"
	"github.com/spf13/cobra"
)

type AWSConfig struct {
  AccessKey string `json:"aws_access_key_id`
  SecretKey string `json:"aws_secret_access_key`
  Region string `json:aws_region`
}

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
    

    config := AWSConfig{
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


var createflowCmd = &cobra.Command {
  Use: "createFlow",
  Short: "Create a new workflow",
  Run: func(cmd *cobra.Command, args []string) {
    reader := bufio.NewReader(os.Stdin)

    fmt.Print("Workflow Name: ")
    name, _ := reader.ReadString('\n')

    fmt.Print("Description: ")
    description, _ := reader.ReadString('\n')

    workflow_name := strings.TrimSpace(name)
    workflow_description := strings.TrimSpace(description)

    wf := workflow.Workflow{
      Name: workflow_name,
      Description: workflow_description,
      Steps: []workflow.Step{},
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

var rootCmd = &cobra.Command{
	Use:   "cli-flow",
	Short: "The CLI Workflow Automation Tool",
	Long: `Some type of super long description
  `,
	Run: func(cmd *cobra.Command, args []string) { 
    fmt.Println("Welcome to CLI Flow!")
  },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
    fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
  rootCmd.AddCommand(configureCmd)
  rootCmd.AddCommand(createflowCmd)
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cli-flow.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


