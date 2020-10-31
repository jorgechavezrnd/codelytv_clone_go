package healthcheck

import "github.com/gofiber/fiber/v2"

func GetControllerHandler(c *fiber.Ctx) error {
	status := struct {
		MoocBackend string `json:"mooc_backend"`
	}{
		MoocBackend: "ok",
	}

	return c.JSON(status)
}
