package auth

import "github.com/tp-study-ai/backend/internal/app/models"

type UseCaseAuth struct {
	Repo Repository
}

func NewUseCaseAuth(TaskRepo Repository) *UseCaseAuth {
	return &UseCaseAuth{
		Repo: TaskRepo,
	}
}

func (u *UseCaseAuth) Register(User *models.UserJson) (UserId models.UserId, err error) {
	UserId, err = u.Repo.CreateUser(&models.UserDB{Username: User.Username, Password: User.Password})
	if err != nil {
		return models.UserId(0), err
	}
	return UserId, nil
}
