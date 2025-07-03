package runner

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/AlexSTJO/cli-flow/internal/services"
	"github.com/AlexSTJO/cli-flow/internal/structures"
)

var (
	cyan   = color.New(color.FgCyan).SprintFunc()
	green  = color.New(color.FgGreen).SprintFunc()
	yellow = color.New(color.FgYellow).SprintFunc()
	red    = color.New(color.FgRed).SprintFunc()
	bold   = color.New(color.Bold).SprintFunc()
	faint  = color.New(color.Faint).SprintFunc()
)

func RunWorkflow(wf structures.Workflow) error {
	queue := append([]structures.Step{}, wf.Steps...)
	ctx := &structures.Context{}

	stepCount := 0

	for len(queue) > 0 {
		stepCount++
		current := queue[0]
		queue = queue[1:]

		svc, ok := services.Registry[current.Service]
		if !ok {
			return fmt.Errorf("%s: %s", red("[Fatal] Unknown service"), cyan(current.Service))
		}

		fmt.Printf("\n%s %d: %s\n", bold("▶ Running Step"), stepCount, cyan(current.Name))
		fmt.Printf("   %s: %s\n", yellow("Service"), green(current.Service))

		nextSteps, err := svc.Run(current, ctx)
		if err != nil {
			return fmt.Errorf("%s in step '%s': %v", red("[Error] Task failed"), current.Name, err)
		}

		fmt.Printf("   %s\n", green("✓ Completed"))

		queue = append(nextSteps, queue...)
	}

	return nil
}

