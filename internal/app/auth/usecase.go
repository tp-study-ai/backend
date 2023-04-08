package auth

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/tp-study-ai/backend/internal/app/models"
)

type UseCaseAuth struct {
	Repo Repository
}

func NewUseCaseAuth(TaskRepo Repository) *UseCaseAuth {
	return &UseCaseAuth{
		Repo: TaskRepo,
	}
}

func (u *UseCaseAuth) Register(User *models.UserJson) (UserId models.UserId, err error) {
	username, err := u.Repo.GetUser(&models.UserDB{Username: User.Username, Password: User.Password})
	//fmt.Println(err.Error(), username, User.Username, User.Password)
	if err == nil && username == User.Username {
		fmt.Println(username)
		return models.UserId(0), errors.Errorf("такой пользователь уже существует")
	}

	UserId, err = u.Repo.CreateUser(&models.UserDB{Username: User.Username, Password: User.Password})
	if err != nil {
		return models.UserId(0), err
	}
	return UserId, nil
}

func (u *UseCaseAuth) Login(User *models.UserJson) (bool, error) {
	User1, err := u.Repo.Login(&models.UserDB{Username: User.Username, Password: User.Password})
	if err != nil {
		return false, err
	}
	if User1.Username == User.Username && User1.Password == User.Password {
		return true, nil
	}
	return false, errors.Errorf("username || password не верно")
}

func (u *UseCaseAuth) GetUserById(id int) (models.ResponseUserJson, error) {
	user, err := u.Repo.GetUserByd(id)
	if err != nil {
		return models.ResponseUserJson{}, err
	}

	user1 := models.ResponseUserJson{Id: user.Id, Username: user.Username}
	return user1, nil
}
