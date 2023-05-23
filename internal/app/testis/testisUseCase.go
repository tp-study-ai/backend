package testis

import "github.com/tp-study-ai/backend/internal/app/models"

type UseCase interface {
	CheckSolution(solution models.CheckSolutionRequest, userId int) (*models.CheckSolutionUseCaseResponse, error)
}
