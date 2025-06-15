package main

import (
	"regexp"
	"fmt"
	"strings"
)


func ParseExpression(expression string, context map[string]map[string]string) string {
	var varPattern = `\$\{([^}]+)\}`

	re := regexp.MustCompile(varPattern)
	updated := re.ReplaceAllStringFunc(expression , func (m string) string {
		inner := re.FindStringSubmatch(m)[1]

		parts := strings.SplitN(inner, ".", 2)
		if len(parts) != 2 {
			return m
		}

		node, key := parts[0], parts[1]
		if nodeData, ok := context[node]; ok {
			if val, ok := nodeData[key]; ok {
				return val
			}
		}

		return m
	})
	fmt.Printf("[Parser] Returning parsed string: %s", updated)
	return updated
}

/* tester
func main(){
	context :=  map[string]map[string]string{}

	context["node1"] = map[string]string{
			"status": "success", "exit_code": "0",
	}

	
	ParseExpression("${node1.status}", context)
}

*/





