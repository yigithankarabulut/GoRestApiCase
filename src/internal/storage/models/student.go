package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	StudentNumber int    `json:"student_number" gorm:"unique"`
	Name          string `json:"name,omitempty"`
	Lastname      string `json:"lastname,omitempty"`
	Password      string `json:"password"`
	Plans         []Plan `gorm:"foreignKey:StudentNumber;references:StudentNumber"`
}
