package mlUseCase

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	che "github.com/tp-study-ai/backend/internal/app/ml"
	"github.com/tp-study-ai/backend/internal/app/models"
	"io/ioutil"
	"net/http"
)

type UseCaseML struct {
	Repo    che.Repository
	Secret1 string
	Secret2 string
	Secret3 string
	Secret4 string
	Secret5 string
}

func NewUseCaseML(TaskRepo che.Repository, secret string, secret1 string, secret2 string, secret3 string, secret4 string) *UseCaseML {
	return &UseCaseML{
		Repo:    TaskRepo,
		Secret1: secret,
		Secret2: secret1,
		Secret3: secret2,
		Secret4: secret3,
		Secret5: secret4,
	}
}

func (u *UseCaseML) GetSimilar(solution models.SimilarRequest) (*models.Tasks, error) {
	solution.NRecs = 6

	result, err := json.Marshal(solution)
	if err != nil {
		return nil, err
	}

	req := bytes.NewBuffer(result)
	resp, err := http.Post(u.Secret2, "application/json", req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var TestisResponse []models.MlTaskResponse

	err = json.Unmarshal(body, &TestisResponse)
	if err != nil {
		return nil, err
	}

	fmt.Println(TestisResponse)

	Tasks := &models.Tasks{}

	for i := 0; i < len(TestisResponse); i++ {
		task, err1 := u.Repo.GetTaskByLink("https://codeforces.com" + TestisResponse[i].ProblemUrl + "?locale=ru")
		if err1 != nil {
			continue
			//return nil, err1
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

		Tasks.Tasks = append(Tasks.Tasks, models.TaskJSON{
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

	return Tasks, err
}

func (u *UseCaseML) Recommendations(UserId int) (*models.RecResponse, error) {
	easyTask, err := u.Repo.GetEasyTasksForUser(UserId)
	fmt.Println("GetEasyTasksForUser", easyTask)

	hardTask, err := u.Repo.GetHardTasksForUser(UserId)
	fmt.Println("GetHardTasksForUser", *hardTask)

	doneTask, err := u.Repo.GetDoneTask(UserId)
	fmt.Println("GetDoneTask", doneTask)

	newEasyTask := make([]int, 0)

	for _, at := range *easyTask {
		che := false
		for _, dt := range *doneTask {
			if at == dt {
				che = true
			}
		}
		if che == false {
			newEasyTask = append(newEasyTask, at)
		}
	}

	newHardTask := make([]int, 0)

	for _, at := range *hardTask {
		che := false
		for _, dt := range *doneTask {
			if at == dt {
				che = true
			}
		}
		if che == false {
			newHardTask = append(newHardTask, at)
		}
	}

	Story := &models.Story1{}

	if len(newEasyTask) != 0 {
		for _, item := range newEasyTask {
			var buff models.StoryItem1

			task, err1 := u.Repo.GetTaskById(item)
			if err1 != nil {
				return nil, err1
			}

			buff.ProblemUrl = task.ShortLink
			buff.Rating = task.CfRating

			buff.ProblemUrl = task.ShortLink
			buff.Rating = task.CfRating

			tagsId := make([]int, 0)
			if task.CfTags.Elements[0].Int != 0 {
				for i := 0; i < len(task.CfTags.Elements); i++ {
					tagsId = append(tagsId, int(task.CfTags.Elements[i].Int))
				}
			}

			buff.Tags = tagsId

			counterAttention := 0

			submisTask, err2 := u.Repo.GetSendTaskByTaskId(UserId, item)
			if err2 != nil {
				buff.NAttempts = counterAttention
			} else {
				for _, i := range submisTask.Tasks {
					if i.TestsTotal == i.TestsPassed && i.TestsTotal != 0 {
						break
					} else {
						counterAttention += 1
					}
				}

				buff.NAttempts = counterAttention
			}

			Story.TooEasy = append(Story.TooEasy, buff)
		}
	} else {
		Story.TooEasy = make([]models.StoryItem1, 0)
	}

	if len(newHardTask) != 0 {
		for _, item := range newHardTask {
			var buff models.StoryItem1

			task, err1 := u.Repo.GetTaskById(item)
			if err1 != nil {
				return nil, err1
			}

			buff.ProblemUrl = task.ShortLink
			buff.Rating = task.CfRating

			buff.ProblemUrl = task.ShortLink
			buff.Rating = task.CfRating

			tagsId := make([]int, 0)
			if task.CfTags.Elements[0].Int != 0 {
				for i := 0; i < len(task.CfTags.Elements); i++ {
					tagsId = append(tagsId, int(task.CfTags.Elements[i].Int))
				}
			}

			buff.Tags = tagsId

			counterAttention := 0

			submisTask, err2 := u.Repo.GetSendTaskByTaskId(UserId, item)
			if err2 != nil {
				buff.NAttempts = counterAttention
			} else {
				for _, i := range submisTask.Tasks {
					if i.TestsTotal == i.TestsPassed && i.TestsTotal != 0 {
						break
					} else {
						counterAttention += 1
					}
				}

				buff.NAttempts = counterAttention
			}

			Story.TooHard = append(Story.TooHard, buff)
		}
	} else {
		Story.TooHard = make([]models.StoryItem1, 0)
	}

	if len(*doneTask) != 0 {
		for _, item := range *doneTask {
			var buff models.StoryItem1

			task, err1 := u.Repo.GetTaskById(item)
			if err1 != nil {
				return nil, err1
			}

			buff.ProblemUrl = task.ShortLink
			buff.Rating = task.CfRating

			buff.ProblemUrl = task.ShortLink
			buff.Rating = task.CfRating

			tagsId := make([]int, 0)
			if task.CfTags.Elements[0].Int != 0 {
				for i := 0; i < len(task.CfTags.Elements); i++ {
					tagsId = append(tagsId, int(task.CfTags.Elements[i].Int))
				}
			}

			buff.Tags = tagsId

			counterAttention := 0

			submisTask, err2 := u.Repo.GetSendTaskByTaskId(UserId, item)
			if err2 != nil {
				buff.NAttempts = counterAttention
			} else {
				for _, i := range submisTask.Tasks {
					if i.TestsTotal == i.TestsPassed && i.TestsTotal != 0 {
						break
					} else {
						counterAttention += 1
					}
				}

				buff.NAttempts = counterAttention
			}

			Story.Solved = append(Story.Solved, buff)
		}
	} else {
		Story.Solved = make([]models.StoryItem1, 0)
	}

	result, err := json.Marshal(Story)
	if err != nil {
		return nil, err
	}

	req := bytes.NewBuffer(result)
	resp, err := http.Post(u.Secret3, "application/json", req)
	if err != nil {
		return nil, errors.Errorf("1127 " + err.Error())
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(body))

	var MlResponse models.Rec

	err = json.Unmarshal(body, &MlResponse.Rec)
	if err != nil {
		return nil, errors.Errorf("1141 " + err.Error() + " " + string(body) + " " + string(result) + " " + fmt.Sprint(doneTask) + " " + fmt.Sprint(newEasyTask) + " " + fmt.Sprint(newHardTask))
	}

	RecommendationResponse := &models.RecResponse{}

	for _, itemRec := range MlResponse.Rec {
		var buff models.RecommendedResponse
		buff.RecommendedTag = models.TagDict[itemRec.RecommendedTag][1]
		buff.Priority = itemRec.Priority

		for i, itemRecTask := range itemRec.Problems {
			var task models.TaskDB
			if i == 0 {
				task, err = u.Repo.GetTaskByLink("https://codeforces.com/contest/1373/problem/B?locale=ru")
			} else {
				task, err = u.Repo.GetTaskByLink("https://codeforces.com" + itemRecTask.ProblemUrl + "?locale=ru")
			}
			if err != nil {
				continue
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

			buff.Problems = append(buff.Problems, models.TaskJSON{
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

		if buff.Problems == nil {
			continue
		}
		RecommendationResponse.Rec = append(RecommendationResponse.Rec, buff)
	}

	return RecommendationResponse, nil
}

func (u *UseCaseML) ColdStart(UserId int) (*models.ColdStartResponse, error) {
	easyTask, _ := u.Repo.GetEasyTasksForUser(UserId)
	fmt.Println("difficultyTask", easyTask)

	hardTask, _ := u.Repo.GetHardTasksForUser(UserId)
	fmt.Println("submissionTask", hardTask)

	doneTask, _ := u.Repo.GetDoneTask(UserId)
	fmt.Println("submissionTask", doneTask)

	newEasyTask := make([]int, 0)

	for _, at := range *easyTask {
		che := false
		for _, dt := range *doneTask {
			if at == dt {
				che = true
			}
		}
		if che == false {
			newEasyTask = append(newEasyTask, at)
		}
	}

	newHardTask := make([]int, 0)

	for _, at := range *hardTask {
		che := false
		for _, dt := range *doneTask {
			if at == dt {
				che = true
			}
		}
		if che == false {
			newHardTask = append(newHardTask, at)
		}
	}

	Story := &models.Story1{}

	if len(newEasyTask) != 0 {
		for _, item := range newEasyTask {
			var buff models.StoryItem1

			task, err1 := u.Repo.GetTaskById(item)
			if err1 != nil {
				return nil, err1
			}

			buff.ProblemUrl = task.ShortLink
			buff.Rating = task.CfRating

			buff.ProblemUrl = task.ShortLink
			buff.Rating = task.CfRating

			tagsId := make([]int, 0)
			if task.CfTags.Elements[0].Int != 0 {
				for i := 0; i < len(task.CfTags.Elements); i++ {
					tagsId = append(tagsId, int(task.CfTags.Elements[i].Int))
				}
			}

			buff.Tags = tagsId

			counterAttention := 0

			submisTask, err2 := u.Repo.GetSendTaskByTaskId(UserId, item)
			if err2 != nil {
				buff.NAttempts = counterAttention
			} else {
				for _, i := range submisTask.Tasks {
					if i.TestsTotal == i.TestsPassed && i.TestsTotal != 0 {
						break
					} else {
						counterAttention += 1
					}
				}

				buff.NAttempts = counterAttention
			}

			Story.TooEasy = append(Story.TooEasy, buff)
		}
	} else {
		Story.TooEasy = make([]models.StoryItem1, 0)
	}

	if len(newHardTask) != 0 {
		for _, item := range newHardTask {
			var buff models.StoryItem1

			task, err1 := u.Repo.GetTaskById(item)
			if err1 != nil {
				return nil, err1
			}

			buff.ProblemUrl = task.ShortLink
			buff.Rating = task.CfRating

			buff.ProblemUrl = task.ShortLink
			buff.Rating = task.CfRating

			tagsId := make([]int, 0)
			if task.CfTags.Elements[0].Int != 0 {
				for i := 0; i < len(task.CfTags.Elements); i++ {
					tagsId = append(tagsId, int(task.CfTags.Elements[i].Int))
				}
			}

			buff.Tags = tagsId

			counterAttention := 0

			submisTask, err2 := u.Repo.GetSendTaskByTaskId(UserId, item)
			if err2 != nil {
				buff.NAttempts = counterAttention
			} else {
				for _, i := range submisTask.Tasks {
					if i.TestsTotal == i.TestsPassed && i.TestsTotal != 0 {
						break
					} else {
						counterAttention += 1
					}
				}

				buff.NAttempts = counterAttention
			}

			Story.TooHard = append(Story.TooHard, buff)
		}
	} else {
		Story.TooHard = make([]models.StoryItem1, 0)
	}

	if len(*doneTask) != 0 {
		for _, item := range *doneTask {
			var buff models.StoryItem1

			task, err1 := u.Repo.GetTaskById(item)
			if err1 != nil {
				return nil, err1
			}

			buff.ProblemUrl = task.ShortLink
			buff.Rating = task.CfRating

			buff.ProblemUrl = task.ShortLink
			buff.Rating = task.CfRating

			tagsId := make([]int, 0)
			if task.CfTags.Elements[0].Int != 0 {
				for i := 0; i < len(task.CfTags.Elements); i++ {
					tagsId = append(tagsId, int(task.CfTags.Elements[i].Int))
				}
			}

			buff.Tags = tagsId

			counterAttention := 0

			submisTask, err2 := u.Repo.GetSendTaskByTaskId(UserId, item)
			if err2 != nil {
				buff.NAttempts = counterAttention
			} else {
				for _, i := range submisTask.Tasks {
					if i.TestsTotal == i.TestsPassed && i.TestsTotal != 0 {
						break
					} else {
						counterAttention += 1
					}
				}

				buff.NAttempts = counterAttention
			}

			Story.Solved = append(Story.Solved, buff)
		}
	} else {
		Story.Solved = make([]models.StoryItem1, 0)
	}

	fmt.Println(Story)

	result, err := json.Marshal(Story)
	if err != nil {
		return nil, err
	}

	req := bytes.NewBuffer(result)
	resp, err := http.Post(u.Secret4, "application/json", req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(body))

	var ColdStartML models.ColdStartML

	err = json.Unmarshal(body, &ColdStartML)
	if err != nil {
		return nil, err
	}

	if ColdStartML.Finished == true {
		err1 := u.Repo.UpdateUserColdStart(UserId)
		if err1 != nil {
			return nil, err1
		}
		return &models.ColdStartResponse{
			Finished: ColdStartML.Finished,
			Progress: make([]models.Progress, 0),
		}, nil
	}

	task, err := u.Repo.GetTaskByLink("https://codeforces.com" + ColdStartML.ProblemUrl + "?locale=ru")
	if err != nil {
		return nil, err
	}

	var tagsId []int
	var tagsEn []string
	var tagsRu []string

	if task.CfTags.Elements[0].Int != 0 {
		for i := 0; i < len(task.CfTags.Elements); i++ {
			tagsId = append(tagsId, int(task.CfTags.Elements[i].Int))
			tagsEn = append(tagsEn, models.TagDict[tagsId[i]][0])
			tagsRu = append(tagsRu, models.TagDict[tagsId[i]][1])
		}
	}

	task1 := models.TaskJSON{
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
	}

	response := &models.ColdStartResponse{
		Finished: false,
		Progress: ColdStartML.Progress,
		Task:     task1,
	}

	return response, nil
}
