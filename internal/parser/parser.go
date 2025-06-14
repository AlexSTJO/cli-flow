package main

import (
	"regexp"
	"fmt"
)


func ParseExpression(expression string, context map[string]map[string]string){
	var varPattern = `\$\{([^}]+)\}`

	re := regexp.MustCompile(varPattern)
	updated := re.ReplaceAllFunc(///NEEDS INPUT))
	fmt.Println(updated)
}


func main(){
	context :=  map[string]map[string]string{}

	context["node11"] = map[string]string{
			"status": "success", "exit_code": "0",
	}

	
	ParseExpression("${bruh}", context)
}







