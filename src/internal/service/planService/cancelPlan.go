package planService

import (
	"context"
	"fmt"
)

func (p *planStoreService) CancelPlan(ctx context.Context, req *SetStateRequest) (*PlanResponse, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		updateData := make(map[string]interface{})
		updateData["status"] = "cancelled"
		plan, err := p.storage.Update(req.StudentNumber, req.PlanNumber, updateData)
		if err != nil {
			return nil, fmt.Errorf("planService.CancelPlan: %w", err)
		}
		return &PlanResponse{Plan: plan}, nil
	}
}
