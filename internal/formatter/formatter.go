package formatter

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/AlexSTJO/cli-flow/internal/parser"
	"github.com/AlexSTJO/cli-flow/internal/structures"
)

func PrintStepTree(prefix string, step structures.Step, isLast bool) {
	connector := "├──"
	nextPrefix := prefix + "│   "
	if isLast {
		connector = "└──"
		nextPrefix = prefix + "    "
	}

	fmt.Printf("%s%s %s\n", prefix, connector, color.New(color.FgCyan).Sprint(step.Name))
	fmt.Printf("%s    └── Service: %s\n", prefix, color.New(color.FgGreen).Sprint(step.Service))

	switch step.Service {
	case "loop":
		inners, err := parser.ParseSteps(step.Config["steps"])
		if err != nil || len(inners) == 0 {
			fmt.Printf("%s    └── %s\n", prefix, color.New(color.Faint).Sprint("(empty loop)"))
			return
		}
		for i, inner := range inners {
			PrintStepTree(nextPrefix, inner, i == len(inners)-1)
		}

	case "if":
		trueSteps, _ := parser.ParseSteps(step.Config["true_steps"])
		falseSteps, _ := parser.ParseSteps(step.Config["false_steps"])

		fmt.Printf("%s├── %s\n", nextPrefix, color.New(color.FgYellow).Sprint("True"))
		if len(trueSteps) == 0 {
			fmt.Printf("%s│   └── %s\n", nextPrefix, color.New(color.Faint).Sprint("(empty)"))
		} else {
			for i, inner := range trueSteps {
				PrintStepTree(nextPrefix+"│   ", inner, i == len(trueSteps)-1)
			}
		}

		fmt.Printf("%s└── %s\n", nextPrefix, color.New(color.FgYellow).Sprint("False"))
		if len(falseSteps) == 0 {
			fmt.Printf("%s    └── %s\n", nextPrefix, color.New(color.Faint).Sprint("(empty)"))
		} else {
			for i, inner := range falseSteps {
				PrintStepTree(nextPrefix+"    ", inner, i == len(falseSteps)-1)
			}
		}
	}
}

