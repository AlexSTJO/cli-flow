package services

import (
	"fmt"
	"bufio"
	"os"
	"strings"

	"github.com/AlexSTJO/cli-flow/internal/structures"
)


type IfService struct {}

func (s *IfService) Name() string {
	return "if"
}

func (s *IfService) ConfigSpec() []string {
	return []string{"Name", "Statement"}
}

func (s *IfService) Run(step structures.Step) (Context, error) {

}
