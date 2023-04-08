package auth

import (
	"github.com/jackc/pgx"
	"github.com/tp-study-ai/backend/internal/app/models"
)

type RepositoryAuth struct {
	DB *pgx.ConnPool
}

func NewRepositoryAuth(db *pgx.ConnPool) *RepositoryAuth {
	return &RepositoryAuth{DB: db}
}

func (r *RepositoryAuth) GetUserByd(id int) (user models.ResponseUserDb, err error) {
	err = r.DB.QueryRow(`select "id", "username" from "users" where "id" = $1`, id).Scan(&user.Id, &user.Username)
	if err != nil {
		return models.ResponseUserDb{}, err
	}
	return
}

func (r *RepositoryAuth) GetUser(UserRequest *models.UserDB) (string, error) {
	var username string
	err := r.DB.QueryRow(`select "username" from "users" where "username" = $1 limit 1`, UserRequest.Username).Scan(&username)
	if err != nil {
		return "", err
	}
	return username, nil
}

func (r *RepositoryAuth) CreateUser(UserRequest *models.UserDB) (UserId models.UserId, err error) {
	err = r.DB.QueryRow(
		`INSERT INTO "users" ("username","password") VALUES ($1,$2) RETURNING id`,
		UserRequest.Username, UserRequest.Password,
	).Scan(&UserId)
	if err != nil {
		return models.UserId(0), err
	}
	return UserId, err
}

func (r *RepositoryAuth) Login(UserRequest *models.UserDB) (models.User, error) {
	var U models.User
	err := r.DB.QueryRow(
		`SELECT * FROM "users" WHERE "username" = $1 and "password" = $2`,
		UserRequest.Username, UserRequest.Password).Scan(
		&U.Id, &U.Username, &U.Password)
	if err != nil {
		return models.User{}, err
	}
	return U, nil
}
