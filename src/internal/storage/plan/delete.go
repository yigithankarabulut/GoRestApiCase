package plan

import (
	"fmt"
	"github.com/yigithankarabulut/vatansoftgocase/src/internal/storage/models"
)

func (p *planStorage) Delete(studentNumber int, planNumber int) error {
	if _, err := p.GetPlan(studentNumber, planNumber); err != nil {
		return fmt.Errorf("err: plan not found")
	}
	result := p.db.Where("student_number = ? AND plan_number = ?", studentNumber, planNumber).Delete(&models.Plan{})
	if result.Error != nil {
		return fmt.Errorf("err: plan could not be deleted")
	}
	return nil
}
