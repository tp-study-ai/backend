package chatGPTRepository

import (
	"github.com/jackc/pgx"
	"github.com/tp-study-ai/backend/internal/app/models"
)

type RepositoryChatGPT struct {
	DB *pgx.ConnPool
}

func NewRepositoryChatGPT(db *pgx.ConnPool) *RepositoryChatGPT {
	return &RepositoryChatGPT{DB: db}
}

func (r *RepositoryChatGPT) GetTaskForChatGPT(id int) (*models.TaskDbForChatGPT, error) {
	Task := &models.TaskDbForChatGPT{}
	sql := `select * from "tasks" where id = $1;`

	err := r.DB.QueryRow(
		sql,
		id,
	).Scan(
		&Task.Description,
		&Task.MasterSolution,
	)
	if err != nil {
		return nil, err
	}

	return Task, nil
}
