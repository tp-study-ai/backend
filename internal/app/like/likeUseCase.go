package like

import "github.com/tp-study-ai/backend/internal/app/models"

type UseCase interface {
	LikeTask(like models.LikeJson) (err error)
	DeleteLike(like models.LikeJson) (err error)
	GetLikeTask(UserId models.UserId) (*models.LikeTasks, error)
}
