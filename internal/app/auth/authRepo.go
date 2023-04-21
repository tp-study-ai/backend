package auth

import "github.com/tp-study-ai/backend/internal/app/models"

type Repository interface {
	GetUserById(id models.UserId) (*models.UserDB, error)
	GetUser(UserRequest *models.UserDB) (*models.UserDB, error)
	CreateUser(UserRequest *models.UserDB) (*models.UserDB, error)
	Login(UserRequest *models.UserDB) (*models.UserDB, error)
	UpdateUsername(UserRequest *models.UpdateUsernameDb) (*models.UserDB, error)
	UpdatePassword(UserRequest *models.UpdatePasswordDb) (*models.UserDB, error)
}
