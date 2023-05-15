package likeRepository

import (
	"fmt"
	"github.com/jackc/pgx"
	"github.com/tp-study-ai/backend/internal/app/models"
)

type RepositoryLike struct {
	DB *pgx.ConnPool
}

func NewRepositoryLike(db *pgx.ConnPool) *RepositoryLike {
	return &RepositoryLike{DB: db}
}

func (r *RepositoryLike) GetTaskById(id int) (*models.TaskDB, error) {
	Task := &models.TaskDB{}
	sql := `select * from "tasks" where id = $1;`

	err := r.DB.QueryRow(sql, id).Scan(
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
	if err != nil {
		return nil, err
	}

	return Task, nil
}

func (r *RepositoryLike) LikeTask(like models.LikeDb) (err error) {
	sql := `insert into "likes" ("user_id", "task_id") values ($1, $2);`
	var newPostsData []interface{}
	newPostsData = append(newPostsData, like.UserId)
	newPostsData = append(newPostsData, like.TaskId)
	_, err = r.DB.Query(sql, newPostsData...)
	return err
}

func (r *RepositoryLike) DeleteLike(like models.LikeDb) (err error) {
	sql := `DELETE FROM "likes" WHERE "user_id" = $1 AND "task_id" = $2;`
	var newPostsData []interface{}
	newPostsData = append(newPostsData, like.UserId)
	newPostsData = append(newPostsData, like.TaskId)
	_, err = r.DB.Query(sql, newPostsData...)
	return err
}

func (r *RepositoryLike) GetLike(like models.LikeDb) (*models.LikeDb, error) {
	like1 := &models.LikeDb{}
	sql := `SELECT "id", "user_id", "task_id" FROM "likes" WHERE "user_id" = $1 AND "task_id" = $2;`
	err := r.DB.QueryRow(sql, like.UserId, like.TaskId).Scan(&like1.Id, &like1.UserId, &like1.TaskId)
	if err != nil {
		return nil, err
	}
	return like1, nil
}

func (r *RepositoryLike) GetLikes(UserId models.UserId) (*models.LikesDb, error) {
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
