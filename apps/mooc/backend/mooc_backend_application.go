package apps

import (
	"tv/codely/apps/mooc/backend/command"
	healthcheck "tv/codely/apps/mooc/backend/controller/health_check"
	"tv/codely/src/shared/infrastructure/cli"

	"github.com/gofiber/fiber/v2"
)

type MoocBackendApplication struct {
	app *fiber.App
}

func NewMoocBackendApplication() *MoocBackendApplication {
	newApp := &MoocBackendApplication{
		app: fiber.New(),
	}

	newApp.setupRoutes()

	return newApp
}

func (m *MoocBackendApplication) setupRoutes() {
	m.app.Get("/health-check", healthcheck.GetControllerHandler)
}

func (m *MoocBackendApplication) StartServer() {
	_ = m.app.Listen(":3000")
}

func (m *MoocBackendApplication) Commands() map[string]cli.ConsoleCommand {
	commands := make(map[string]cli.ConsoleCommand)

	commands["fake"] = command.NewFakeCommand()
	commands["another_fake"] = command.NewAnotherFakeCommand()

	return commands
}

func (m *MoocBackendApplication) App() *fiber.App {
	return m.app
}
