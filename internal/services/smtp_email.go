package services

import (
  "os"
  "net/smtp"
  "fmt"


	"github.com/AlexSTJO/cli-flow/internal/structures"
  "github.com/AlexSTJO/cli-flow/internal/config"
)

type SmtpEmailService struct{}

func (s SmtpEmailService) Name() string {
	return "smtp_email"
}

func (s SmtpEmailService) ConfigSpec() []string {
	return []string{"destination_email", "subject", "body", "port", "host"}
}

func (s SmtpEmailService) Run(step structures.Step) (structures.Context, error) {
  if err := config.HandleSmtpConfig(); err != nil{
    return nil, fmt.Errorf("Error getting smtp config: %w", err)
  }
  from := os.Getenv("SMTPEmailAddress")
  pwd := os.Getenv("SMTPEmailPassword")
  host, ok := step.Config["host"].(string)
  if !ok {
    return nil, fmt.Errorf("host must be a string")
  }
  port, ok := step.Config["port"].(string)
  if !ok {
    return nil, fmt.Errorf("port must be a string")
  }

  subject, ok := step.Config["subject"].(string)
  if !ok {
    return nil, fmt.Errorf("subject must be a string")
  }

  body, ok := step.Config["body"].(string)
  if !ok {
    return nil, fmt.Errorf("body must be string")
  }

  destinationEmail, ok := step.Config["destination_email"].(string)
  if !ok {
    fmt.Errorf("Destination Email must be string")
  }

  auth := smtp.PlainAuth("", from, pwd, host)

  msg := []byte(
    "To: " + destinationEmail + "\r\n" +
    "Subject: " + subject + "\r\n" +
    "\r\n" + 
    body + "\r\n")

  
  err := smtp.SendMail(host+":"+port, auth, from, []string{destinationEmail}, msg)
  if err != nil{
    return nil, fmt.Errorf("failed to send email: %w", err)
  }

  fmt.Println("Email has been sent!")

  return structures.Context{ "exit_code": 0, "status": "success"}, nil
}


func init(){
  Registry["smtp_email"] = &SmtpEmailService{}
}
