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
	var count int
	err = r.DB.QueryRow(
		`select count(*) from "tasks"`,
	).Scan(&count)

	err = r.DB.QueryRow(
		`select *
		from "tasks"
		where id = $1;`,
		rand.Intn(count-1)+1,
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

func (r *RepositoryTask) GetTaskById(id int) (Task models.TaskResponse, err error) {
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
		&Task.TaskRu,
		&Task.Input,
		&Task.Output,
		&Task.Note,
	)
	return
}

func (r *RepositoryTask) GetTaskByLimit(id int) (*models.TasksResponse, error) {
	tasks := &models.TasksResponse{}

	var newPostsData []interface{}
	newPostsData = append(newPostsData, 10)
	newPostsData = append(newPostsData, 10*id)

	rows, err := r.DB.Query(`select * from "tasks" order by "id" limit $1 offset $2`, newPostsData...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var buff models.TaskResponse
		err = rows.Scan(
			&buff.Id,
			&buff.Name,
			&buff.Description,
			&buff.PublicTests,
			&buff.PrivateTests,
			&buff.GeneratedTests,
			&buff.Difficulty,
			&buff.CfContestId,
			&buff.CfIndex,
			&buff.CfPoints,
			&buff.CfRating,
			&buff.CfTags,
			&buff.TimeLimit,
			&buff.MemoryLimitBytes,
			&buff.Link,
			&buff.TaskRu,
			&buff.Input,
			&buff.Output,
			&buff.Note,
		)

		if err != nil {
			return nil, err
		}

		tasks.Tasks = append(tasks.Tasks, buff)
	}

	return tasks, err
}
