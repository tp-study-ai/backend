package task

import "github.com/tp-study-ai/backend/internal/app/models"

type UseCase interface {
	GetTask() (task models.TaskJSON, err error)
	GetTaskById(id int) (Task models.TaskJSON, err error)
	CheckSolution(solution models.CheckSolutionRequest) (cheche models.CheckSolutionUseCaseResponse, err error)
	GetTaskByLimit(id int, sort string, tag []int) (*models.Tasks, error)
}
