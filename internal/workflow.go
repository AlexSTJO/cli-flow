package workflow

type Step struct {
  Service string `json:"service"`
  Name string `json:"name"`
  Command string `json:"command"`
}

type Workflow struct {
  Name string `json:"name"`
  Description string `json:"description"`
  Steps []Step `json:"steps`
}
