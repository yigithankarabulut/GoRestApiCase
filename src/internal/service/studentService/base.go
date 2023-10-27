package studentService

import (
	"context"
	"github.com/yigithankarabulut/vatansoftgocase/src/internal/storage/student"
)

type StudentStoreService interface {
	Set(context.Context, *SetStudentRequest) (*StudentResponse, error)
	Get(context.Context, int) (*StudentResponse, error)
	Update(context.Context, *UpdateStudentRequest) (*StudentResponse, error)
	Delete(context.Context, int) error
	List(context.Context) (*StudentListResponse, error)
}

type studentStoreService struct {
	storage student.StudentStorer
}

type StudentStoreServiceOption func(*studentStoreService)

func WithStorage(storage student.StudentStorer) StudentStoreServiceOption {
	return func(s *studentStoreService) {
		s.storage = storage
	}
}

func New(opts ...StudentStoreServiceOption) StudentStoreService {
	s := &studentStoreService{}
	for _, opt := range opts {
		opt(s)
	}
	return s
}
