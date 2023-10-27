package httpservice

import (
	"context"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/yigithankarabulut/vatansoftgocase/src/internal/service/planService"
	"time"
)

func (s *httpServiceHandler) ListMonthlyPlans(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), s.Handler.CancelTimeout)
	defer cancel()

	var listMonthlyRequest planService.ListMonthlyPlansRequest
	if err := c.BodyParser(&listMonthlyRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "request body could not be parsed",
		})
	}
	if listMonthlyRequest.Month <= 0 || listMonthlyRequest.Month > 12 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "month is not valid",
		})
	}
	if listMonthlyRequest.Year > time.Now().Year()+20 || listMonthlyRequest.Year < 1970 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "year is not valid",
		})
	}
	studentNumber, err := s.JwtGetStudentNumber(c.Cookies("jwt"))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "jwt token is not valid",
		})
	}
	listMonthlyRequest.StudentNumber = studentNumber
	result, err := s.planService.ListMonthly(ctx, &listMonthlyRequest)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return c.Status(fiber.StatusGatewayTimeout).JSON(fiber.Map{
				"error": "context deadline exceeded",
			})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error": "plans not found",
		})
	}
	if result == nil || len(*result) < 1 {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error": "plans not found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "plans listed",
		"plans":   result,
	})
}
