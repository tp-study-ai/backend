package auth

import "github.com/tp-study-ai/backend/internal/app/models"

type UseCase interface {
	Register(User *models.UserJson) (*models.ResponseUserJson, error)
	Login(*models.UserJson) (*models.ResponseUserJson, error)
	GetUserById(id models.UserId) (*models.ResponseUserJson, error)
	Update(UserRequest *models.UpdateJson, UserId models.UserId) (*models.ResponseUserJson, error)
}
