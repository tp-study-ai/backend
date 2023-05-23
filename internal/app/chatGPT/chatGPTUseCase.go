package chatGPT

import "github.com/tp-study-ai/backend/internal/app/models"

type UseCase interface {
	Chat(Message models.ChatGPT) (*models.Message, error)
}
