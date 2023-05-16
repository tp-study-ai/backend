package task

import (
	"github.com/tp-study-ai/backend/internal/app/models"
)

type Repository interface {
	GetTask() (*models.TaskDB, error)
	GetTaskById(id int) (*models.TaskDB, error)
	GetTaskByLimit(id int, sort string, tag []int, minRating int, maxRating int) (*models.TasksResponse, int, error)
	GetTaskByLink(link string) (*models.TaskDB, error)
	GetSendTask(UserId int) (*models.SendTasks, error)
	GetSendTaskByTaskId(UserId int, TaskId int) (*models.SendTasks, error)
	GetAllUserTask(id int) (*[]int, error)
	GetDoneTask(id int) (*[]int, error)
	SetDifficultyTask(difficulty models.DifficultyDb) error
	GetSetDifficultyTasks(UserId int) (*[]int, error)
	GetSetDifficultyTask(UserId int, TaskId int) (*models.DifficultyDb, error)
	GetEasyTasksForUser(UserId int) (*[]int, error)
	GetHardTasksForUser(UserId int) (*[]int, error)
	//UpdateUserColdStart(UserId int) error
}
