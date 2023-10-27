package planService

import "github.com/yigithankarabulut/vatansoftgocase/src/internal/storage/models"

type SetPlanRequest struct {
	Plan models.Plan
}

type SetStateRequest struct {
	PlanNumber    int `json:"plan_number"`
	StudentNumber int `json:"student_number,omitempty"`
}

type GetPlanRequest struct {
	PlanNumber    int `json:"plan_number"`
	StudentNumber int `json:"student_number,omitempty"`
}

type GetPlanByStateRequest struct {
	PlanState     string `json:"plan_state"`
	StudentNumber int    `json:"student_number,omitempty"`
}

type UpdatePlanRequest struct {
	PlanNumber    int                    `json:"plan_number,omitempty"`
	StudentNumber int                    `json:"student_number,omitempty"`
	UpdateData    map[string]interface{} `json:"update_data,omitempty"`
}

type DeletePlanRequest struct {
	PlanNumber    int `json:"plan_number"`
	StudentNumber int `json:"student_number,omitempty"`
}

type ListAllPlansRequest struct {
	StudentNumber int `json:"student_number,omitempty"`
}

type ListWeeklyPlansRequest struct {
	StudentNumber int `json:"student_number,omitempty"`
	LastWeek      int `json:"last_week"`
}

type ListMonthlyPlansRequest struct {
	StudentNumber int `json:"student_number,omitempty"`
	Month         int `json:"month"`
	Year          int `json:"year"`
}
