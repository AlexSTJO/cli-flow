package services

import (
	"fmt"

	"github.com/AlexSTJO/cli-flow/internal/parser"
	"github.com/AlexSTJO/cli-flow/internal/structures"
	"github.com/Knetic/govaluate"
)


type IfService struct {}

func (s *IfService) Name() string {
	return "if"
}

func (s *IfService) ConfigSpec() []string {
	return []string{"statement"}
}

func (s *IfService) Run(step structures.Step, ctx *structures.Context) ([]structures.Step, error) {
	expression := step.Config["statement"].(string)
	handledExpression := parser.ParseExpression(expression, (*ctx))

	fmt.Printf("Evaluating Expression: %s\n", handledExpression) 

	gvExpression, err := govaluate.NewEvaluableExpression(handledExpression)
	if err != nil {
		return nil, err
	}

	result, err := gvExpression.Evaluate(nil)
	if err != nil {
		return nil, err
	}

	boolResult, ok := result.(bool)
	if !ok {
		return nil, fmt.Errorf("Expression did not return a boolean")
	}
	fmt.Printf("Evaluated to: %t\n",boolResult)
	
	(*ctx)[step.Name] = map[string]any{
		"bool": boolResult,
		"exit_code": "0",
		"status": "success",
	}


	return nil, nil
}


func init(){
	Registry["if"] = &IfService{}
}
