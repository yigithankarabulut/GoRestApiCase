package studentService

import (
	"context"
	"fmt"
)

func (s *studentStoreService) Delete(ctx context.Context, studentNumber int) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		if err := s.storage.Delete(studentNumber); err != nil {
			return fmt.Errorf("studentService.Delete: %w", err)
		}
		return nil
	}
}
