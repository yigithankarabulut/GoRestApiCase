package planService

import (
	"context"
	"fmt"
)

func (p *planStoreService) StartPlan(ctx context.Context, req *SetStateRequest) (*PlanResponse, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		updateData := make(map[string]interface{})
		updateData["status"] = "continue"
		plan, err := p.storage.Update(req.StudentNumber, req.PlanNumber, updateData)
		if err != nil {
			return nil, fmt.Errorf("planService.StartPlan: %w", err)
		}
		return &PlanResponse{Plan: plan}, nil
	}
}
