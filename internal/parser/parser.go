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

