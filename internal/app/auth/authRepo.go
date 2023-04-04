package auth

import "github.com/tp-study-ai/backend/internal/app/models"

type Repository interface {
	CreateUser(UserRequest *models.UserDB) (UserId models.UserId, err error)
}
