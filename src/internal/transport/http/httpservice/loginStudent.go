package httpservice

import (
	"context"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/yigithankarabulut/vatansoftgocase/src/internal/storage/models"
	"time"
)

func (s *httpServiceHandler) LoginStudent(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), s.Handler.CancelTimeout)
	defer cancel()

	var student models.Student
	if err := c.BodyParser(&student); err != nil {
		c.Cookie(&fiber.Cookie{
			Name:     "jwt",
			Value:    "",
			Expires:  time.Now().Add(-time.Hour),
			HTTPOnly: true,
		})
		if len(c.Body()) == 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "request body is empty",
			})
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "request body is not valid",
		})
	}
	if student.StudentNumber == 0 || student.Password == "" {
		c.Cookie(&fiber.Cookie{
			Name:     "jwt",
			Value:    "",
			Expires:  time.Now().Add(-time.Hour),
			HTTPOnly: true,
		})
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "student number or password is empty",
		})
	}
	if len(student.Name) > 0 || len(student.Lastname) > 0 {
		c.Cookie(&fiber.Cookie{
			Name:     "jwt",
			Value:    "",
			Expires:  time.Now().Add(-time.Hour),
			HTTPOnly: true,
		})
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "please only enter student number and password",
		})
	}
	result, err := s.studentService.Get(ctx, student.StudentNumber)
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
	if s.Handler.ComparePasswordHash(student.Password, result.Student.Password) == false {
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
	jwtErr, jwtToken := s.JwtCreate(result.Student.StudentNumber)
	if jwtErr != nil {
		c.Cookie(&fiber.Cookie{
			Name:     "jwt",
			Value:    "",
			Expires:  time.Now().Add(-time.Hour),
			HTTPOnly: true,
		})
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error": "jwt token cannot created",
		})
	}
	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    jwtToken,
		Expires:  time.Now().Add(time.Hour * 6),
		HTTPOnly: true,
	})
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":        "student logged in successfully",
		"student number": result.Student.StudentNumber,
	})
}
