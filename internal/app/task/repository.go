package task

import (
	"github.com/jackc/pgx"
	"github.com/lib/pq"
	"github.com/tp-study-ai/backend/internal/app/models"
	"math/rand"
)

type RepositoryTask struct {
	DB *pgx.ConnPool
}

func NewRepositoryTask(db *pgx.ConnPool) *RepositoryTask {
	return &RepositoryTask{DB: db}
}

func (r *RepositoryTask) GetTask() (Task models.TaskDB, err error) {
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

func (r *RepositoryTask) GetTaskById(id int) (Task models.TaskDB, err error) {
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

func (r *RepositoryTask) GetTaskByLimit(id int, sort string, tag []int) (*models.TasksResponse, error) {
	tasks := &models.TasksResponse{}

	var newPostsData []interface{}
	newPostsData = append(newPostsData, 15)
	newPostsData = append(newPostsData, 15*id)

	sql := `select * from "tasks" `

	if len(tag) != 0 {
		sql = sql + `where $3 <@ (cf_tags)`
		newPostsData = append(newPostsData, pq.Array(tag))
	}

	if sort == "" {
		sql = sql + `order by "id"`
	}

	if sort == "rating_up" {
		sql = sql + `order by "cf_rating"`
	}

	if sort == "rating_down" {
		sql = sql + `order by "cf_rating" desc`
	}

	sql = sql + ` limit $1 offset $2`

	rows, err := r.DB.Query(sql, newPostsData...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var buff models.TaskDB
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

func (r *RepositoryTask) SendTask(task *models.SendTask) (*models.SendTask, error) {
	Task1 := &models.SendTask{}
	err := r.DB.QueryRow(
		`INSERT INTO "send_task" (
			"user_id", "task_id", "check_time", "build_time", "check_result", "check_message", "tests_passed", "tests_total", "lint_success", "code_text"
			 ) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) 
	 RETURNING "id", "user_id", "task_id", "check_time", "build_time", "check_result", "check_message", "tests_passed", "tests_total", "lint_success", "code_text", get_ru_date(date)`,
		task.UserId, task.TaskId, task.CheckTime, task.BuildTime, task.CheckResult, task.CheckMessage, task.TestsPassed, task.TestsTotal, task.LintSuccess, task.CodeText,
	).Scan(
		&Task1.ID, &Task1.UserId, &Task1.TaskId, &Task1.CheckTime, &Task1.BuildTime, &Task1.CheckResult, &Task1.CheckMessage, &Task1.TestsPassed, &Task1.TestsTotal, &Task1.LintSuccess, &Task1.CodeText, &Task1.Date,
	)
	if err != nil {
		return nil, err
	}
	return Task1, nil
}
