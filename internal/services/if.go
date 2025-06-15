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

func (s *IfService) Run(step structures.Step) (structures.Context, error) {
	expression := step.Config["statement"].(string)
	context := step.Config["__context"].(structures.Context)

	handledExpression := parser.ParseExpression(expression, context)

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

	return structures.Context {
		"bool": boolResult,
		"exit_code": "0",
		"status": "success",
	}, nil
}


func init(){
	Registry["if"] = &IfService{}
}
