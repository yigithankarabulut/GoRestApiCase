package plan

import (
	"fmt"
	"github.com/yigithankarabulut/vatansoftgocase/src/internal/storage/models"
)

func (p *planStorage) GetPlan(studentNumber int, planNumber int) (models.Plan, error) {
	var plan models.Plan
	result := p.db.Where("student_number = ? AND plan_number = ?", studentNumber, planNumber).First(&plan)
	if result.Error != nil {
		return models.Plan{}, fmt.Errorf("err: plan not found")
	}
	return plan, nil
}
