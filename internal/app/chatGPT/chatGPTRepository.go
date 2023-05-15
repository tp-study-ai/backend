package chatGPT

import "github.com/tp-study-ai/backend/internal/app/models"

type Repository interface {
	GetTaskById(id int) (Task models.TaskDB, err error)
}
