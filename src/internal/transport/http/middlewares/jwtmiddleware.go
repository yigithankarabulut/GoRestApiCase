package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

func JWTMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		cookie := c.Cookies("jwt")
		var env []byte = []byte(os.Getenv("JWT_SECRET"))
		if cookie == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}
		token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
			return env, nil
		})
		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}
		expirationTime := time.Unix(int64(claims["ExpiresAt"].(float64)), 0)
		currentTime := time.Now()
		if currentTime.After(expirationTime) {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}
		return c.Next()
	}
}
