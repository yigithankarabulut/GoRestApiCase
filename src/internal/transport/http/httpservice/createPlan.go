package httpservice

import (
	"context"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/yigithankarabulut/vatansoftgocase/src/internal/service/planService"
	"github.com/yigithankarabulut/vatansoftgocase/src/internal/storage/models"
)

func (s *httpServiceHandler) CreatePlan(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), s.Handler.CancelTimeout)
	defer cancel()

	var plan models.Plan
	if err := c.BodyParser(&plan); err != nil {
		if len(c.Body()) == 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "request body is empty",
			})
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "request body is not valid",
		})
	}
	if plan.PlanNumber == 0 || plan.PlanDescription == "" || len(plan.Date) < 1 || len(plan.EndHour) < 1 || len(plan.StartHour) < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "empty field/fields",
		})
	}
	dateFormat := "02.01.2006"
	timeFormat := "15:04"

	date, err := s.Handler.TimeFormatChecker(dateFormat, plan.Date)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "date format is not valid. date format must be dd.mm.yyyy",
		})
	}
	startHour, err := s.Handler.TimeFormatChecker(timeFormat, plan.StartHour)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "start hour format is not valid. start hour format must be hh:mm",
		})
	}
	endHour, err := s.Handler.TimeFormatChecker(timeFormat, plan.EndHour)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "end hour format is not valid. end hour format must be hh:mm",
		})
	}

	if err := s.Handler.TimeValidChecker(date, startHour, endHour); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	studentNumber, err := s.JwtGetStudentNumber(c.Cookies("jwt"))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "jwt token is not valid",
		})
	}
	plan.StudentNumber = studentNumber
	result, err := s.planService.Get(ctx, &planService.GetPlanRequest{PlanNumber: plan.PlanNumber, StudentNumber: plan.StudentNumber})
	if err == nil {
		if result.Plan.PlanNumber == plan.PlanNumber {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "plan number must be unique",
			})
		}
	} else if errors.Is(err, context.DeadlineExceeded) {
		return c.Status(fiber.StatusGatewayTimeout).JSON(fiber.Map{
			"error": "context deadline exceeded",
		})
	}
	if len(plan.Status) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "please only change plan state with set state endpoint",
		})
	}
	plan.Status = "created"

	response, err := s.planService.Set(ctx, &planService.SetPlanRequest{Plan: plan})
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return c.Status(fiber.StatusGatewayTimeout).JSON(fiber.Map{
				"error": "context deadline exceeded",
			})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error": "plan cannot created",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "plan created successfully",
		"plan":    response.Plan,
	})
} // TODO: check plan date and time other plans
