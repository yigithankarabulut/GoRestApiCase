package httpservice

import (
	"context"
	"errors"
	"github.com/gofiber/fiber/v2"
	ss "github.com/yigithankarabulut/vatansoftgocase/src/internal/service/studentService"
	"github.com/yigithankarabulut/vatansoftgocase/src/internal/storage/models"
	"time"
)

func (s *httpServiceHandler) RegisterStudent(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), s.Handler.CancelTimeout)
	defer cancel()

	var student models.Student
	if err := c.BodyParser(&student); err != nil {
		if len(c.Body()) == 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "request body is empty",
			})
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "request body is not valid",
		})
	}
	if student.StudentNumber == 0 || student.Password == "" || student.Name == "" || student.Lastname == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "empty field/fields",
		})
	}
	if getResult, err := s.studentService.Get(ctx, student.StudentNumber); err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return c.Status(fiber.StatusGatewayTimeout).JSON(fiber.Map{
				"error": "request timeout",
			})
		}
	} else if getResult.Student.StudentNumber == student.StudentNumber {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "student already exists",
		})
	}
	hashPassword, err := s.Handler.GeneratePasswordHash(student.Password)
	if err != nil {
		return c.Status(fiber.StatusGatewayTimeout).JSON(fiber.Map{
			"error": "hashing password failed",
		})
	}
	student.Password = hashPassword
	result, err := s.studentService.Set(ctx, &ss.SetStudentRequest{Student: student})
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return c.Status(fiber.StatusGatewayTimeout).JSON(fiber.Map{
				"error": "context deadline exceeded",
			})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	jwtErr, jwtToken := s.JwtCreate(result.Student.StudentNumber)
	if jwtErr != nil {
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
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "student registered successfully",
	})
}
