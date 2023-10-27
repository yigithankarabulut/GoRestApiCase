package httpservice

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"time"
)

func (s *httpServiceHandler) LogoutStudent(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), s.Handler.CancelTimeout)
	defer cancel()
	select {
	case <-ctx.Done():
		return c.Status(fiber.StatusGatewayTimeout).JSON(fiber.Map{
			"error": "context deadline exceeded",
		})
	default:
		c.Cookie(&fiber.Cookie{
			Name:     "jwt",
			Value:    "",
			Expires:  time.Now().Add(-time.Hour),
			HTTPOnly: true,
		})
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "logout successful",
		})
	}
}
