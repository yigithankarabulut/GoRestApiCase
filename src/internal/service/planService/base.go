package planService

import (
	"context"
	"github.com/yigithankarabulut/vatansoftgocase/src/internal/storage/plan"
)

type PlanStoreService interface {
	Set(context.Context, *SetPlanRequest) (*PlanResponse, error)
	StartPlan(context.Context, *SetStateRequest) (*PlanResponse, error)
	CancelPlan(context.Context, *SetStateRequest) (*PlanResponse, error)
	CompletePlan(context.Context, *SetStateRequest) (*PlanResponse, error)
	Get(context.Context, *GetPlanRequest) (*PlanResponse, error)
	GetByState(context.Context, *GetPlanByStateRequest) (*PlanListResponse, error)
	Update(context.Context, *UpdatePlanRequest) (*PlanResponse, error)
	Delete(context.Context, *DeletePlanRequest) error
	ListAll(context.Context, *ListAllPlansRequest) (*PlanListResponse, error)
	ListWeekly(context.Context, *ListWeeklyPlansRequest) (*PlanListResponse, error)
	ListMonthly(context.Context, *ListMonthlyPlansRequest) (*PlanListResponse, error)
}

type planStoreService struct {
	storage plan.PlanStorer
}

type PlanStoreServiceOption func(*planStoreService)

func WithStorage(storage plan.PlanStorer) PlanStoreServiceOption {
	return func(s *planStoreService) {
		s.storage = storage
	}
}

func New(opts ...PlanStoreServiceOption) PlanStoreService {
	s := &planStoreService{}
	for _, opt := range opts {
		opt(s)
	}
	return s
}
