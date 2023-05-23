package testisRepository

import (
	"github.com/jackc/pgx"
	"github.com/tp-study-ai/backend/internal/app/models"
)

type RepositoryTask struct {
	DB *pgx.ConnPool
}

func NewRepositoryTask(db *pgx.ConnPool) *RepositoryTask {
	return &RepositoryTask{DB: db}
}

func (r *RepositoryTask) GetTaskForTestis(id int) (*models.TaskDBForTestis, error) {
	Task := &models.TaskDBForTestis{}
	sql := `select "private_tests", "time_limit" from "tasks" where id = $1;`

	err := r.DB.QueryRow(sql, id).Scan(
		&Task.PrivateTests,
		&Task.TimeLimit,
	)
	if err != nil {
		return nil, err
	}

	return Task, nil
}

func (r *RepositoryTask) SendTask(task *models.SendTask) (*models.SendTask, error) {
	Task1 := &models.SendTask{}
	err := r.DB.QueryRow(
		`INSERT INTO "send_task" (
			"user_id", "task_id", "check_time", "build_time", "check_result", "check_message", "tests_passed", "tests_total", "lint_success", "code_text"
			 ) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) 
	 RETURNING "id", "user_id", "task_id", "check_time", "build_time", "check_result", "check_message", "tests_passed", "tests_total", "lint_success", "code_text", "date"`,
		task.UserId, task.TaskId, task.CheckTime, task.BuildTime, task.CheckResult, task.CheckMessage, task.TestsPassed, task.TestsTotal, task.LintSuccess, task.CodeText,
	).Scan(
		&Task1.ID, &Task1.UserId, &Task1.TaskId, &Task1.CheckTime, &Task1.BuildTime, &Task1.CheckResult, &Task1.CheckMessage, &Task1.TestsPassed, &Task1.TestsTotal, &Task1.LintSuccess, &Task1.CodeText, &Task1.Date,
	)
	if err != nil {
		return nil, err
	}
	return Task1, nil
}
