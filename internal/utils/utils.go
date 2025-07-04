package utils


import (
	"fmt"


	"github.com/AlexSTJO/cli-flow/internal/parser"
	"github.com/AlexSTJO/cli-flow/internal/structures"
)




func Search(steps []structures.Step, newName string) (bool, error) {
	exists := false
	for _, step := range steps {
		if step.Name == newName {
			return true, nil
		}
		if step.Service == "loop" {
			inners, err := parser.ParseSteps(step.Config["true_steps"])
			if err != nil {
				return false, fmt.Errorf("Could not parse if steps")
			}
    	exists = Search(inners)
			if exists {
				return true, nil
			}
    } else if step.Service == "loop" {
			inners_t, err := parser.ParseSteps(step.Config["true_steps"])
			if err != nil {
				return false, fmt.Errorf("Could not parse loop steps")
			}
    	exists = Search(inners_t)
			if exists {
				return true, nil
			}

			inners_f, err := parser.ParseSteps(step.Config["false_steps"])
			exists = Search(inners_f)
			if exists {
				return true
			}
  	}

	}
	return exists
}



