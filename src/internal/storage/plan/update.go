package plan

import (
	"fmt"
	"github.com/yigithankarabulut/vatansoftgocase/src/internal/storage/models"
)

func (p *planStorage) Update(studentNumber int, planNumber int, updateData map[string]interface{}) (models.Plan, error) {
	var plan models.Plan
	result := p.db.Where("student_number = ? AND plan_number = ?", studentNumber, planNumber).First(&plan)
	if result.Error != nil {
		return models.Plan{}, fmt.Errorf("err: plan not found")
	}
	result = p.db.Model(&plan).Updates(updateData)
	if result.Error != nil {
		return models.Plan{}, fmt.Errorf("err: plan could not be updated")
	}
	if updateData["plan_number"] != nil {
		updatesNumber := updateData["plan_number"].(float64)
		planNumber = int(updatesNumber)
	}
	return p.GetPlan(studentNumber, planNumber)
}
