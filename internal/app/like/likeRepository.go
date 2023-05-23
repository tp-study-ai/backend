package like

import "github.com/tp-study-ai/backend/internal/app/models"

type Repository interface {
	GetTaskById(id int) (*models.TaskDB, error)
	LikeTask(like models.LikeDb) (err error)
	DeleteLike(like models.LikeDb) (err error)
	GetLike(like models.LikeDb) (*models.LikeDb, error)
	GetLikes(UserId models.UserId) (*models.LikesDb, error)
}
