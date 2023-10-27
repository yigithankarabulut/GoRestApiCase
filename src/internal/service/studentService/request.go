package studentService

import "github.com/yigithankarabulut/vatansoftgocase/src/internal/storage/models"

type SetStudentRequest struct {
	Student models.Student
}

type UpdateStudentRequest struct {
	StudentNumber int
	UpdateData    map[string]string
}
