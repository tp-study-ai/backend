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

var TagDict = map[int][]string{
	1:  {"*special", "*особая задача"},
	2:  {"2-sat", "2-sat"},
	3:  {"binary search", "бинарный поиск"},
	4:  {"bitmasks", "битмаски"},
	5:  {"brute force", "перебор"},
	6:  {"chinese remainder theorem", "китайская теорема об остатках"},
	7:  {"combinatorics", "комбинаторика"},
	8:  {"constructive algorithms", "конструктив"},
	9:  {"data structures", "структуры данных"},
	10: {"dfs and similar", "поиск в глубину и подобное"},
	11: {"divide and conquer", "разделяй и властвуй"},
	12: {"dp", "дп"},
	13: {"dsu", "системы непересекающихся множеств"},
	14: {"expression parsing", "разбор выражений"},
	15: {"fft", "быстрое преобразование Фурье"},
	16: {"flows", "потоки"},
	17: {"games", "игры"},
	18: {"geometry", "геометрия"},
	19: {"graph matchings", "паросочетания"},
	20: {"graphs", "графы"},
	21: {"greedy", "жадные алгоритмы"},
	22: {"hashing", "хэши"},
	23: {"implementation", "реализация"},
	24: {"interactive", "интерактив"},
	25: {"math", "математика"},
	26: {"matrices", "матрицы"},
	27: {"meet-in-the-middle", "meet-in-the-middle"},
	28: {"number theory", "теория чисел"},
	29: {"probabilities", "теория вероятностей"},
	30: {"schedules", "расписания"},
	31: {"shortest paths", "кратчайшие пути"},
	32: {"sortings", "сортировки"},
	33: {"string suffix structures", "строковые суфф. структуры"},
	34: {"strings", "строки"},
	35: {"ternary search", "тернарный поиск"},
	36: {"trees", "деревья"},
	37: {"two pointers", "два указателя"},
}

func (u *UseCaseLike) LikeTask(like models.LikeJson) (err error) {
	_, err = u.Repo.GetLike(models.LikeDb{UserId: like.UserId, TaskId: like.TaskId})
	if err == nil {
		return errors.Errorf("такой лайк уже есть")
	}
	//if like.UserId == like1.UserId && like.TaskId == like1.TaskId {
	//	return errors.Errorf("такой лайк уже есть")
	//}
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
				//fmt.Println(tagsId)
				tagsEn = append(tagsEn, TagDict[tagsId[j]][0])
				//fmt.Println(tagsRu)
				tagsRu = append(tagsRu, TagDict[tagsId[j]][1])
				//fmt.Println(tagsEn)
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
