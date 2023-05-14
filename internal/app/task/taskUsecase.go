package task

import "github.com/tp-study-ai/backend/internal/app/models"

type UseCase interface {
	GetTask() (task models.TaskJSON, err error)
	GetTaskById(id int) (Task models.TaskJSON, err error)
	//CheckSolution(solution models.CheckSolutionRequest, userId int) (*models.CheckSolutionUseCaseResponse, error)
	GetTaskByLimit(id int, sort string, tag []int, minRating int, maxRating int) (*models.TasksPagination, error)
	GetSimilar(solution models.SimilarRequest) (*models.Tasks, error)
	GetSendTask(UserId int) (*models.SendTasksJson, error)
	GetSendTaskByTaskId(UserId int, TaskId int) (*models.SendTasksJson, error)
	LikeTask(like models.LikeJson) (err error)
	GetLikeTask(UserId models.UserId) (*models.LikeTasks, error)
	DeleteLike(like models.LikeJson) (err error)
	GetCountTaskOfDate(id int) (*models.Days, error)
	GetShockMode(id int) (*models.ShockMode, error)
	GetDoneTask(id int) (*models.DoneTask, error)
	GetNotDoneTask(id int) (*models.DoneTask, error)
	SetDifficultyTask(difficulty models.DifficultyJson) error
	Recommendations(UserId int) (*models.RecResponse, error)
	Recommendations1(UserId int) (*models.RecResponse, error)
	ColdStart(UserId int) (*models.ColdStartResponse, error)
	Chat(Message models.ChatGPT) (*models.Message, error)
}
