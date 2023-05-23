package task

import (
	"github.com/tp-study-ai/backend/internal/app/models"
)

type Repository interface {
	GetTask() (*models.TaskDB, error)
	GetTaskById(id int) (*models.TaskDB, error)
	GetTaskByLimit(id int, sort string, tag []int, minRating int, maxRating int) (*models.TasksResponse, int, error)
	GetSendTask(UserId int) (*models.SendTasks, error)
	GetSendTaskByTaskId(UserId int, TaskId int) (*models.SendTasks, error)
	GetAllUserTask(id int) (*[]int, error)
	GetDoneTask(id int) (*[]int, error)
	SetDifficultyTask(difficulty models.DifficultyDb) error
}
