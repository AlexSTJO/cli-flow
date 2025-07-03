package formatter


import (
	"fmt"

	"github.com/AlexSTJO/cli-flow/internal/structures"
	"github.com/AlexSTJO/cli-flow/internal/parser"
)


func printStepTree(prefix string, step structures.Step, isLast bool) {
	connector := "├──"
	nextPrefix := prefix + "│   "
	if isLast {
		connector = "└──"
		nextPrefix = prefix + "    "
	}

	fmt.Printf("%s%s %s\n", prefix, connector, step.Name)
	fmt.Printf("%s    └── Service: %s\n", prefix, step.Service)

	switch step.Service {
	case "loop":
		inners, err := parser.ParseSteps(step.Config["steps"])
		if err != nil || len(inners) == 0 {
			fmt.Printf("%s    (empty loop)\n", nextPrefix)
			return
		}
		for i, inner := range inners {
			printStepTree(nextPrefix, inner, i == len(inners)-1)
		}
	case "if":
		trueSteps, _ := parser.ParseSteps(step.Config["true_steps"])
		falseSteps, _ := parser.ParseSteps(step.Config["false_steps"])

		fmt.Printf("%s    ├── <True>\n", prefix)
		if len(trueSteps) == 0 {
			fmt.Printf("%s    │   └── (empty)\n", prefix)
		} else {
			for i, inner := range trueSteps {
				printStepTree(prefix+"    │   ", inner, i == len(trueSteps)-1)
			}
		}

		fmt.Printf("%s    └── <False>\n", prefix)
		if len(falseSteps) == 0 {
			fmt.Printf("%s        └── (empty)\n", prefix)
		} else {
			for i, inner := range falseSteps {
				printStepTree(prefix+"        ", inner, i == len(falseSteps)-1)
			}
		}
	}
}
