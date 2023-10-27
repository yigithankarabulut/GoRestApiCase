package studentService

import (
	"context"
	"fmt"
)

func (s *studentStoreService) Get(ctx context.Context, studentNumber int) (*StudentResponse, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		student, err := s.storage.Get(studentNumber)
		if err != nil {
			return nil, fmt.Errorf("studentService.Get: %w", err)
		}
		return &StudentResponse{Student: student}, nil
	}
}
