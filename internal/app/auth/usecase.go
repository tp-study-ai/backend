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

func (u *UseCaseAuth) Register(User *models.UserJson) (*models.ResponseUserJson, error) {
	User1, err := u.Repo.GetUser(&models.UserDB{Username: User.Username, Password: User.Password})
	if err == nil && User1.Username == User.Username {
		return nil, errors.Errorf("такой пользователь уже существует")
	}

	User2, err := u.Repo.CreateUser(&models.UserDB{Username: User.Username, Password: User.Password})
	if err != nil {
		return nil, err
	}

	if User.Username != User2.Username || User.Password != User2.Password {
		return nil, errors.Errorf("некорректаня работа чего то там")
	}

	return &models.ResponseUserJson{Id: User2.Id, Username: User2.Username}, nil
}

func (u *UseCaseAuth) Login(User *models.UserJson) (*models.ResponseUserJson, error) {
	User1, err := u.Repo.Login(&models.UserDB{Username: User.Username, Password: User.Password})
	if err != nil {
		return nil, err
	}
	if User1.Username == User.Username && User1.Password == User.Password {
		return &models.ResponseUserJson{Id: User1.Id, Username: User1.Username}, nil
	}
	return nil, errors.Errorf("username || password не верно")
}

func (u *UseCaseAuth) GetUserById(id models.UserId) (*models.ResponseUserJson, error) {
	user, err := u.Repo.GetUserById(id)
	if err != nil {
		return nil, err
	}
	return &models.ResponseUserJson{Id: user.Id, Username: user.Username}, nil
}

func (u *UseCaseAuth) Update(UserRequest *models.UpdateJson) (*models.ResponseUserJson, error) {
	fmt.Println(UserRequest.NewUsername)
	if len(UserRequest.NewUsername) != 0 {
		fmt.Println("nice")
		UserResponse, err := u.Repo.UpdateUsername(
			&models.UpdateUsernameDb{
				Id:          UserRequest.Id,
				Username:    UserRequest.Username,
				NewUsername: UserRequest.NewUsername,
			},
		)
		fmt.Println(err)
		if err != nil {
			return nil, err
		}
		if len(UserRequest.NewPassword) == 0 {
			return &models.ResponseUserJson{Id: UserResponse.Id, Username: UserResponse.Username}, nil
		}
	}

	fmt.Println(UserRequest.NewPassword)
	if len(UserRequest.NewPassword) != 0 {
		fmt.Println("nice")
		UserResponse, err := u.Repo.UpdatePassword(
			&models.UpdatePasswordDb{
				Id:          UserRequest.Id,
				Username:    UserRequest.Username,
				NewPassword: UserRequest.NewPassword,
			},
		)
		if err != nil {
			return nil, err
		}
		return &models.ResponseUserJson{Id: UserResponse.Id, Username: UserResponse.Username}, nil
	}
	return nil, errors.Errorf("error")
}
