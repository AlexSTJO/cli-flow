package services

import (

	"github.com/AlexSTJO/cli-flow/internal/structures"
)

type SmtpEmailService struct{}

func (s SmtpEmailService) Name() string {
	return "smtp_email"
}

func (s SmtpEmailService) ConfigSpec() []string {
	return []string{"destination_email", "content", "port", "host"}
}

func (s SmtpEmailService) Run(structures.Step) (error) {
	return nil

}
