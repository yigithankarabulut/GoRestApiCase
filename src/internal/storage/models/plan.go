package models

import (
	"gorm.io/gorm"
)

type Plan struct {
	gorm.Model
	PlanNumber      int    `json:"plan_number" gorm:"primaryKey"`
	StudentNumber   int    `json:"student_number,omitempty"`
	PlanDescription string `json:"plan_description"`
	Date            string `json:"date"`
	StartHour       string `json:"start_hour"`
	EndHour         string `json:"end_hour"`
	Status          string `json:"status,omitempty"`
}
