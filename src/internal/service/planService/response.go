package planService

import "github.com/yigithankarabulut/vatansoftgocase/src/internal/storage/models"

type PlanResponse struct {
	Plan models.Plan
}

type PlanListResponse []PlanResponse
