package services

import (
  "fmt"
  "io"
  "net/http"
  "os"

  "github.com/AlexSTJO/cli-flow/internal/structures"
)

type HttpFetchService struct{}

func (s HttpFetchService) Run(step structures.Step) error {
  urlRaw, ok := step.Config["url"]
  if !ok {
    return fmt.Errorf("[Error] Missing 'url' in config")
  }

  destRaw, ok := step.Config["destination"]
  if !ok {
    return fmt.Errorf("[Error] Missing 'destination' in config")
  }

  url, ok := urlRaw.(string)
  if !ok {
    return fmt.Errorf("[Error] 'url' must be a string")
  }

  dest, ok := destRaw.(string)
  if !ok {
    return fmt.Errorf("[Error] 'dest' must be a string")
  }

  fmt.Printf("[Download] Fetching: %s\n", url)

  resp, err := http.Get(url)
  if err != nil {
    return fmt.Errorf("[Error] Non-200 response: %d", resp.StatusCode)
  }

  outFile, err := os.Create(dest)
  if err != nil {
    return fmt.Errorf("[Error] Failed to create file: %w", err)
  }
  def outFile.Close()

  _, err = io.Copy(outFile, resp.Body)
  if err != nil {
    return fmt.Errorf("[Error] Failed to write to file: %w", err)
  }

  fmt.Printf("[Download] Saved to: %s\n", dest)
  return nil

}

func (s HttpFetchService) ConfigSpec() []string {
  return []string{"url", "destination"}
}

func init() {
  Registry["http_fetch"] = HttpFetchService{}
}
