package services

import (
  "fmt"
  "os"
  "os/exec"


  "github.com/AlexSTJO/cli-flow/intenral/structures"
)

type ShellService struct {
}


func (s ShellService) Name() string {
  return "shell"
}


func (s ShellService) Run(step structures.Step) error {
  cmdRaw, ok := step.Config["command"]
  if !ok {
    return fmt.Error("[Error] Missing 'command' field in step config")
  }

  cmdStr, ok := cmdRaw.(string)

  if (!ok){
    return fmt.Error("[Error] 'command' must be a non-empty string")
  }

  fmt.Println("[Shell] Shell execution initiated")
  fmt.Printf("[Shell] Executing the following command: %s\n", cmdStr)

  cmd := exec.Command("sh", "-c", cdmStr)

  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr
  cmd.Stdin = os.Stdin


  err := cmd.Run()

  if (err != nil) {
    return fmt.Error("[Error] Shell Error Occurred: %w", err)
  }
  return nil
}


func (s ShellService) ConfigSpec() []string {
  return []string{"command"}
}

func init() {
  Registry["shell"] = ShellService{}
}
