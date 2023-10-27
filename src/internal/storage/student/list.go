package student

import (
	"fmt"
	"github.com/yigithankarabulut/vatansoftgocase/src/internal/storage/models"
)

func (s *studentStorage) List() ([]models.Student, error) {
	var students []models.Student
	result := s.db.Find(&students)
	if result.Error != nil {
		return nil, fmt.Errorf("err: students not found")
	}
	return students, nil
}
