package planService

import (
	"context"
	"fmt"
)

func (p *planStoreService) Delete(ctx context.Context, req *DeletePlanRequest) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		if err := p.storage.Delete(req.StudentNumber, req.PlanNumber); err != nil {
			return fmt.Errorf("planService.Delete: %w", err)
		}
		return nil
	}
}
