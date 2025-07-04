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
			inners, err := parser.ParseSteps(step.Config["steps"])
			if err != nil {
				return false, fmt.Errorf("Could not parse loop steps")
			}
			exists, err := Search(inners, newName)
			if err != nil {
				return false, fmt.Errorf("Could not parse loop steps")
			}
			if exists {
				return true, nil
			}
    } else if step.Service == "if" {
			inners_t, err := parser.ParseSteps(step.Config["true_steps"])
			if err != nil {
				return false, fmt.Errorf("Could not parse if steps")
			}
			exists_t, err := Search(inners_t, newName)
			if err != nil {
				return false, fmt.Errorf("Could not parse if steps")
			}
			if exists_t {
				return true, nil
			}

			inners_f, err := parser.ParseSteps(step.Config["false_steps"])
			exists_f, err := Search(inners_f, newName)
			if err != nil {
				return false, fmt.Errorf("Could not parse if steps")
			}
			if exists_f {
				return true, nil
			}
  	}

	}
	return exists, nil
}



