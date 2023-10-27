package student

import (
	"github.com/yigithankarabulut/vatansoftgocase/src/internal/storage/models"
	"gorm.io/gorm"
)

type StudentStorer interface {
	Set(student models.Student) (models.Student, error)
	Get(studentNumber int) (models.Student, error)
	Update(studentNumber int, updateData map[string]string) (models.Student, error)
	Delete(studentNumber int) error
	List() ([]models.Student, error)
	Migrate() error
}

type studentStorage struct {
	db *gorm.DB
}

type StudentStorageOption func(*studentStorage)

func WithStudentDb(db *gorm.DB) StudentStorageOption {
	return func(s *studentStorage) {
		s.db = db
	}
}

func NewStudentStorage(opts ...StudentStorageOption) StudentStorer {
	s := &studentStorage{}
	for _, opt := range opts {
		opt(s)
	}
	return s
}
