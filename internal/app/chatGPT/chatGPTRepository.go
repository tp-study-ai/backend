package chatGPT

import "github.com/tp-study-ai/backend/internal/app/models"

type Repository interface {
	GetTaskForChatGPT(id int) (*models.TaskDbForChatGPT, error)
}
