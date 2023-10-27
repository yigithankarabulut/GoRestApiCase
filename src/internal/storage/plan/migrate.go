package plan

import (
	"github.com/yigithankarabulut/vatansoftgocase/src/internal/storage/models"
)

func (p *planStorage) Migrate() error {
	err := p.db.AutoMigrate(&models.Plan{})
	if err != nil {
		return err
	}
	return nil
}
