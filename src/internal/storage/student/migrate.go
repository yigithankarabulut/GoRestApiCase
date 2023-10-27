package student

import (
	"github.com/yigithankarabulut/vatansoftgocase/src/internal/storage/models"
)

func (s *studentStorage) Migrate() error {
	err := s.db.AutoMigrate(&models.Student{})
	if err != nil {
		return err
	}
	return nil
}
