package plan

import (
	"fmt"
	"github.com/yigithankarabulut/vatansoftgocase/src/internal/storage/models"
)

func (p *planStorage) GetPlanByState(studentNumber int, state string) ([]models.Plan, error) {
	var plans []models.Plan
	result := p.db.Where("student_number = ? AND status = ?", studentNumber, state).Find(&plans)
	if result.Error != nil {
		return []models.Plan{}, fmt.Errorf("err: plans not found")
	}
	return plans, nil
}
