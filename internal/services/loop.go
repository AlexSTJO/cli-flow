package services

import (
  "fmt"
  "bufio"
  "os"
  "strings"

  "github.com/AlexSTJO/cli-flow/internal/structures"
)

type LoopService struct {}

func (s *LoopService) Name() string {
  return "loopservice"
}

func (s *LoopService) ConfigSpec() []string {
  return []string{}
}

func (s LoopService) PromptForConfig() (map[string]any, error) {
  reader := bufio.NewReader(os.Stdin)

  config := make(map[string]any)  
  fmt.Printf("Please choose one of the following \n 1. Specify Iterations\n 2. End Conditionally\n")
  fmt.Printf("Enter 1 or 2: ")
  mode, _ := reader.ReadString('\n')
  mode = strings.TrimSpace(mode)

  switch mode {
  case "1":
    fmt.Printf("How many iterations should it loop?: ")
    times, _ := reader.ReadString('\n')
    times = strings.TrimSpace(times)
    
    config["times"] = times
  
  case "2":
    fmt.Printf("Node ID to be checked: ")
    nodeId, _ := reader.ReadString('\n')
    nodeId = strings.TrimSpace(nodeId)

    fmt.Printf("Key of Returned structures.Context (i.e 'exit_code' or 'status': ")
    contextKey, _ := reader.ReadString('\n')
    contextKey = strings.TrimSpace(contextKey)

		fmt.Printf("Value of Returned structures.Context to End Loop (i.e '0' or 'success'): ")
    contextValue, _ := reader.ReadString('\n')
    contextValue = strings.TrimSpace(contextValue)

    config["node_id"] = nodeId
    config["context_key"] = contextKey
    config["context_value"] = contextValue

  default:
    return nil, fmt.Errorf("unknown loop mode")
  }

  config["steps"] = []any{}
  return config,nil
}

func (s *LoopService) Run(step structures.Step, ctx *structures.Context) ([]structures.Step, error) {
  config := step.Config

  rawSteps, ok := config["steps"].([]any)

  if !ok || len(rawSteps) == 0 {
    return nil, fmt.Errorf("|| Loop || Steps can not be empty in loop, add steps to loop by using 'addtoloop' command")
  }

  var steps []structures.Step

  for _, raw := range rawSteps {
    stepMap, ok := raw.(map[string]any)
    if !ok {
      return nil, fmt.Errorf("|| Loop || Invalid Step Format in Loop")
    }

    st := structures.Step{
      Name: stepMap["name"].(string),
      Service: stepMap["service"].(string),
      Config: stepMap["config"].(map[string]any),
    }

    steps = append(steps, st)
  }

  var loopCount int
  var maxIterations int = 10 
  var useCondition bool
  var nodeID, contextKey, contextValue string

  if raw, ok := config["times"]; ok {
    if str, ok := raw.(string); ok {
      _, err := fmt.Sscanf(str, "%d", &loopCount)
      if err != nil {
          return nil, fmt.Errorf("invalid integer for 'times': %v", err)
      }
    } else {
      return nil, fmt.Errorf("|| Loop || Invalid format for 'times' paramater")
    }
  } else if nid, hasNid := config["node_id"].(string); hasNid {
    	nodeID = nid
    	contextKey = config["context_key"].(string)
    	contextValue = config["context_value"].(string)
    	useCondition = true
  } else {
    return nil, fmt.Errorf("|| Loop || Requires 'times' or 'node_id' parameter")
  }

  var finalCtx structures.Context = map[string]any{}
  attempt := 0

  for {
    if loopCount > 0 && attempt >= loopCount {
      break
    }
    if attempt > maxIterations {
      return nil, fmt.Errorf("|| Loop || Exceeded max safe loop iterations (1000)")
    }
    

    attempt++
    fmt.Printf("|| Loop || Iteration #%d\n", attempt)

    stepCtx := &structures.Context{}
	 	queue := append([]structures.Step{}, steps...) 
    for len(queue) > 0 {
			c := queue[0]
			queue = queue[1:]
      svc, ok := Registry[c.Service]
      if !ok {
        return nil, fmt.Errorf("|| Loop || Unknown service '%s' in loop", c.Service)
      }

      nextSteps, err := svc.Run(c, stepCtx)

      if err != nil {
        return nil, fmt.Errorf("|| Loop || Error in step '%s' : '%w' ", c.Name, err)
      }

			queue = append(nextSteps, queue...)

      for k,v := range (*stepCtx) {
        finalCtx[k] = v
      }

    }

    if useCondition {
      targetCtxRaw, ok := (*stepCtx)[nodeID]
      if !ok {
        return nil, fmt.Errorf("|| Loop || Could not find context for step '%s'", nodeID)
      }

			targetCtx, ok := targetCtxRaw.(map[string]any)
			if !ok {
				return nil, fmt.Errorf("|| Loop || Type error with your target ctx: '%v'")
			}

      val, ok := targetCtx[contextKey]
			if !ok {
				return nil, fmt.Errorf("|| Loop || ContextKey not found in targetCtx: '%v'", contextKey)
			}
      if ok && fmt.Sprintf("%v", val) == contextValue {
        break
      }
    }
  }

  finalCtx["Attempts"] = attempt
	(*ctx)[step.Name] = finalCtx
  return nil,nil
}

func init(){
  Registry["loop"] = &LoopService{}
}
