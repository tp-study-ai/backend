package task

import (
	"fmt"
	"github.com/jackc/pgx"
	"github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/tp-study-ai/backend/internal/app/models"
	"math/rand"
	"time"
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
		&Task.ShortLink,
		&Task.NameRu,
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
		&Task.ShortLink,
		&Task.NameRu,
		&Task.TaskRu,
		&Task.Input,
		&Task.Output,
		&Task.Note,
	)

	return
}

func (r *RepositoryTask) GetTaskByLimit(id int, sort string, tag []int, minRating int, maxRating int) (*models.TasksResponse, int, error) {
	tasks := &models.TasksResponse{}

	var ForTaskCount []interface{}
	var TaskCount int

	var newPostsData []interface{}
	newPostsData = append(newPostsData, 15)
	newPostsData = append(newPostsData, 15*id)

	sql := `select * from "tasks"`
	sqlCount := `select count(*) from "tasks"`

	if len(tag) != 0 {
		sql = sql + ` where $3 <@ (cf_tags)`
		sqlCount = sqlCount + ` where $1 <@ (cf_tags)`
		newPostsData = append(newPostsData, pq.Array(tag))
		ForTaskCount = append(ForTaskCount, pq.Array(tag))
	}

	if minRating == 0 && maxRating == 0 {
		fmt.Println("no rating")
	} else {
		if minRating > 0 && maxRating > 0 {
			if len(tag) != 0 {
				sql = sql + ` and $4 <= cf_rating and cf_rating <= $5`
				newPostsData = append(newPostsData, minRating)
				newPostsData = append(newPostsData, maxRating)
				sqlCount = sqlCount + ` and $2 <= cf_rating and cf_rating <= $3`
				ForTaskCount = append(ForTaskCount, minRating)
				ForTaskCount = append(ForTaskCount, maxRating)
			} else {
				sql = sql + ` where $3 <= cf_rating and cf_rating <= $4`
				newPostsData = append(newPostsData, minRating)
				newPostsData = append(newPostsData, maxRating)
				sqlCount = sqlCount + ` where $1 <= cf_rating and cf_rating <= $2`
				ForTaskCount = append(ForTaskCount, minRating)
				ForTaskCount = append(ForTaskCount, maxRating)
			}
		} else {
			if minRating > 0 {
				if len(tag) != 0 {
					sql = sql + ` and $4 <= cf_rating`
					newPostsData = append(newPostsData, minRating)
					sqlCount = sqlCount + ` and $2 <= cf_rating`
					ForTaskCount = append(ForTaskCount, minRating)
				} else {
					sql = sql + ` where $3 <= cf_rating`
					newPostsData = append(newPostsData, minRating)
					sqlCount = sqlCount + ` where $2 <= cf_rating`
					ForTaskCount = append(ForTaskCount, minRating)
				}
			} else {
				if maxRating > 0 {
					if len(tag) != 0 {
						sql = sql + ` and cf_rating <= $4`
						newPostsData = append(newPostsData, maxRating)
						sqlCount = sqlCount + ` and cf_rating <= $2`
						ForTaskCount = append(ForTaskCount, maxRating)
					} else {
						sql = sql + ` where cf_rating <= $3`
						newPostsData = append(newPostsData, maxRating)
						sqlCount = sqlCount + ` where cf_rating <= $1`
						ForTaskCount = append(ForTaskCount, maxRating)
					}
				}
			}
		}
	}

	if sort == "" {
		sql = sql + `order by "id"`
	}

	if sort == "rating_up" {
		sql = sql + ` order by "cf_rating"`
	}

	if sort == "rating_down" {
		sql = sql + ` order by "cf_rating" desc`
	}

	sql = sql + ` limit $1 offset $2`

	fmt.Println(sqlCount)
	fmt.Println(sql)

	err := r.DB.QueryRow(sqlCount, ForTaskCount...).Scan(&TaskCount)
	if err != nil {
		return nil, 0, err
	}

	rows, err := r.DB.Query(sql, newPostsData...)
	if err != nil {
		return nil, 0, err
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
			&buff.ShortLink,
			&buff.NameRu,
			&buff.TaskRu,
			&buff.Input,
			&buff.Output,
			&buff.Note,
		)

		if err != nil {
			return nil, 0, err
		}

		tasks.Tasks = append(tasks.Tasks, buff)
	}

	return tasks, TaskCount, nil
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

func (r *RepositoryTask) GetTaskByLink(link string) (Task models.TaskDB, err error) {
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
	)

	return
}

func (r *RepositoryTask) GetSendTask(UserId int) (*models.SendTasks, error) {
	Task1 := &models.SendTasks{}

	var newPostsData []interface{}
	newPostsData = append(newPostsData, UserId)
	sql := `SELECT "id", "user_id", "task_id", "check_time", "build_time", "check_result", "check_message", "tests_passed", "tests_total", "lint_success", "code_text", "date"
		FROM "send_task" where "user_id" = $1`

	rows, err := r.DB.Query(sql, newPostsData...)
	if err != nil {
		return nil, err
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
			return nil, err
		}

		Task1.Tasks = append(Task1.Tasks, buff)
	}

	return Task1, nil
}

func (r *RepositoryTask) GetSendTaskByTaskId(UserId int, TaskId int) (*models.SendTasks, error) {
	Task1 := &models.SendTasks{}

	var newPostsData []interface{}
	newPostsData = append(newPostsData, UserId)
	newPostsData = append(newPostsData, TaskId)
	//sql := `SELECT "id", "user_id", "task_id", "check_time", "build_time", "check_result", "check_message", "tests_passed", "tests_total", "lint_success", "code_text", "date"
	//	FROM "send_task" where "user_id" = $1 and "task_id" = $2`

	sql := `SELECT * FROM "send_task" where "user_id" = $1 and "task_id" = $2`

	rows, err := r.DB.Query(sql, newPostsData...)
	if err != nil {
		return nil, err
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
			return nil, err
		}

		Task1.Tasks = append(Task1.Tasks, buff)
	}

	return Task1, nil
}

func (r *RepositoryTask) LikeTask(like models.LikeDb) (err error) {
	fmt.Println("RepositoryTask", like)
	sql := `insert into "likes" ("user_id", "task_id") values ($1, $2);`
	var newPostsData []interface{}
	newPostsData = append(newPostsData, like.UserId)
	newPostsData = append(newPostsData, like.TaskId)
	_, err = r.DB.Query(sql, newPostsData...)
	return err
}

func (r *RepositoryTask) DeleteLike(like models.LikeDb) (err error) {
	fmt.Println("RepositoryTask", like)
	sql := `DELETE FROM "likes" WHERE "user_id" = $1 AND "task_id" = $2;`
	var newPostsData []interface{}
	newPostsData = append(newPostsData, like.UserId)
	newPostsData = append(newPostsData, like.TaskId)
	_, err = r.DB.Query(sql, newPostsData...)
	return err
}

func (r *RepositoryTask) GetLike(like models.LikeDb) (*models.LikeDb, error) {
	like1 := &models.LikeDb{}
	sql := `SELECT "id", "user_id", "task_id" FROM "likes" WHERE "user_id" = $1 AND "task_id" = $2;`
	err := r.DB.QueryRow(sql, like.UserId, like.TaskId).Scan(&like1.Id, &like1.UserId, &like1.TaskId)
	if err != nil {
		return nil, err
	}
	return like1, nil
}

func (r *RepositoryTask) GetLikes(UserId models.UserId) (*models.LikesDb, error) {
	Likes := &models.LikesDb{}
	var newPostsData []interface{}
	newPostsData = append(newPostsData, UserId)

	sql := `SELECT "id", "user_id", "task_id" FROM "likes" WHERE "user_id" = $1`

	rows, err := r.DB.Query(sql, newPostsData...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var buff models.LikeDb
		err = rows.Scan(
			&buff.Id,
			&buff.UserId,
			&buff.TaskId,
		)
		if err != nil {
			return nil, err
		}
		fmt.Println(buff)

		Likes.Likes = append(Likes.Likes, buff)
	}

	return Likes, nil
}

func (r *RepositoryTask) GetCountTaskOfDate(id int, day time.Time) (int, error) {
	var countTask int
	sql := `select count(*) from send_task where "user_id" = $1 and $2 < date::date + interval '3 hours' and date::date + interval '3 hours' < $3 and "tests_passed" = "tests_total";`
	err := r.DB.QueryRow(sql, id, day.Format("2006-01-02"), day.Add(+24*time.Hour).Format("2006-01-02")).Scan(
		&countTask,
	)
	if err != nil {
		return 0, err
	}
	return countTask, nil
}

// GetAllUserTask получение всех задач которые пользователь решал удачно или нет
func (r *RepositoryTask) GetAllUserTask(id int) (*[]int, error) {
	var allTask []int
	var newPostsData []interface{}
	newPostsData = append(newPostsData, id)

	sql := `select DISTINCT task_id from send_task where user_id = $1;`

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

		allTask = append(allTask, buff)
	}

	return &allTask, nil
}

// GetDoneTask получение всех задач которые пользователь решал удачно
func (r *RepositoryTask) GetDoneTask(id int) (*[]int, error) {
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

func (r *RepositoryTask) SetDifficultyTask(difficulty models.DifficultyDb) error {
	fmt.Println("difficulty", difficulty)
	sql := `insert into "difficulty_task" ("user_id", "task_id", "difficulty") values ($1, $2, $3);`
	var newPostsData []interface{}
	newPostsData = append(newPostsData, difficulty.UserId)
	newPostsData = append(newPostsData, difficulty.TaskId)
	newPostsData = append(newPostsData, difficulty.Difficulty)
	_, err := r.DB.Query(sql, newPostsData...)
	return err
}

func (r *RepositoryTask) GetSetDifficultyTasks(UserId int) (*[]int, error) {
	var doneTask []int
	var newPostsData []interface{}
	newPostsData = append(newPostsData, UserId)

	sql := `SELECT "task_id" FROM "difficulty_task" WHERE "user_id" = $1`

	rows, err := r.DB.Query(sql, newPostsData...)
	if err != nil {
		//return nil, err
		return nil, errors.Errorf("GetSetDifficultyTasks")
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

		doneTask = append(doneTask, buff)
	}

	return &doneTask, nil
}

func (r *RepositoryTask) GetSetDifficultyTask(UserId int, TaskId int) (*models.DifficultyDb, error) {
	doneTask := &models.DifficultyDb{}
	var newPostsData []interface{}
	newPostsData = append(newPostsData, UserId)
	newPostsData = append(newPostsData, TaskId)

	sql := `SELECT "user_id", "task_id", "difficulty" FROM "difficulty_task" WHERE "user_id" = $1 and "task_id" = $2`

	err := r.DB.QueryRow(sql, UserId, TaskId).Scan(
		&doneTask.UserId,
		&doneTask.TaskId,
		&doneTask.Difficulty,
	)
	if err != nil {
		//return nil, err
		return nil, errors.Errorf("GetSetDifficultyTask")
	}

	return doneTask, nil
}
