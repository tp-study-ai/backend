package task

import "github.com/tp-study-ai/backend/internal/app/models"

type Repository interface {
	GetTask() (Task models.TaskDB, err error)
	GetTaskById(id int) (Task models.TaskDB, err error)
	GetTaskByLimit(id int, sort string, tag []int) (*models.TasksResponse, error)
	SendTask(task *models.SendTask) (*models.SendTask, error)
	GetTaskByLink(link string) (Task models.TaskDB, err error)
	GetSendTask(UserId int) (*models.SendTasks, error)
}
