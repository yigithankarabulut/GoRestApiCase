package apiserver

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/yigithankarabulut/vatansoftgocase/database"
	"github.com/yigithankarabulut/vatansoftgocase/src/internal/service/planService"
	"github.com/yigithankarabulut/vatansoftgocase/src/internal/service/studentService"
	"github.com/yigithankarabulut/vatansoftgocase/src/internal/storage/plan"
	"github.com/yigithankarabulut/vatansoftgocase/src/internal/storage/student"
	"github.com/yigithankarabulut/vatansoftgocase/src/internal/transport/http/httpservice"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const ContextCancelTimeout = 5 * time.Second

type apiServer struct {
	logLevel  slog.Level
	logger    *slog.Logger
	serverEnv string
}

type Option func(*apiServer)

func WithLogger(logger *slog.Logger) Option {
	return func(s *apiServer) {
		s.logger = logger
	}
}

func WithLogLevel(level string) Option {
	return func(s *apiServer) {
		var logLevel slog.Level

		switch level {
		case "DEBUG":
			logLevel = slog.LevelDebug
		case "INFO":
			logLevel = slog.LevelInfo
		case "WARN":
			logLevel = slog.LevelWarn
		case "ERROR":
			logLevel = slog.LevelError
		default:
			logLevel = slog.LevelInfo
		}
		s.logLevel = logLevel
	}
}

func WithServerEnv(env string) Option {
	return func(s *apiServer) {
		s.serverEnv = env
	}
}

func New(opts ...Option) error {
	apiServer := &apiServer{
		logLevel: slog.LevelInfo,
	}
	db, err := database.ConnectDB()
	if err != nil {
		return err
	}
	app := fiber.New()
	app.Use(recover.New())

	for _, opt := range opts {
		opt(apiServer)
	}

	if apiServer.logger == nil {
		logHandlerOpts := &slog.HandlerOptions{Level: apiServer.logLevel}
		logHandler := slog.NewJSONHandler(os.Stdout, logHandlerOpts)
		apiServer.logger = slog.New(logHandler)
	}

	slog.SetDefault(apiServer.logger)
	if apiServer.serverEnv == "" {
		apiServer.serverEnv = "development"
	}

	logger := apiServer.logger

	studentStorage := student.NewStudentStorage(student.WithStudentDb(db))
	planStorage := plan.NewPlanStorage(plan.WithPlanDb(db))

	if err := studentStorage.Migrate(); err != nil {
		return fmt.Errorf("studentStorage.Migrate: %w", err)
	}
	if err := planStorage.Migrate(); err != nil {
		return fmt.Errorf("planStorage.Migrate: %w", err)
	}

	studentServices := studentService.New(studentService.WithStorage(studentStorage))
	planServices := planService.New(planService.WithStorage(planStorage))

	httpStoreHandler := httpservice.New(
		httpservice.WithStudentService(studentServices),
		httpservice.WithPlanService(planServices),
		httpservice.WithContextTimeout(ContextCancelTimeout),
		httpservice.WithServerEnv(apiServer.serverEnv),
		httpservice.WithLogger(logger),
	)

	httpStoreHandler.Router(app)

	shutdown := make(chan os.Signal, 1)
	apiError := make(chan error, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		logger.Info("Starting api server...")
		apiError <- app.Listen(":8080")
	}()

	select {
	case err := <-apiError:
		return fmt.Errorf("listen error: %w", err)
	case <-shutdown:
		logger.Info("Starting shutdown", "pid", os.Getpid())
		time.Sleep(1 * time.Second)
		defer logger.Info("Shutdown complete", "pid", os.Getpid())
	}
	return nil
}
