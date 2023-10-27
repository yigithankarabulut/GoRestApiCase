package httpservice

import (
	"context"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/yigithankarabulut/vatansoftgocase/src/internal/service/planService"
	"gorm.io/gorm"
	"strconv"
)

func (s *httpServiceHandler) UpdatePlan(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), s.Handler.CancelTimeout)
	defer cancel()

	planNumber, err := strconv.Atoi(c.Query("planNumber"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "plan number is not valid",
		})
	}
	var updatePlanRequest planService.UpdatePlanRequest
	updatePlanRequest.UpdateData = make(map[string]interface{})
	if err := c.BodyParser(&updatePlanRequest.UpdateData); err != nil {
		if len(c.Body()) == 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "request body is empty",
			})
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "request body is not valid",
		})
	}
	studentNumber, err := s.JwtGetStudentNumber(c.Cookies("jwt"))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "jwt token is not valid",
		})
	}
	updatePlanRequest.PlanNumber = planNumber
	updatePlanRequest.StudentNumber = studentNumber

	result, err := s.planService.Get(ctx, &planService.GetPlanRequest{PlanNumber: updatePlanRequest.PlanNumber, StudentNumber: updatePlanRequest.StudentNumber})
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

	for k, v := range updatePlanRequest.UpdateData {
		if k == "status" {
			if v != "continue" && v != "completed" && v != "cancelled" && v != "created" {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error": "please enter a valid status",
				})
			}
		} else if k == "plan_description" || k == "plan_number" || k == "date" || k == "start_hour" || k == "end_hour" {
			continue
		} else {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "please only enter status, plan_number, plan_description, date, start_hour, end_hour",
			})
		}
	}

	if int(updatePlanRequest.UpdateData["plan_number"].(float64)) > 0 {
		_, err := s.planService.Get(ctx, &planService.GetPlanRequest{PlanNumber: int(updatePlanRequest.UpdateData["plan_number"].(float64)), StudentNumber: studentNumber})
		if err == nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "plan number already exists",
			})
		} else if errors.Is(err, context.DeadlineExceeded) {
			return c.Status(fiber.StatusGatewayTimeout).JSON(fiber.Map{
				"error": "context deadline exceeded",
			})
		}
		result.Plan.PlanNumber = int(updatePlanRequest.UpdateData["plan_number"].(float64))
	}
	if updatePlanRequest.UpdateData["plan_description"] != nil {
		result.Plan.PlanDescription = updatePlanRequest.UpdateData["plan_description"].(string)
	}
	if updatePlanRequest.UpdateData["date"] != nil {
		result.Plan.Date = updatePlanRequest.UpdateData["date"].(string)
	}
	if updatePlanRequest.UpdateData["start_hour"] != nil {
		result.Plan.StartHour = updatePlanRequest.UpdateData["start_hour"].(string)
	}
	if updatePlanRequest.UpdateData["end_hour"] != nil {
		result.Plan.EndHour = updatePlanRequest.UpdateData["end_hour"].(string)
	}
	if updatePlanRequest.UpdateData["status"] != nil {
		result.Plan.Status = updatePlanRequest.UpdateData["status"].(string)
	}

	dateFormat := "02.01.2006"
	timeFormat := "15:04"
	date, err := s.Handler.TimeFormatChecker(dateFormat, result.Plan.Date)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "date format is not valid. date format must be dd.mm.yyyy",
		})
	}
	startHour, err := s.Handler.TimeFormatChecker(timeFormat, result.Plan.StartHour)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "start hour format is not valid. start hour format must be hh:mm",
		})
	}
	endHour, err := s.Handler.TimeFormatChecker(timeFormat, result.Plan.EndHour)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "end hour format is not valid. end hour format must be hh:mm",
		})
	}

	if err = s.Handler.TimeValidChecker(date, startHour, endHour); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	uptResult, err := s.planService.Update(ctx, &updatePlanRequest)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return c.Status(fiber.StatusGatewayTimeout).JSON(fiber.Map{
				"error": "context deadline exceeded",
			})
		}
		if errors.Is(err, gorm.ErrInvalidValue) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "student number cannot be changed",
			})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error": "plan could not be updated",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "plan updated",
		"plan":    uptResult.Plan,
	})
}
