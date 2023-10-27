package planService

import (
	"context"
	"fmt"
)

func (p *planStoreService) Set(ctx context.Context, req *SetPlanRequest) (*PlanResponse, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		plan, err := p.storage.CreatePlan(req.Plan)
		if err != nil {
			return nil, fmt.Errorf("planService.Set: %w", err)
		}
		return &PlanResponse{Plan: plan}, nil
	}
}
