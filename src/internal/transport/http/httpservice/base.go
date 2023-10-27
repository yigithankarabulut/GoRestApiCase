package httpservice

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yigithankarabulut/vatansoftgocase/src/internal/service/planService"
	"github.com/yigithankarabulut/vatansoftgocase/src/internal/service/studentService"
	"github.com/yigithankarabulut/vatansoftgocase/src/internal/transport/http/basehttphandler"
	"log/slog"
	"time"
)

type HttpService interface {
	// Student operations
	RegisterStudent(c *fiber.Ctx) error
	LoginStudent(c *fiber.Ctx) error
	LogoutStudent(c *fiber.Ctx) error
	UpdatePassword(c *fiber.Ctx) error

	// Plan operations
	GetPlan(c *fiber.Ctx) error
	GetPlanByState(c *fiber.Ctx) error
	CreatePlan(c *fiber.Ctx) error
	CancelPlan(c *fiber.Ctx) error
	CompletePlan(c *fiber.Ctx) error
	StartPlan(c *fiber.Ctx) error
	UpdatePlan(c *fiber.Ctx) error
	DeletePlan(c *fiber.Ctx) error
	ListAllPlans(c *fiber.Ctx) error
	ListWeeklyPlans(c *fiber.Ctx) error
	ListMonthlyPlans(c *fiber.Ctx) error

	// Route operations
	Router(app fiber.Router)
}

type httpServiceHandler struct {
	basehttphandler.Handler
	studentService studentService.StudentStoreService
	planService    planService.PlanStoreService
}

type HttpServiceOption func(*httpServiceHandler)

func WithStudentService(studentService studentService.StudentStoreService) HttpServiceOption {
	return func(s *httpServiceHandler) {
		s.studentService = studentService
	}
}

func WithPlanService(planService planService.PlanStoreService) HttpServiceOption {
	return func(s *httpServiceHandler) {
		s.planService = planService
	}
}

func WithContextTimeout(time time.Duration) HttpServiceOption {
	return func(s *httpServiceHandler) {
		s.Handler.CancelTimeout = time
	}
}

func WithServerEnv(env string) HttpServiceOption {
	return func(s *httpServiceHandler) {
		s.Handler.ServerEnv = env
	}
}

func WithLogger(logger *slog.Logger) HttpServiceOption {
	return func(s *httpServiceHandler) {
		s.Handler.Logger = logger
	}
}

func New(opts ...HttpServiceOption) HttpService {
	s := &httpServiceHandler{
		Handler: basehttphandler.Handler{},
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}
