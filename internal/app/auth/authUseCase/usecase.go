package authUseCase

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/tp-study-ai/backend/internal/app/auth"
	"github.com/tp-study-ai/backend/internal/app/models"
	"github.com/tp-study-ai/backend/tools"
)

type UseCaseAuth struct {
	Repo auth.Repository
}

func NewUseCaseAuth(TaskRepo auth.Repository) *UseCaseAuth {
	return &UseCaseAuth{
		Repo: TaskRepo,
	}
}

func (u *UseCaseAuth) Register(User *models.UserJson) (*models.ResponseUserJson, error) {
	User1, err := u.Repo.GetUser(&models.UserDB{Username: User.Username, Password: tools.GetMD5Hash(User.Password)})
	if err == nil && User1.Username == User.Username {
		return nil, errors.Errorf("такой пользователь уже существует")
	}

	User2, err := u.Repo.CreateUser(&models.UserDB{Username: User.Username, Password: tools.GetMD5Hash(User.Password)})
	if err != nil {
		return nil, err
	}

	if User.Username != User2.Username || tools.GetMD5Hash(User.Password) != User2.Password {
		return nil, errors.Errorf("некорректаня работа чего то там")
	}

	return &models.ResponseUserJson{Id: User2.Id, Username: User2.Username, ColdStart: User2.ColdStart}, nil
}

func (u *UseCaseAuth) Login(User *models.UserJson) (*models.ResponseUserJson, error) {
	User1, err := u.Repo.Login(&models.UserDB{Username: User.Username, Password: tools.GetMD5Hash(User.Password)})
	if err != nil {
		return nil, err
	}
	if User1.Username == User.Username && User1.Password == tools.GetMD5Hash(User.Password) {
		return &models.ResponseUserJson{Id: User1.Id, Username: User1.Username}, nil
	}
	return nil, errors.Errorf("username || password не верно")
}

func (u *UseCaseAuth) GetUserById(id models.UserId) (*models.ResponseUserJson, error) {
	user, err := u.Repo.GetUserById(id)
	if err != nil {
		return nil, err
	}
	return &models.ResponseUserJson{Id: user.Id, Username: user.Username, ColdStart: user.ColdStart}, nil
}

func (u *UseCaseAuth) Update(UserRequest *models.UpdateJson, UserId models.UserId) (*models.ResponseUserJson, error) {
	Uuser, err := u.Repo.GetUserById(UserId)
	if err != nil {
		return nil, err
	}

	fmt.Println(Uuser)

	fmt.Println(UserRequest.NewUsername)
	if len(UserRequest.NewUsername) != 0 {
		fmt.Println("nice")
		UserResponse, err1 := u.Repo.UpdateUsername(
			&models.UpdateUsernameDb{
				Id:          Uuser.Id,
				Username:    Uuser.Username,
				NewUsername: UserRequest.NewUsername,
			},
		)
		fmt.Println(err1)
		if err1 != nil {
			return nil, err1
		}
		if len(UserRequest.NewPassword) == 0 {
			return &models.ResponseUserJson{Id: UserResponse.Id, Username: UserResponse.Username, ColdStart: UserResponse.ColdStart}, nil
		}
	}

	fmt.Println(UserRequest.NewPassword)
	if len(UserRequest.NewPassword) != 0 {
		fmt.Println("nice")
		UserResponse, err1 := u.Repo.UpdatePassword(
			&models.UpdatePasswordDb{
				Id:          Uuser.Id,
				Username:    Uuser.Username,
				NewPassword: tools.GetMD5Hash(UserRequest.NewPassword),
			},
		)
		if err1 != nil {
			return nil, err1
		}
		return &models.ResponseUserJson{Id: UserResponse.Id, Username: UserResponse.Username, ColdStart: UserResponse.ColdStart}, nil
	}
	return nil, errors.Errorf("error")
}
