package studentService

import (
	"context"
	"fmt"
)

func (s *studentStoreService) List(ctx context.Context) (*StudentListResponse, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		items, err := s.storage.List()
		if err != nil {
			return nil, fmt.Errorf("studentService.List: %w", err)
		}
		response := make(StudentListResponse, len(items))
		for i, v := range items {
			response[i] = StudentResponse{Student: v}
		}
		return &response, nil
	}
}
