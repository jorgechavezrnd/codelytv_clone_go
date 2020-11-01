package apps

import (
	"tv/codely/apps/mooc/backend/command"
	healthcheck "tv/codely/apps/mooc/backend/controller/health_check"
	"tv/codely/src/shared/infrastructure/cli"

	"github.com/gofiber/fiber/v2"
)

type MoocBackendApplication struct {
	*fiber.App
}

func NewMoocBackendApplication() *MoocBackendApplication {
	return &MoocBackendApplication{}
}

func (m *MoocBackendApplication) ConfigureServer() {
	m.App = fiber.New()

	m.setupRoutes()
}

func (m *MoocBackendApplication) StartServer() {
	_ = m.Listen(":3000")
}

func (m *MoocBackendApplication) Commands() map[string]cli.ConsoleCommand {
	commands := make(map[string]cli.ConsoleCommand)

	commands["fake"] = command.NewFakeCommand()
	commands["another_fake"] = command.NewAnotherFakeCommand()

	return commands
}

func (m *MoocBackendApplication) setupRoutes() {
	m.Get("/health-check", healthcheck.GetControllerHandler)
}
