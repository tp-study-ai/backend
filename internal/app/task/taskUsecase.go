package task

import "github.com/tp-study-ai/backend/internal/app/models"

type Ucase interface {
	GetTask() (Task models.TaskResponse, err error)
	GetTaskById(id int) (Task models.TaskResponse, err error)
}