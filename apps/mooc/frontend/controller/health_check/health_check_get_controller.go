package healthcheck

import "github.com/gofiber/fiber/v2"

func GetControllerHandler(c *fiber.Ctx) error {
	status := struct {
		MoocFrontend string `json:"mooc_frontend"`
	}{
		MoocFrontend: "ok",
	}

	return c.JSON(status)
}
