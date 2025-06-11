package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"path/filepath"
	"strings"
	"os"

	"github.com/spf13/cobra"
	"github.com/AlexSTJO/cli-flow/internal/structures"

)


var configuresmtpCmd = &cobra.Command{
	Use: "config-smtp",
	Short: "Configuration command to store smtp email credentials",
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Please enter email address: ")
		email_address, _ := reader.ReadString('\n')

		fmt.Print("Please enter email password: ")
		password, _ := reader.ReadString('\n')

		config := structures.SMTPConfig {
			EmailAddress: strings.TrimSpace(email_address),
			EmailPassword: strings.TrimSpace(password),
		}

		home_directory, _ := os.UserHomeDir()
		configPath := filepath.Join(home_directory, ".cli_flow")
		os.MkdirAll(configPath, os.ModePerm)

		file, err := os.Create(filepath.Join(configPath, "config_smtp.json"))
		if err != nil {
			fmt.Println("Error saving config: ", err)
			return
		}
		defer file.Close()

		encoder := json.NewEncoder(file)
		encoder.SetIndent("", "  ")
		encoder.Encode(config)

		fmt.Println("Config Saved Succesfully")

	},
}
