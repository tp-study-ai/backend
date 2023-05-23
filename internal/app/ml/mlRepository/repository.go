package mlRepository

import (
	"fmt"
	"github.com/jackc/pgx"
	"github.com/pkg/errors"
	"github.com/tp-study-ai/backend/internal/app/models"
)

type RepositoryML struct {
	DB *pgx.ConnPool
}

func NewRepositoryML(db *pgx.ConnPool) *RepositoryML {
	return &RepositoryML{DB: db}
}

func (r *RepositoryML) GetTaskByLink(link string) (Task models.TaskDB, err error) {
	err = r.DB.QueryRow(
		`select *
		from "tasks"
		where link = $1;`,
		link,
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
		&Task.NameRu,
		&Task.TaskRu,
		&Task.Input,
		&Task.Output,
		&Task.Note,
		&Task.MasterSolution,
	)

	return
}

func (r *RepositoryML) GetEasyTasksForUser(UserId int) (*[]int, error) {
	var easyTask []int
	var newPostsData []interface{}
	newPostsData = append(newPostsData, UserId)

	sql := `SELECT "task_id" FROM "difficulty_task" WHERE "user_id" = $1 and "difficulty" = -1;`

	rows, err := r.DB.Query(sql, newPostsData...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var buff int
		err = rows.Scan(
			&buff,
		)
		if err != nil {
			return nil, err
		}
		//fmt.Println(buff)

		easyTask = append(easyTask, buff)
	}

	return &easyTask, nil
}

func (r *RepositoryML) GetHardTasksForUser(UserId int) (*[]int, error) {
	var hardTask []int
	var newPostsData []interface{}
	newPostsData = append(newPostsData, UserId)

	sql := `SELECT "task_id" FROM "difficulty_task" WHERE "user_id" = $1 and "difficulty" = 1`

	rows, err := r.DB.Query(sql, newPostsData...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var buff int
		err = rows.Scan(
			&buff,
		)
		if err != nil {
			return nil, err
		}
		//fmt.Println(buff)

		hardTask = append(hardTask, buff)
	}

	return &hardTask, nil
}

func (r *RepositoryML) GetDoneTask(id int) (*[]int, error) {
	var doneTask []int
	var newPostsData []interface{}
	newPostsData = append(newPostsData, id)

	sql := `select DISTINCT task_id from send_task where user_id = $1 and "tests_passed" = "tests_total";`

	rows, err := r.DB.Query(sql, newPostsData...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var buff int
		err = rows.Scan(
			&buff,
		)
		if err != nil {
			return nil, err
		}
		fmt.Println(buff)

		doneTask = append(doneTask, buff)
	}

	return &doneTask, nil
}

func (r *RepositoryML) GetTaskById(id int) (Task models.TaskDB, err error) {
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

func (r *RepositoryML) GetSendTaskByTaskId(UserId int, TaskId int) (*models.SendTasks, error) {
	Task1 := &models.SendTasks{}

	var newPostsData []interface{}
	newPostsData = append(newPostsData, UserId)
	newPostsData = append(newPostsData, TaskId)
	//sql := `SELECT "id", "user_id", "task_id", "check_time", "build_time", "check_result", "check_message", "tests_passed", "tests_total", "lint_success", "code_text", "date"
	//	FROM "send_task" where "user_id" = $1 and "task_id" = $2`

	sql := `SELECT * FROM "send_task" where "user_id" = $1 and "task_id" = $2`

	rows, err := r.DB.Query(sql, newPostsData...)
	if err != nil {
		//return nil, err
		return nil, errors.Errorf("GetSendTaskByTaskId1" + err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var buff models.SendTask
		err = rows.Scan(
			&buff.ID,
			&buff.UserId,
			&buff.TaskId,
			&buff.CheckTime,
			&buff.BuildTime,
			&buff.CheckResult,
			&buff.CheckMessage,
			&buff.TestsPassed,
			&buff.TestsTotal,
			&buff.LintSuccess,
			&buff.CodeText,
			&buff.Date,
		)

		if err != nil {
			//return nil, err
			return nil, errors.Errorf("GetSendTaskByTaskId2" + " " + err.Error() + " " + fmt.Sprint(UserId) + " " + fmt.Sprint(TaskId))
		}

		Task1.Tasks = append(Task1.Tasks, buff)
	}

	return Task1, nil
}

func (r *RepositoryML) UpdateUserColdStart(UserId int) error {
	var newPostsData []interface{}
	newPostsData = append(newPostsData, UserId)

	sql := `UPDATE "users" SET "cold_start" = 'true' WHERE "id"=$1`

	rows, err := r.DB.Query(sql, newPostsData...)
	if err != nil {
		return err
	}
	defer rows.Close()
	return nil
}
