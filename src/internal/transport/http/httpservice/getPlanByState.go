package httpservice

import (
	"context"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/yigithankarabulut/vatansoftgocase/src/internal/service/planService"
)

func (s *httpServiceHandler) GetPlanByState(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), s.Handler.CancelTimeout)
	defer cancel()

	planState := c.Query("state")
	if planState == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "plan state is empty",
		})
	}
	var validPlanStates = []string{"continue", "completed", "cancelled", "created"}
	switch planState {
	case validPlanStates[0]:
		planState = "continue"
	case validPlanStates[1]:
		planState = "completed"
	case validPlanStates[2]:
		planState = "cancelled"
	case validPlanStates[3]:
		planState = "created"
	default:
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "plan state is not valid",
		})
	}
	studentNumber, err := s.JwtGetStudentNumber(c.Cookies("jwt"))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "jwt token is not valid",
		})
	}
	result, err := s.planService.GetByState(ctx, &planService.GetPlanByStateRequest{PlanState: planState, StudentNumber: studentNumber})
	if err != nil || len(*result) < 1 {
		if errors.Is(err, context.DeadlineExceeded) {
			return c.Status(fiber.StatusGatewayTimeout).JSON(fiber.Map{
				"error": "context deadline exceeded",
			})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error": "plan not found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"plans": result,
	})
}
