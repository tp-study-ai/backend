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
