package task

import (
	"github.com/tp-study-ai/backend/internal/app/models"
	"time"
)

type Repository interface {
	GetTask() (Task models.TaskDB, err error)
	GetTaskById(id int) (Task models.TaskDB, err error)
	GetTaskByLimit(id int, sort string, tag []int) (*models.TasksResponse, int, error)
	SendTask(task *models.SendTask) (*models.SendTask, error)
	GetTaskByLink(link string) (Task models.TaskDB, err error)
	GetSendTask(UserId int) (*models.SendTasks, error)
	LikeTask(like models.LikeDb) (err error)
	GetLike(like models.LikeDb) (*models.LikeDb, error)
	GetLikes(UserId models.UserId) (*models.LikesDb, error)
	DeleteLike(like models.LikeDb) (err error)
	GetCountTaskOfDate(id int, day time.Time) (int, error)
}
