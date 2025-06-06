package services

import (
  "fmt"
  "os"
  "os/exec"

  "github.com/AlexSTJO/cli-flow/internal/structures"
)


type S3Service struct {}

func (s *S3Service) Name() string {
  return "s3"
}

func (s *S3Service) ConfigSpec() []string {
  return []string{"action", "bucket", "key", "path"}
}

func (s *S3Service) Run(step structures.Step) error {
  action:=step.Config["action"].(string)
  bucket:=step.Config["bucket"].(string)
  key:=step.Config["key"].(string)
  path:=step.Config["path"].(string)

  if action != "upload" && action != "download" {
    return fmt.Errorf("Invalid action: %s (must be 'upload' or 'download')", action)
  }

  var cmd *exec.Cmd
  s3Uri := fmt.Sprintf("s3://%s/%s", bucket, key)

  if action == "upload" {
    _, err := os.Stat(path)

    if err != nil {
      return fmt.Errorf("Upload failed: local file %s does not exist", path)
    }
    cmd = exec.Command("aws", "s3", "cp", path, s3Uri)
  } else {
    cmd = exec.Command("aws", "s3", "cp", path, s3Uri)
  }

  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr

  fmt.Printf("[s3] %s from %s to %s\n", action, s3Uri, path)
  return cmd.Run()
}

func init(){
  Registry["s3"] = &S3Service{}
}

