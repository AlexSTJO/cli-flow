package parser

import (
	"regexp"
	"fmt"
	"strings"


	"github.com/AlexSTJO/cli-flow/internal/structures"
)


func ParseExpression(expression string, context structures.Context) string {
	var varPattern = `\$\{([^}]+)\}`

	re := regexp.MustCompile(varPattern)
	updated := re.ReplaceAllStringFunc(expression , func (m string) string {
		inner := re.FindStringSubmatch(m)[1]

		parts := strings.SplitN(inner, ".", 2)
		if len(parts) != 2 {
			return m
		}

		nodeKey, fieldkey := parts[0], parts[1]


		nodeRaw, ok := context[nodeKey]
		if !ok {
			return m
		}

		nodeMap, ok := nodeRaw.(map[string]any)
		if !ok {
			return m
		}

		valRaw, ok := nodeMap[fieldkey]
		if !ok {
			return m
		}

		valStr, ok := valRaw.(string)
		if !ok {
			return m
		}
		return fmt.Sprintf("%q", valStr)
	})
	fmt.Printf("[Parser] Returning parsed string: %s\n", updated)
	return updated
}


func ParseSteps (raw any) ([]structures.Step, error) {
	items, ok := raw.([]interface{})
	if !ok {
		return nil, fmt.Errorf("Invalid input in Step Parser")
	}

	var steps []structures.Step

	for _, item := range items {
		stepMap, ok := item.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("Invalid input in Step Parser")
		}

		step := structures.Step{
			Name: stepMap["name"].(string),
			Service: stepMap["service"].(string),
			Config: stepMap["config"].(map[string]interface{}),
		}
		steps = append(steps,step)
	}

	return steps, nil
}
