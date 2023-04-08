package auth

import "github.com/tp-study-ai/backend/internal/app/models"

type Repository interface {
	GetUser(UserRequest *models.UserDB) (username string, err error)
	CreateUser(UserRequest *models.UserDB) (UserId models.UserId, err error)
	Login(UserRequest *models.UserDB) (models.User, error)
	GetUserByd(id int) (user models.ResponseUserDb, err error)
}
