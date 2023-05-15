package ml

import "github.com/tp-study-ai/backend/internal/app/models"

type UseCase interface {
	GetSimilar(solution models.SimilarRequest) (*models.Tasks, error)
	Recommendations(UserId int) (*models.RecResponse, error)
	ColdStart(UserId int) (*models.ColdStartResponse, error)
}
