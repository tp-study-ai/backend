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

func (r *RepositoryAuth) GetUserById(id models.UserId) (*models.UserDB, error) {
	UserResponse := &models.UserDB{}
	err := r.DB.QueryRow(
		`select "id", "username", "password", "cold_start" from "users" where "id" = $1`,
		id,
	).Scan(&UserResponse.Id, &UserResponse.Username, &UserResponse.Password, &UserResponse.ColdStart)
	if err != nil {
		return nil, err
	}
	return UserResponse, nil
}

func (r *RepositoryAuth) GetUser(UserRequest *models.UserDB) (*models.UserDB, error) {
	UserResponse := &models.UserDB{}
	err := r.DB.QueryRow(
		`select "id", "username", "password", "cold_start" from "users" where "username" = $1 limit 1`,
		UserRequest.Username,
	).Scan(&UserResponse.Id, &UserResponse.Username, &UserResponse.Password, &UserResponse.ColdStart)
	if err != nil {
		return nil, err
	}
	return UserResponse, nil
}

func (r *RepositoryAuth) CreateUser(UserRequest *models.UserDB) (*models.UserDB, error) {
	UserResponse := &models.UserDB{}
	err := r.DB.QueryRow(
		`INSERT INTO "users" ("username","password","cold_start") VALUES ($1,$2,'false') RETURNING "id", "username", "password", "cold_start"`,
		UserRequest.Username, UserRequest.Password,
	).Scan(&UserResponse.Id, &UserResponse.Username, &UserResponse.Password, &UserResponse.ColdStart)
	if err != nil {
		return nil, err
	}
	return UserResponse, nil
}

func (r *RepositoryAuth) Login(UserRequest *models.UserDB) (*models.UserDB, error) {
	UserResponse := &models.UserDB{}
	err := r.DB.QueryRow(
		`SELECT * FROM "users" WHERE "username" = $1 and "password" = $2`,
		UserRequest.Username, UserRequest.Password,
	).Scan(&UserResponse.Id, &UserResponse.Username, &UserResponse.Password, &UserResponse.ColdStart)
	if err != nil {
		return nil, err
	}
	return UserResponse, nil
}

func (r *RepositoryAuth) UpdateUsername(UserRequest *models.UpdateUsernameDb) (*models.UserDB, error) {
	UserResponse := &models.UserDB{}
	sql := `UPDATE "users" SET "username" = $1 WHERE "username" = $2 and "id" = $3 RETURNING "id", "username", "password", "cold_start";`
	err := r.DB.QueryRow(
		sql,
		UserRequest.NewUsername, UserRequest.Username, UserRequest.Id,
	).Scan(&UserResponse.Id, &UserResponse.Username, &UserResponse.Password, &UserResponse.ColdStart)
	if err != nil {
		return nil, err
	}
	return UserResponse, nil
}

func (r *RepositoryAuth) UpdatePassword(UserRequest *models.UpdatePasswordDb) (*models.UserDB, error) {
	UserResponse := &models.UserDB{}
	sql := `UPDATE "users" SET "password" = $1 WHERE "username" = $2 and "id" = $3 RETURNING "id", "username", "password", "cold_start";`
	err := r.DB.QueryRow(
		sql,
		UserRequest.NewPassword, UserRequest.Username, UserRequest.Id,
	).Scan(&UserResponse.Id, &UserResponse.Username, &UserResponse.Password, &UserResponse.ColdStart)
	if err != nil {
		return nil, err
	}
	return UserResponse, nil
}
