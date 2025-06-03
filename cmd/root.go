/*
Copyright © 2025 Alexandros St John alexandros.georgakoudi@gmail.com
*/
package cmd

import (
	"github.com/spf13/cobra"
  "fmt"
  "os"
)




var rootCmd = &cobra.Command{
	Use:   "cli-flow",
	Short: "The CLI Workflow Automation Tool",
	Long: `Some type of super long description
  `,
	Run: func(cmd *cobra.Command, args []string) { 
    fmt.Println("Welcome to CLI Flow!")
  },
}

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
  rootCmd.AddCommand(addstepCmd)
  rootCmd.AddCommand(runflowCmd)
  rootCmd.AddCommand(listflowCmd) 
  rootCmd.AddCommand(liststepsCmd)
  rootCmd.AddCommand(removeflowCmd)
  rootCmd.AddCommand(removestepCmd)
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


