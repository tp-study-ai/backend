package task

import (
	"github.com/tp-study-ai/backend/internal/app/models"
	"time"
)

type Repository interface {
	GetTask() (Task models.TaskDB, err error)
	GetTaskById(id int) (Task models.TaskDB, err error)
	GetTaskByLimit(id int, sort string, tag []int, minRating int, maxRating int) (*models.TasksResponse, int, error)
	SendTask(task *models.SendTask) (*models.SendTask, error)
	GetTaskByLink(link string) (Task models.TaskDB, err error)
	GetSendTask(UserId int) (*models.SendTasks, error)
	GetSendTaskByTaskId(UserId int, TaskId int) (*models.SendTasks, error)
	LikeTask(like models.LikeDb) (err error)
	GetLike(like models.LikeDb) (*models.LikeDb, error)
	GetLikes(UserId models.UserId) (*models.LikesDb, error)
	DeleteLike(like models.LikeDb) (err error)
	GetCountTaskOfDate(id int, day time.Time) (int, error)
	GetAllUserTask(id int) (*[]int, error)
	GetDoneTask(id int) (*[]int, error)
	SetDifficultyTask(difficulty models.DifficultyDb) error
	GetSetDifficultyTasks(UserId int) (*[]int, error)
	GetSetDifficultyTask(UserId int, TaskId int) (*models.DifficultyDb, error)
	GetEasyTasksForUser(UserId int) (*[]int, error)
	GetHardTasksForUser(UserId int) (*[]int, error)
	UpdateUserColdStart(UserId int) error
}
