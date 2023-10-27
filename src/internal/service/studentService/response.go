package studentService

import "github.com/yigithankarabulut/vatansoftgocase/src/internal/storage/models"

type StudentResponse struct {
	Student models.Student
}

type StudentListResponse []StudentResponse
