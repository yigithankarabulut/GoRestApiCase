package planService

import (
	"context"
	"fmt"
)

func (p *planStoreService) GetByState(ctx context.Context, req *GetPlanByStateRequest) (*PlanListResponse, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		plans, err := p.storage.GetPlanByState(req.StudentNumber, req.PlanState)
		if err != nil {
			return nil, fmt.Errorf("planService.GetByState: %w", err)
		}
		response := make(PlanListResponse, len(plans))
		for i, v := range plans {
			response[i] = PlanResponse{Plan: v}
		}
		return &response, nil
	}
}
