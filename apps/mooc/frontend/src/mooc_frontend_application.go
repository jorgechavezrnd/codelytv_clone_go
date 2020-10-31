package apps

import (
	"tv/codely/apps/mooc/frontend/src/command"
	healthcheck "tv/codely/apps/mooc/frontend/src/controller/health_check"
	"tv/codely/src/shared/infrastructure/cli"

	"github.com/gofiber/fiber/v2"
)

type MoocFrontendApplication struct{}

func NewMoocFrontendApplication() *MoocFrontendApplication {
	return &MoocFrontendApplication{}
}

func (m *MoocFrontendApplication) StartServer() {
	app := fiber.New()

	app.Get("/health-check", healthcheck.GetControllerHandler)

	_ = app.Listen(":4000")
}

func (m *MoocFrontendApplication) Commands() map[string]cli.ConsoleCommand {
	commands := make(map[string]cli.ConsoleCommand)

	commands["fake"] = command.NewFakeCommand()
	commands["another_fake"] = command.NewAnotherFakeCommand()

	return commands
}
