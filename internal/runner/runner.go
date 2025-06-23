package runner

import (
	"fmt"

	"github.com/AlexSTJO/cli-flow/internal/services"
	"github.com/AlexSTJO/cli-flow/internal/structures"
)


func RunWorkflow(wf structures.Workflow) error{
	steps := wf.Steps
	queue := append([]structures.Step{}, steps...)
	
	ctx := &structures.Context{}

	for len(queue) > 0 {
		c := queue[0]
		queue = queue[1:]
		
		svc, ok := services.Registry[c.Service]
		if !ok {
			return fmt.Errorf("Service Name not recognized")
		}

		nextSteps, err := svc.Run(c, ctx)
		if err != nil {
			return fmt.Errorf("Error while running task", err)
		}

		queue = append(nextSteps, queue...)
	}

	return nil

}
