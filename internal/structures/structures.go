package structures 

type Step struct {
	Name    string                 `json:"name"`
	Service string                 `json:"service"`
	Config  map[string]interface{} `json:"config,omitempty"` 
}

type Workflow struct {
  Name string `json:"name"`
  Description string `json:"description"`
  Steps []Step `json:"steps"`
}

type AWSConfig struct {
  AccessKey string `json:"aws_access_key_id"`
  SecretKey string `json:"aws_secret_access_key"`
  Region string `json:"aws_region"`
}

type SMTPConfig struct {
	EmailAddress string `json:"email_address"`
	EmailPassword string `json:"email_password"`
}


