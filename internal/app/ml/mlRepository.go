package ml

import "github.com/tp-study-ai/backend/internal/app/models"

type Repository interface {
	GetTaskByLink(link string) (Task models.TaskDB, err error)
	GetEasyTasksForUser(UserId int) (*[]int, error)
	GetHardTasksForUser(UserId int) (*[]int, error)
	GetDoneTask(id int) (*[]int, error)
	GetTaskById(id int) (Task models.TaskDB, err error)
	GetSendTaskByTaskId(UserId int, TaskId int) (*models.SendTasks, error)
	UpdateUserColdStart(UserId int) error
}
