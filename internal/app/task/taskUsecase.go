package task

import "github.com/tp-study-ai/backend/internal/app/models"

type UseCase interface {
	GetTask() (task models.TaskJSON, err error)
	GetTaskById(id int) (Task models.TaskJSON, err error)
	GetTaskByLimit(id int, sort string, tag []int, minRating int, maxRating int) (*models.TasksPagination, error)
	GetSendTask(UserId int) (*models.SendTasksJson, error)
	GetSendTaskByTaskId(UserId int, TaskId int) (*models.SendTasksJson, error)
	GetDoneTask(id int) (*models.DoneTask, error)
	GetNotDoneTask(id int) (*models.DoneTask, error)
	SetDifficultyTask(difficulty models.DifficultyJson) error
}
