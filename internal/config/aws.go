package config

import (
  "encoding/json"
  "fmt"
  "os"
  "path/filepath"

  "github.com/AlexSTJO/cli-flow/internal/structures"
)

func LoadAWSConfig() (structures.AWSConfig, error){
  var cfg structures.AWSConfig
  fmt.Println("[Config] Checking for AWS Credentials")

  home, err := os.UserHomeDir()
  if err != nil {
    return cfg,err
  }

  path := filepath.Join(home, ".cli_flow", "config_aws.json")

  data, err := os.ReadFile(path)

  if err != nil {
    return cfg, fmt.Errorf("missing config.json; please run `cli-flow configure-aws` first: %w", err)
 }

  if err := json.Unmarshal(data, &cfg); err!=nil{
    return cfg, fmt.Errorf("invalid config.json: %w", err)
  }

  return cfg, nil
}

func SetAWSEnvVars(cfg structures.AWSConfig) {
	fmt.Println("[Config] AWS Env Vars Being Set")
	os.Setenv("AWS_ACCESS_KEY_ID", cfg.AccessKey)
	os.Setenv("AWS_SECRET_ACCESS_KEY", cfg.SecretKey)
	os.Setenv("AWS_REGION", cfg.Region)
}


func UnsetAWSEnvVars() {
	fmt.Println("[Config] AWS Env Vars Being Unset")
	os.Unsetenv("AWS_ACCESS_KEY_ID")
	os.Unsetenv("AWS_SECRET_ACCESS_KEY")
	os.Unsetenv("AWS_REGION")
}
