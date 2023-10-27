package student

import (
	"fmt"
	"github.com/yigithankarabulut/vatansoftgocase/src/internal/storage/models"
)

func (s *studentStorage) Set(Student models.Student) (models.Student, error) {
	if _, err := s.Get(Student.StudentNumber); err == nil {
		return models.Student{}, fmt.Errorf("err: student already exists")
	}
	result := s.db.Create(&Student)
	if result.Error != nil {
		return models.Student{}, fmt.Errorf("err: student could not be created")
	}
	return Student, nil
}
