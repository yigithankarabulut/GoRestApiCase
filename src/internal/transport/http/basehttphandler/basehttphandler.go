package basehttphandler

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"log/slog"
	"os"
	"time"
)

type Handler struct {
	ServerEnv     string
	Logger        *slog.Logger
	CancelTimeout time.Duration
}

func (h *Handler) JwtCreate(studentNumber int) (error, string) {
	var env []byte = []byte(os.Getenv("JWT_SECRET"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"StudentNumber": studentNumber,
		"ExpiresAt":     time.Now().Add(time.Hour * 6).Unix(),
	})
	tokenString, err := token.SignedString(env)
	if err != nil {
		return fmt.Errorf("err: %w", err), ""
	}
	return nil, tokenString
}

func (h *Handler) JwtGetStudentNumber(jwtKey string) (int, error) {
	var env []byte = []byte(os.Getenv("JWT_SECRET"))
	token, err := jwt.Parse(jwtKey, func(token *jwt.Token) (interface{}, error) {
		return env, nil
	})
	if err != nil || !token.Valid {
		return 0, fmt.Errorf("err: %w", err)
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, fmt.Errorf("err: %w", err)
	}
	result := claims["StudentNumber"].(float64)
	return int(result), nil
}

func (h *Handler) GeneratePasswordHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (h *Handler) ComparePasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (h *Handler) TimeFormatChecker(format, srcTime string) (time.Time, error) {
	result, err := time.Parse(format, srcTime)
	if err != nil {
		return time.Time{}, err
	}
	return result, nil
}

func (h *Handler) TimeValidChecker(date, startHour, endHour time.Time) error {
	if date.Before(time.Now()) {
		if date.Day() == time.Now().Day() {
			if startHour.Hour() < time.Now().Hour() {
				return fmt.Errorf("start hour must be after current hour")
			} else if startHour.Hour() == time.Now().Hour() && startHour.Minute() <= time.Now().Minute() {
				return fmt.Errorf("start hour must be after current hour")
			}
		} else {
			return fmt.Errorf("date must be after today")
		}
	}
	if startHour.Before(endHour) == false {
		return fmt.Errorf("start hour must be before end hour")
	}
	return nil
}
