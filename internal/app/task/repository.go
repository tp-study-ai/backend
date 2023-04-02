package task

import (
	"github.com/jackc/pgx"
	"github.com/tp-study-ai/backend/internal/app/models"
	"math/rand"
)

type RepositoryTask struct {
	DB *pgx.ConnPool
}

func NewRepositoryTask(db *pgx.ConnPool) *RepositoryTask {
	return &RepositoryTask{DB: db}
}

func (r *RepositoryTask) GetTask() (Task models.TaskResponse, err error) {
	err = r.DB.QueryRow(
		`select *
		from "task"
		where id = $1;`,
		rand.Intn(3584-1)+1,
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
		&Task.TaskRu,
		&Task.Input,
		&Task.Output,
		&Task.Note,
	)

	//fmt.Println(Task)
	//fmt.Println(
	//	Task.Id,
	//	Task.Name,
	//	Task.Description,
	//	Task.PublicTests,
	//	Task.PrivateTests,
	//	Task.GeneratedTests,
	//	Task.Difficulty,
	//	Task.CfContestId,
	//	Task.CfIndex,
	//	Task.CfPoints,
	//	Task.CfRating,
	//	Task.CfTags,
	//	Task.TimeLimit,
	//	Task.MemoryLimitBytes,
	//	Task.Link,
	//	Task.TaskRu,
	//	Task.Input,
	//	Task.Output,
	//	Task.Note,
	//)
	return
}

func (r *RepositoryTask) GetTaskById(id int) (Task models.TaskResponse, err error) {
	err = r.DB.QueryRow(
		`select *
		from "task"
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
		&Task.TaskRu,
		&Task.Input,
		&Task.Output,
		&Task.Note,
	)
	return
}
