package testis

import "github.com/tp-study-ai/backend/internal/app/models"

type Repository interface {
	GetTaskById(id int) (Task models.TaskDB, err error)
	SendTask(task *models.SendTask) (*models.SendTask, error)
}
