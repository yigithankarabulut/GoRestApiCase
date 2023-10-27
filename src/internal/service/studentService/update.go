package studentService

import (
	"context"
	"fmt"
)

func (s *studentStoreService) Update(ctx context.Context, req *UpdateStudentRequest) (*StudentResponse, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		student, err := s.storage.Update(req.StudentNumber, req.UpdateData)
		if err != nil {
			return nil, fmt.Errorf("studentService.Update: %w", err)
		}
		return &StudentResponse{Student: student}, nil
	}
}
