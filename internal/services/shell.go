package services

import (
  "fmt"
  "os"
  "os/exec"


  "github.com/AlexSTJO/cli-flow/internal/structures"
)

type ShellService struct {
}


func (s ShellService) Name() string {
  return "shell"
}


func (s ShellService) Run(step structures.Step, ctx *structures.Context) ([]structures.Step , error) {
  cmdRaw, ok := step.Config["command"]
  if !ok {
    return nil, fmt.Errorf("[Error] Missing 'command' field in step config")
  }

  cmdStr, ok := cmdRaw.(string)

  if (!ok){
    return nil, fmt.Errorf("[Error] 'command' must be a non-empty string")
  }

  fmt.Println("[Shell] Shell execution initiated")
  fmt.Printf("[Shell] Executing the following command: %s\n", cmdStr)

  cmd := exec.Command("sh", "-c", cmdStr)

  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr
  cmd.Stdin = os.Stdin


  err := cmd.Run()

  if (err != nil) {
    return nil, fmt.Errorf("[Error] Shell Error Occurred: %v", err)
  }

  (*ctx)[step.Name] = map[string]any{
	"exit_code": 0,
    "status": "success",
  }

  return nil, nil
}

func (s ShellService) ConfigSpec() []string {
  return []string{"command"}
}

func init() {
  Registry["shell"] = ShellService{}
}
