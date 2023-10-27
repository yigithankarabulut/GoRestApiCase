package httpservice

import (
	"context"
	"errors"
	"github.com/gofiber/fiber/v2"
	ss "github.com/yigithankarabulut/vatansoftgocase/src/internal/service/studentService"
	"strings"
	"time"
)

func (s *httpServiceHandler) UpdatePassword(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), s.Handler.CancelTimeout)
	defer cancel()

	var password map[string]string
	if err := c.BodyParser(&password); err != nil {
		if len(c.Body()) == 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "request body is empty",
			})
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "request body is not valid",
		})
	}
	i := 0
	for k, _ := range password {
		if strings.Compare(k, "newPassword") == 0 || strings.Compare(k, "password") == 0 {
			i++
		}
	}
	if i != 2 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "please only enter password and new password",
		})
	}
	if password["password"] == "" || password["newPassword"] == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "password or new password is empty",
		})
	}
	if len(password["password"]) < 6 || len(password["newPassword"]) < 6 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "please enter at least 6 characters",
		})
	}
	if password["password"] == password["newPassword"] {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "password and new password cannot be the same",
		})
	}
	studentNumber, err := s.JwtGetStudentNumber(c.Cookies("jwt"))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "jwt token is not valid",
		})
	}
	getResult, err := s.studentService.Get(ctx, studentNumber)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return c.Status(fiber.StatusGatewayTimeout).JSON(fiber.Map{
				"error": "context deadline exceeded",
			})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error": "student not found",
		})
	}
	if s.Handler.ComparePasswordHash(password["password"], getResult.Student.Password) == false {
		c.Cookie(&fiber.Cookie{
			Name:     "jwt",
			Value:    "",
			Expires:  time.Now().Add(-time.Hour),
			HTTPOnly: true,
		})
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "wrong password",
		})
	}
	hashPassword, err := s.Handler.GeneratePasswordHash(password["newPassword"])
	if err != nil {
		return c.Status(fiber.StatusGatewayTimeout).JSON(fiber.Map{
			"error": "hashing password failed",
		})
	}
	req := ss.UpdateStudentRequest{
		StudentNumber: studentNumber,
		UpdateData:    map[string]string{"password": hashPassword},
	}
	if _, err := s.studentService.Update(ctx, &req); err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return c.Status(fiber.StatusGatewayTimeout).JSON(fiber.Map{
				"error": "context deadline exceeded",
			})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "password updated successfully",
	})
}
