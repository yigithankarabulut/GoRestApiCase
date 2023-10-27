package httpservice

import (
	"context"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/yigithankarabulut/vatansoftgocase/src/internal/service/planService"
	"strconv"
)

func (s *httpServiceHandler) StartPlan(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), s.Handler.CancelTimeout)
	defer cancel()

	planNumber, err := strconv.Atoi(c.Query("planNumber"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "plan number is not valid",
		})
	}
	var stateRequest planService.SetStateRequest
	studentNumber, err := s.JwtGetStudentNumber(c.Cookies("jwt"))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "jwt token is not valid",
		})
	}
	stateRequest.PlanNumber = planNumber
	stateRequest.StudentNumber = studentNumber
	if _, err := s.planService.Get(ctx, &planService.GetPlanRequest{PlanNumber: planNumber, StudentNumber: studentNumber}); err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return c.Status(fiber.StatusGatewayTimeout).JSON(fiber.Map{
				"error": "context deadline exceeded",
			})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error": "plan not found",
		})
	}
	result, err := s.planService.StartPlan(ctx, &stateRequest)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return c.Status(fiber.StatusGatewayTimeout).JSON(fiber.Map{
				"error": "context deadline exceeded",
			})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error": "plan state could not be changed",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "plan started",
		"plan":    result.Plan,
	})
}
