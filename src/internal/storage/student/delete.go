package student

import (
	"fmt"
	"github.com/yigithankarabulut/vatansoftgocase/src/internal/storage/models"
)

func (s *studentStorage) Delete(StudentNumber int) error {
	if _, err := s.Get(StudentNumber); err != nil {
		return fmt.Errorf("err: student not found")
	}
	result := s.db.Delete(&models.Student{}, StudentNumber)
	if result.Error != nil {
		return fmt.Errorf("err: student could not be deleted")
	}
	return nil
}
