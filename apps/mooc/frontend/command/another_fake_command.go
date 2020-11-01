package command

import (
	"fmt"
)

type AnotherFakeCommand struct{}

func NewAnotherFakeCommand() *AnotherFakeCommand {
	return &AnotherFakeCommand{}
}

func (fc *AnotherFakeCommand) Execute() {
	fmt.Println("Another fake command executed from mooc frontend")
}
