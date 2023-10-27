package plan

import (
	"github.com/yigithankarabulut/vatansoftgocase/src/internal/storage/models"
	"gorm.io/gorm"
)

type PlanStorer interface {
	CreatePlan(plan models.Plan) (models.Plan, error) // ** check date time other plans
	GetPlan(studentNumber int, planNumber int) (models.Plan, error)
	GetPlanByState(studentNumber int, state string) ([]models.Plan, error)
	Update(studentNumber int, planNumber int, updateData map[string]interface{}) (models.Plan, error) // Set state this
	Delete(studentNumber int, planNumber int) error
	ListAll(studentNumber int) ([]models.Plan, error)
	ListWeeklyPlans(studentNumber int, lastWeek int) ([]models.Plan, error)
	ListMonthlyPlans(studentNumber int, month int, year int) ([]models.Plan, error)
	Migrate() error
}

type planStorage struct {
	db *gorm.DB
}

type PlanStorageOption func(*planStorage)

func WithPlanDb(db *gorm.DB) PlanStorageOption {
	return func(s *planStorage) {
		s.db = db
	}
}

func NewPlanStorage(opts ...PlanStorageOption) PlanStorer {
	s := &planStorage{}
	for _, opt := range opts {
		opt(s)
	}
	return s
}
