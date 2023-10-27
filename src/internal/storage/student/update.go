package student

import (
	"fmt"
	"github.com/yigithankarabulut/vatansoftgocase/src/internal/storage/models"
)

func (s *studentStorage) Update(studentNumber int, updateData map[string]string) (models.Student, error) {
	var student models.Student

	result := s.db.Where("student_number = ?", studentNumber).First(&student)
	if result.Error != nil {
		return models.Student{}, fmt.Errorf("err: student not found")
	}
	result = s.db.Model(&student).Update("password", updateData["password"])
	if result.Error != nil {
		return models.Student{}, fmt.Errorf("err: student could not be updated")
	}
	return s.Get(studentNumber)
}
