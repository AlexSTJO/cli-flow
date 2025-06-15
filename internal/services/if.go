package services

import (
	"fmt"

	"github.com/AlexSTJO/cli-flow/internal/structures"
	"github.com/AlexSTJO/cli-flow/internal/parser"
)


type IfService struct {}

func (s *IfService) Name() string {
	return "if"
}

func (s *IfService) ConfigSpec() []string {
	return []string{"statement"}
}

func (s *IfService) Run(step structures.Step) (structures.Context, error) {
	expression := step.Config["statement"].(string)
	context := step.Config["__context"].(structures.Context)

	handledExpression := parser.ParseExpression(expression, context)

	fmt.Printf("Expression is: %s\n", handledExpression) 
	return nil, nil
}


func init(){
	Registry["if"] = &IfService{}
}
