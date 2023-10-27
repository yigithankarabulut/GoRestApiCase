package httpservice

import (
	"context"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/yigithankarabulut/vatansoftgocase/src/internal/service/planService"
)

func (s *httpServiceHandler) ListAllPlans(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), s.Handler.CancelTimeout)
	defer cancel()

	studentNumber, err := s.JwtGetStudentNumber(c.Cookies("jwt"))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "jwt token is not valid",
		})
	}
	result, err := s.planService.ListAll(ctx, &planService.ListAllPlansRequest{StudentNumber: studentNumber})
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return c.Status(fiber.StatusGatewayTimeout).JSON(fiber.Map{
				"error": "context deadline exceeded",
			})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error": "plans not found",
		})
	} else if result == nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error": "plans not found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "plans listed",
		"plans":   result,
	})
}
