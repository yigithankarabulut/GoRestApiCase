package plan

import (
	"fmt"
	"github.com/yigithankarabulut/vatansoftgocase/src/internal/storage/models"
)

func (p *planStorage) CreatePlan(plan models.Plan) (models.Plan, error) {
	if _, err := p.GetPlan(plan.StudentNumber, plan.PlanNumber); err == nil {
		return models.Plan{}, fmt.Errorf("err: plan number already exists")
	}
	result := p.db.Create(&plan)
	if result.Error != nil {
		return models.Plan{}, fmt.Errorf("err: plan could not be created")
	}
	return plan, nil
}
