package http_v1

import (
	"github.com/gofiber/fiber/v2"
	"os"
	"time"
)

func (s *server) Home(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json")
	hostname, err := os.Hostname()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			&fiber.Map{
				"message": "Server: error getting hostname",
				"error":   err.Error(),
			},
		)
	}

	return c.Status(fiber.StatusOK).JSON(
		&fiber.Map{
			"message":     "Information about the server",
			"serviceName": "Docker Swarm Auto-Scale Service",
			"hostname":    hostname,
			"timestamp":   time.Now().String(),
			"version":     "1.0.0",
			"error":       "No errors",
		},
	)
}

func (s *server) Ping(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json")

	return c.Status(fiber.StatusOK).JSON(
		&fiber.Map{
			"message": "Pong!",
		},
	)
}
