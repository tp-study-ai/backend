package testis

import "github.com/tp-study-ai/backend/internal/app/models"

type Repository interface {
	GetTaskForTestis(id int) (*models.TaskDBForTestis, error)
	SendTask(task *models.SendTask) (*models.SendTask, error)
}
