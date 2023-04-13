package task

import "github.com/tp-study-ai/backend/internal/app/models"

type UseCase interface {
	GetTask() (task models.TaskJSON, err error)
	GetTaskById(id int) (Task models.TaskJSON, err error)
	CheckSolution(solution models.CheckSolutionRequest, userId int) (*models.CheckSolutionUseCaseResponse, error)
	GetTaskByLimit(id int, sort string, tag []int) (*models.Tasks, error)
	GetSimilar(solution models.SimilarRequest) (*models.Tasks, error)
	GetSendTask(UserId int) (*models.SendTasksJson, error)
	LikeTask(like models.LikeJson) (err error)
	GetLikeTask(UserId models.UserId) (*models.LikeTasks, error)
	DeleteLike(like models.LikeJson) (err error)
}
