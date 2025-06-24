package services

import (
	"fmt"
	"bufio"
	"strings"
	"os"
	
	"github.com/AlexSTJO/cli-flow/internal/parser"
	"github.com/AlexSTJO/cli-flow/internal/structures"
	"github.com/Knetic/govaluate"
)


type IfService struct {}

func (s *IfService) Name() string {
	return "if"
}

func (s *IfService) ConfigSpec() []string {
	return []string{}
}

func (s IfService) PromptForConfig() (map[string]any, error){
	reader := bufio.NewReader(os.Stdin)

	config := map[string]any{}

	fmt.Printf("Please enter a conditional statement (i.e ${node.status} == \"sucess\"): ")
	statement, _ := reader.ReadString('\n')
	statement = strings.TrimSpace(statement)

	config["statement"] = statement
	config["true_steps"] = []any{}
	config["false_steps"] = []any{}

	return config, nil
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
	fmt.Printf("Evaluated to: %t -> Appending Tasks to runtime queue....\n",boolResult)
	(*ctx)[step.Name] = map[string]any{
		"bool": boolResult,
		"exit_code": "0",
		"status": "success",
	}
	

	var rawSteps []interface{}

	
	if (boolResult) {
		rawSteps,ok = step.Config["true_steps"].([]interface{})
		if !ok {
			return nil, fmt.Errorf("|| IfService || if_true must be a list of steps")
		}

	} else {
		rawSteps,ok = step.Config["false_steps"].([]interface{})
		if !ok {
			return nil, fmt.Errorf("|| IfService || if_true must be a list of steps")
		}

	}

		
	var steps []structures.Step

	for _, raw := range rawSteps {
		stepMap, ok := raw.(map[string]interface{})
		if !ok {
    	return nil, fmt.Errorf("|| IfService || invalid step format in if_true")
    }

		st := structures.Step{
        Name:    stepMap["name"].(string),
        Service: stepMap["service"].(string),
        Config:  stepMap["config"].(map[string]interface{}),
    }

    steps = append(steps, st)
	}

	return steps,nil
}



func init(){
	Registry["if"] = &IfService{}
}
