package httpservice

import (
	"context"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/yigithankarabulut/vatansoftgocase/src/internal/service/planService"
	"strconv"
)

func (s *httpServiceHandler) GetPlan(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), s.Handler.CancelTimeout)
	defer cancel()

	planNumber, err := strconv.Atoi(c.Query("planNumber"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "plan number is not valid",
		})
	}
	studentNumber, err := s.JwtGetStudentNumber(c.Cookies("jwt"))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "jwt token is not valid",
		})
	}
	result, err := s.planService.Get(ctx, &planService.GetPlanRequest{PlanNumber: planNumber, StudentNumber: studentNumber})
	if err != nil {
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
		"plan": result.Plan,
	})
}
