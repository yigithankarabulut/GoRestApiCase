package plan

import (
	"fmt"
	"github.com/yigithankarabulut/vatansoftgocase/src/internal/storage/models"
	"time"
)

func (p *planStorage) ListWeeklyPlans(studentNumber int, lastWeek int) ([]models.Plan, error) {
	var plans []models.Plan

	now := time.Now()
	now.Format("02.06.2006")

	lastWeekStartTime := now.AddDate(0, 0, -lastWeek*7)
	lastWeekEndTime := now.AddDate(0, 0, -(lastWeek-1)*7)

	query := `
    SELECT * FROM plans
	WHERE student_number = ?
	AND STR_TO_DATE(date, '%d.%m.%Y') >= ?
	AND STR_TO_DATE(date, '%d.%m.%Y') < ?
	ORDER BY STR_TO_DATE(date, '%d.%m.%Y') DESC
    `

	result := p.db.Raw(query, studentNumber, lastWeekStartTime, lastWeekEndTime).Scan(&plans)
	if result.Error != nil {
		return []models.Plan{}, fmt.Errorf("err: plans not found")
	}
	return plans, nil
}
