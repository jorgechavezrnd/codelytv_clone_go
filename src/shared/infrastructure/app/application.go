package app

import "tv/codely/src/shared/infrastructure/cli"

type Application interface {
	StartServer()
	Commands() map[string]cli.ConsoleCommand
}
