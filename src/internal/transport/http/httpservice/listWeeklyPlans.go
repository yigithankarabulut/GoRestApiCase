package httpservice

import (
	"context"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/yigithankarabulut/vatansoftgocase/src/internal/service/planService"
	"strconv"
)

func (s *httpServiceHandler) ListWeeklyPlans(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), s.Handler.CancelTimeout)
	defer cancel()

	var listWeeklyRequest planService.ListWeeklyPlansRequest
	var err error

	listWeeklyRequest.LastWeek, err = strconv.Atoi(c.Query("lastweek"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "last week is not valid",
		})
	}
	studentNumber, err := s.JwtGetStudentNumber(c.Cookies("jwt"))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "jwt token is not valid",
		})
	}
	listWeeklyRequest.StudentNumber = studentNumber
	result, err := s.planService.ListWeekly(ctx, &listWeeklyRequest)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return c.Status(fiber.StatusGatewayTimeout).JSON(fiber.Map{
				"error": "context deadline exceeded",
			})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error": "plans not found",
		})
	} else if result == nil || len(*result) < 1 {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error": "plans not found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "plans listed",
		"plans":   result,
	})
}
