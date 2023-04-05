package task

import "github.com/tp-study-ai/backend/internal/app/models"

type UseCase interface {
	GetTask() (Task models.TaskResponse, err error)
	GetTaskById(id int) (Task models.TaskResponse, err error)
	CheckSolution(solution models.CheckSolutionRequest) (cheche models.CheckSolutionUseCaseResponse, err error)
	GetTaskByLimit(id int, sort string, tag []int) (*models.Tasks, error)
}
