package command

import (
	"fmt"
)

type FakeCommand struct{}

func NewFakeCommand() *FakeCommand {
	return &FakeCommand{}
}

func (fc *FakeCommand) Execute() {
	fmt.Println("Fake command executed from mooc backend")
}
