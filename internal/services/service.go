package services

import "github.com/AlexSTJO/cli-flow/internal/structures"


// Interface acts like a contract remember that silly goose 
type Service interface {
  Run(step structures.Step) error
  Name() string
  ConfigSpec() []string
}


// Create a registry map so we can init into internal
var Registry = map[string]Service{}
