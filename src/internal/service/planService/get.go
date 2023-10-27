package planService

import (
	"context"
	"fmt"
)

func (p *planStoreService) Get(ctx context.Context, req *GetPlanRequest) (*PlanResponse, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		plan, err := p.storage.GetPlan(req.StudentNumber, req.PlanNumber)
		if err != nil {
			return nil, fmt.Errorf("planService.Get: %w", err)
		}
		return &PlanResponse{Plan: plan}, nil
	}
}
