package plan

import (
	"fmt"
	"github.com/yigithankarabulut/vatansoftgocase/src/internal/storage/models"
	"time"
)

func (p *planStorage) ListMonthlyPlans(studentNumber int, month, year int) ([]models.Plan, error) {
	var plans []models.Plan

	startDay := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	endDay := startDay.AddDate(0, 1, -1) // Ayın son günü

	query := `
	SELECT * FROM plans
	WHERE student_number = ?
	AND STR_TO_DATE(date, '%d.%m.%Y') >= ?
	AND STR_TO_DATE(date, '%d.%m.%Y') <= ?
	ORDER BY STR_TO_DATE(date, '%d.%m.%Y') DESC
	`

	result := p.db.Raw(query, studentNumber, startDay, endDay).Scan(&plans)
	if result.Error != nil {
		return []models.Plan{}, fmt.Errorf("err: plans not found")
	}
	return plans, nil
}
