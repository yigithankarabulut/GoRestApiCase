package studentService

import (
	"context"
	"fmt"
)

func (s *studentStoreService) Set(ctx context.Context, req *SetStudentRequest) (*StudentResponse, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		student, err := s.storage.Set(req.Student)
		if err != nil {
			return nil, fmt.Errorf("studentService.Set: %w", err)
		}
		return &StudentResponse{Student: student}, nil
	}
}
