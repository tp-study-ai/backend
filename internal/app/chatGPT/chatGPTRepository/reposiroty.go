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

func (r *RepositoryChatGPT) GetTaskById(id int) (Task models.TaskDB, err error) {
	err = r.DB.QueryRow(
		`select *
		from "tasks"
		where id = $1;`,
		id,
	).Scan(
		&Task.Id,
		&Task.Name,
		&Task.Description,
		&Task.PublicTests,
		&Task.PrivateTests,
		&Task.GeneratedTests,
		&Task.Difficulty,
		&Task.CfContestId,
		&Task.CfIndex,
		&Task.CfPoints,
		&Task.CfRating,
		&Task.CfTags,
		&Task.TimeLimit,
		&Task.MemoryLimitBytes,
		&Task.Link,
		&Task.ShortLink,
		&Task.NameRu,
		&Task.TaskRu,
		&Task.Input,
		&Task.Output,
		&Task.Note,
		&Task.MasterSolution,
	)

	return
}
