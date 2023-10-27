package planService

import (
	"context"
	"fmt"
)

func (p *planStoreService) Update(ctx context.Context, req *UpdatePlanRequest) (*PlanResponse, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		plan, err := p.storage.Update(req.StudentNumber, req.PlanNumber, req.UpdateData)
		if err != nil {
			return nil, fmt.Errorf("planService.Update: %w", err)
		}
		return &PlanResponse{Plan: plan}, nil
	}
}
