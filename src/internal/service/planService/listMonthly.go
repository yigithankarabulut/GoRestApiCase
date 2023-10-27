package planService

import (
	"context"
	"fmt"
)

func (p *planStoreService) ListMonthly(ctx context.Context, req *ListMonthlyPlansRequest) (*PlanListResponse, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		plans, err := p.storage.ListMonthlyPlans(req.StudentNumber, req.Month, req.Year)
		if err != nil {
			return nil, fmt.Errorf("planService.ListMonthly: %w", err)
		}
		response := make(PlanListResponse, len(plans))
		for i, v := range plans {
			response[i] = PlanResponse{Plan: v}
		}
		return &response, nil
	}
}
