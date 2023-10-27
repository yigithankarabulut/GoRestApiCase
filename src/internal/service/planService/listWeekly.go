package planService

import (
	"context"
	"fmt"
)

func (p *planStoreService) ListWeekly(ctx context.Context, req *ListWeeklyPlansRequest) (*PlanListResponse, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		plans, err := p.storage.ListWeeklyPlans(req.StudentNumber, req.LastWeek)
		if err != nil {
			return nil, fmt.Errorf("planService.ListWeekly: %w", err)
		}
		response := make(PlanListResponse, len(plans))
		for i, v := range plans {
			response[i] = PlanResponse{Plan: v}
		}
		return &response, nil
	}
}
