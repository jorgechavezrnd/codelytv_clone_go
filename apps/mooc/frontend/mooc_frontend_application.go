package apps

import (
	"tv/codely/apps/mooc/frontend/command"
	healthcheck "tv/codely/apps/mooc/frontend/controller/health_check"
	"tv/codely/src/shared/infrastructure/cli"

	"github.com/gofiber/fiber/v2"
)

type MoocFrontendApplication struct {
	*fiber.App
}

func NewMoocFrontendApplication() *MoocFrontendApplication {
	return &MoocFrontendApplication{}
}

func (m *MoocFrontendApplication) ConfigureServer() {
	m.App = fiber.New()

	m.setupRoutes()
}

func (m *MoocFrontendApplication) StartServer() {
	_ = m.Listen(":4000")
}

func (m *MoocFrontendApplication) Commands() map[string]cli.ConsoleCommand {
	commands := make(map[string]cli.ConsoleCommand)

	commands["fake"] = command.NewFakeCommand()
	commands["another_fake"] = command.NewAnotherFakeCommand()

	return commands
}

func (m *MoocFrontendApplication) setupRoutes() {
	m.Get("/health-check", healthcheck.GetControllerHandler)
}
