package student

import (
	"fmt"
	"github.com/yigithankarabulut/vatansoftgocase/src/internal/storage/models"
)

func (s *studentStorage) Get(studentNumber int) (models.Student, error) {
	var student models.Student
	result := s.db.Where("student_number = ?", studentNumber).First(&student)
	if result.Error != nil {
		return models.Student{}, fmt.Errorf("err: student not found")
	}
	return student, nil
}
