package services

import (
  "fmt"
  "io"
  "net/http"
  "os"
  "path/filepath"

  "github.com/AlexSTJO/cli-flow/internal/structures"
)

type HttpFetchService struct{}

func (s HttpFetchService) Name() string {
  return "http_fetch"
}

func (s HttpFetchService) Run(step structures.Step) (structures.Context, error) {
  urlRaw, ok := step.Config["url"]
  if !ok {
    return nil, fmt.Errorf(" Missing 'url' in config")
  }

  destRaw, ok := step.Config["destination"]
  if !ok {
    return nil, fmt.Errorf(" Missing 'destination' in config")
  }

  url, ok := urlRaw.(string)
  if !ok {
    return nil, fmt.Errorf(" 'url' must be a string")
  }

  dest, ok := destRaw.(string)
  if !ok {
    return nil, fmt.Errorf(" 'dest' must be a string")
  }

  fmt.Printf("[Download] Fetching: %s\n", url)

  resp, err := http.Get(url)
  if err != nil {
      return nil, fmt.Errorf("failed to GET URL: %w", err)
  }
  defer resp.Body.Close()

  if resp.StatusCode != 200 {
      return nil, fmt.Errorf("non-200 response: %d", resp.StatusCode)
  }

  if err := os.MkdirAll(filepath.Dir(dest), 0755); err != nil {
    return nil, fmt.Errorf("failed to create directories: %w", err)
  }

  outFile, err := os.Create(dest)
  if err != nil {
    return nil, fmt.Errorf(" Failed to create file: %w", err)
  }
  defer outFile.Close()

  _, err = io.Copy(outFile, resp.Body)
  if err != nil {
    return nil, fmt.Errorf(" Failed to write to file: %w", err)
  }

  fmt.Printf("[Download] Saved to: %s\n", dest)
  return structures.Context {
    "exit_code": 0,
    "status": "success",
  }, nil

}

func (s HttpFetchService) ConfigSpec() []string {
  return []string{"url", "destination"}
}

func init() {
  Registry["http_fetch"] = HttpFetchService{}
}
