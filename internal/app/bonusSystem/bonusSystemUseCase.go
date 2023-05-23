package bonusSystem

import "github.com/tp-study-ai/backend/internal/app/models"

type UseCase interface {
	GetCountTaskOfDate(id int) (*models.Days, error)
	GetShockMode(id int) (*models.ShockMode, error)
}
