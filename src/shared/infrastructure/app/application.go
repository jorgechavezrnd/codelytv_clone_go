package app

import "tv/codely/src/shared/infrastructure/cli"

type Application interface {
	ConfigureServer()
	StartServer()
	Commands() map[string]cli.ConsoleCommand
}
