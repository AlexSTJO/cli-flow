package config

import (
	"fmt"
	"path/filepath"
	"encoding/json"
	"os"

	"github.com/AlexSTJO/cli-flow/internal/structures"

)

func HandleSmtpConfig() error {
	fmt.Println("Loading SMTP Config into Environment Variables")

	home_directory, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("Error occurred loading home directory: %w", err)
	}

	configPath := filepath.Join(home_directory, ".cli_flow", "config_smtp.json")
	data, err := os.ReadFile(configPath)

	var cfg structures.SMTPConfig

	if err := json.Unmarshal(data, &cfg); err != nil {
		return fmt.Errorf("Error occurred parsing json: %w", err)
	}
	
	os.Setenv("SMTPEmailAddress", cfg.EmailAddress)
	os.Setenv("SMTPEmailPassword", cfg.EmailPassword)

	return nil
	
}

func UnsetSmtpEnv() {
	os.Unsetenv("SMTPEmailAddress")
	os.Unsetenv("SMTPEmailAddress")
}
