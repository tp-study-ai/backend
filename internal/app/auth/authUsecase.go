package auth

import "github.com/tp-study-ai/backend/internal/app/models"

type UseCase interface {
	Register(User *models.UserJson) (UserId models.UserId, err error)
	Login(User *models.UserJson) (b bool, err error)
	GetUserById(id int) (models.ResponseUserJson, error)
}
