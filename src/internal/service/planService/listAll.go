package planService

import (
	"context"
	"fmt"
)

func (p *planStoreService) ListAll(ctx context.Context, req *ListAllPlansRequest) (*PlanListResponse, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		plans, err := p.storage.ListAll(req.StudentNumber)
		if err != nil {
			return nil, fmt.Errorf("planService.ListAll: %w", err)
		}
		response := make(PlanListResponse, len(plans))
		for i, v := range plans {
			response[i] = PlanResponse{Plan: v}
		}
		return &response, nil
	}
}
