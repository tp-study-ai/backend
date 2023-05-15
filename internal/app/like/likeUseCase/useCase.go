package likeUseCase

import (
	"fmt"
	"github.com/pkg/errors"
	che "github.com/tp-study-ai/backend/internal/app/like"
	"github.com/tp-study-ai/backend/internal/app/models"
)

type UseCaseLike struct {
	Repo che.Repository
}

func NewUseCaseLike(likeRepo che.Repository) *UseCaseLike {
	return &UseCaseLike{
		Repo: likeRepo,
	}
}

func (u *UseCaseLike) LikeTask(like models.LikeJson) (err error) {
	_, err = u.Repo.GetLike(models.LikeDb{UserId: like.UserId, TaskId: like.TaskId})
	if err == nil {
		return errors.Errorf("такой лайк уже есть")
	}
	err = u.Repo.LikeTask(models.LikeDb{UserId: like.UserId, TaskId: like.TaskId})
	return
}

func (u *UseCaseLike) DeleteLike(like models.LikeJson) (err error) {
	_, err = u.Repo.GetLike(models.LikeDb{UserId: like.UserId, TaskId: like.TaskId})
	if err != nil {
		return errors.Errorf("такого лайк нет")
	}
	err = u.Repo.DeleteLike(models.LikeDb{UserId: like.UserId, TaskId: like.TaskId})
	return
}

func (u *UseCaseLike) GetLikeTask(UserId models.UserId) (*models.LikeTasks, error) {
	likes, err := u.Repo.GetLikes(UserId)
	if err != nil {
		return nil, err
	}

	tasks := &models.LikeTasks{}

	tasks.CountTasks = 0

	for _, like := range likes.Likes {
		fmt.Println(like.TaskId)
		task, err1 := u.Repo.GetTaskById(like.TaskId)
		fmt.Println(task)
		if err1 != nil {
			return nil, err
		}

		var tagsId []int
		var tagsEn []string
		var tagsRu []string

		if task.CfTags.Elements[0].Int != 0 {
			for j := 0; j < len(task.CfTags.Elements); j++ {
				tagsId = append(tagsId, int(task.CfTags.Elements[j].Int))
				tagsEn = append(tagsEn, models.TagDict[tagsId[j]][0])
				tagsRu = append(tagsRu, models.TagDict[tagsId[j]][1])
			}
		}

		tasks.CountTasks += 1
		tasks.TasksIdList = append(tasks.TasksIdList, task.Id)

		tasks.Tasks = append(tasks.Tasks, models.TaskJSON{
			Id:               task.Id,
			Name:             task.Name,
			Description:      task.Description,
			PublicTests:      task.PublicTests,
			Difficulty:       task.Difficulty,
			CfContestId:      task.CfContestId,
			CfIndex:          task.CfIndex,
			CfPoints:         task.CfPoints,
			CfRating:         task.CfRating,
			CfTagsID:         tagsId,
			CfTagsRu:         tagsRu,
			CfTagsEN:         tagsEn,
			TimeLimit:        task.TimeLimit,
			MemoryLimitBytes: task.MemoryLimitBytes,
			Link:             task.Link,
			ShortLink:        task.ShortLink,
			NameRu:           task.NameRu,
			TaskRu:           task.TaskRu,
			Input:            task.Input,
			Output:           task.Output,
			Note:             task.Note,
		})
	}

	return tasks, nil
}
