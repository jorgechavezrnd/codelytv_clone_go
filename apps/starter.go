package main

import (
	"log"
	"os"
	"strings"
	moocBackend "tv/codely/apps/mooc/backend"
	moocFrontend "tv/codely/apps/mooc/frontend"
	"tv/codely/src/shared/infrastructure/app"
	"tv/codely/src/shared/infrastructure/cli"
)

var applications map[string]app.Application
var commands map[string]map[string]cli.ConsoleCommand

func main() {
	args := os.Args[1:]

	if len(args) < 2 {
		log.Println("There are not enough arguments")
		os.Exit(1)
	}

	initializeApplications()

	applicationName := args[0]
	commandName := args[1]
	isServerCommand := strings.Compare("server", commandName) == 0

	ensureApplicationExist(applicationName)
	ensureCommandExist(applicationName, commandName)

	application := applications[applicationName]

	if isServerCommand {
		application.ConfigureServer()
		application.StartServer()
	} else {
		consoleCommand := commands[applicationName][commandName]

		consoleCommand.Execute()
	}
}

func initializeApplications() {
	registerApplications()
	registerCommands()
}

func registerApplications() {
	applications = make(map[string]app.Application)

	applications["mooc_backend"] = moocBackend.NewMoocBackendApplication()
	applications["mooc_frontend"] = moocFrontend.NewMoocFrontendApplication()
}

func registerCommands() {
	commands = make(map[string]map[string]cli.ConsoleCommand)

	commands["mooc_backend"] = applications["mooc_backend"].Commands()
	commands["mooc_frontend"] = applications["mooc_frontend"].Commands()
}

func ensureApplicationExist(applicationName string) {
	if _, exist := applications[applicationName]; !exist {
		log.Printf("The application <%s> doesn't exist. Valids:%s", applicationName, createApplicationsStringList(applications))
		os.Exit(1)
	}
}

func ensureCommandExist(applicationName string, commandName string) {
	if !(strings.Compare("server", commandName) == 0) && !existCommand(applicationName, commandName) {
		log.Printf("The command <%s> for application <%s> doesn't exist. Valids (application.command):\n- server%s", commandName, applicationName, createCommandsStringList(commands[applicationName]))
		os.Exit(1)
	}
}

func existCommand(applicationName string, commandName string) bool {
	_, appExist := applications[applicationName]

	if appExist {
		_, commandExist := commands[applicationName][commandName]

		return commandExist
	}

	return false
}

func createApplicationsStringList(mapWithValues map[string]app.Application) string {
	stringList := ""

	for k := range mapWithValues {
		stringList += "\n- " + k
	}

	return stringList
}

func createCommandsStringList(mapWithValues map[string]cli.ConsoleCommand) string {
	stringList := ""

	for k := range mapWithValues {
		stringList += "\n- " + k
	}

	return stringList
}
