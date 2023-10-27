package httpservice

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yigithankarabulut/vatansoftgocase/src/internal/transport/http/middlewares"
)

func (r *httpServiceHandler) Router(app fiber.Router) {
	app.Post("/register", r.RegisterStudent)
	app.Post("/login", r.LoginStudent)

	app.Use(middlewares.JWTMiddleware())

	app.Post("/logout", r.LogoutStudent)
	app.Put("/update", r.UpdatePassword)

	app.Post("/plan/create", r.CreatePlan)

	app.Get("/plan/get/", r.GetPlan)               //?planNumber=1
	app.Get("/plan/getByState/", r.GetPlanByState) //?state=completed

	app.Put("/plan/cancel/", r.CancelPlan)     //?planNumber=1
	app.Put("/plan/complete/", r.CompletePlan) //?planNumber=1
	app.Put("/plan/start/", r.StartPlan)       //?planNumber=1
	app.Put("/plan/update/", r.UpdatePlan)     //?planNumber=1

	app.Delete("/plan/delete/", r.DeletePlan) //?planNumber=1

	app.Get("/plan/listAll", r.ListAllPlans)
	app.Post("/plan/listWeekly/", r.ListWeeklyPlans) //?lastweek=1
	app.Post("/plan/listMonthly", r.ListMonthlyPlans)
}
