package task

import (
	"fmt"
	"github.com/tp-study-ai/backend/internal/app/models"
	"time"
)

type UseCaseTask struct {
	Repo    Repository
	Secret1 string
	Secret2 string
	Secret3 string
	Secret4 string
	Secret5 string
}

func NewUseCaseTask(TaskRepo Repository, secret string, secret1 string, secret2 string, secret3 string, secret4 string) *UseCaseTask {
	return &UseCaseTask{
		Repo:    TaskRepo,
		Secret1: secret,
		Secret2: secret1,
		Secret3: secret2,
		Secret4: secret3,
		Secret5: secret4,
	}
}

func (u *UseCaseTask) GetTask() (task models.TaskJSON, err error) {
	Task, err := u.Repo.GetTask()
	if err != nil {
		return
	}

	var tagsId []int
	var tagsEn []string
	var tagsRu []string

	if Task.CfTags.Elements[0].Int != 0 {
		for i := 0; i < len(Task.CfTags.Elements); i++ {
			tagsId = append(tagsId, int(Task.CfTags.Elements[i].Int))
			tagsEn = append(tagsEn, models.TagDict[tagsId[i]][0])
			tagsRu = append(tagsRu, models.TagDict[tagsId[i]][1])
		}
	}

	task = models.TaskJSON{
		Id:               Task.Id,
		Name:             Task.Name,
		Description:      Task.Description,
		PublicTests:      Task.PublicTests,
		Difficulty:       Task.Difficulty,
		CfContestId:      Task.CfContestId,
		CfIndex:          Task.CfIndex,
		CfPoints:         Task.CfPoints,
		CfRating:         Task.CfRating,
		CfTagsID:         tagsId,
		CfTagsRu:         tagsRu,
		CfTagsEN:         tagsEn,
		TimeLimit:        Task.TimeLimit,
		MemoryLimitBytes: Task.MemoryLimitBytes,
		Link:             Task.Link,
		ShortLink:        Task.ShortLink,
		NameRu:           Task.NameRu,
		TaskRu:           Task.TaskRu,
		Input:            Task.Input,
		Output:           Task.Output,
		Note:             Task.Note,
	}

	return
}

func (u *UseCaseTask) GetTaskById(id int) (task models.TaskJSON, err error) {
	Task, err := u.Repo.GetTaskById(id)
	if err != nil {
		return
	}

	var tagsId []int
	var tagsEn []string
	var tagsRu []string

	if Task.CfTags.Elements[0].Int != 0 {
		for i := 0; i < len(Task.CfTags.Elements); i++ {
			tagsId = append(tagsId, int(Task.CfTags.Elements[i].Int))
			tagsEn = append(tagsEn, models.TagDict[tagsId[i]][0])
			tagsRu = append(tagsRu, models.TagDict[tagsId[i]][1])
		}
	}

	task = models.TaskJSON{
		Id:               Task.Id,
		Name:             Task.Name,
		Description:      Task.Description,
		PublicTests:      Task.PublicTests,
		Difficulty:       Task.Difficulty,
		CfContestId:      Task.CfContestId,
		CfIndex:          Task.CfIndex,
		CfPoints:         Task.CfPoints,
		CfRating:         Task.CfRating,
		CfTagsID:         tagsId,
		CfTagsRu:         tagsRu,
		CfTagsEN:         tagsEn,
		TimeLimit:        Task.TimeLimit,
		MemoryLimitBytes: Task.MemoryLimitBytes,
		Link:             Task.Link,
		ShortLink:        Task.ShortLink,
		NameRu:           Task.NameRu,
		TaskRu:           Task.TaskRu,
		Input:            Task.Input,
		Output:           Task.Output,
		Note:             Task.Note,
	}

	return
}

func (u *UseCaseTask) GetTaskByLimit(id int, sort string, tag []int, minRating int, maxRating int) (*models.TasksPagination, error) {
	tasks, taskCount, err := u.Repo.GetTaskByLimit(id, sort, tag, minRating, maxRating)
	if err != nil {
		return nil, err
	}

	reqTasks := &models.TasksPagination{
		Tasks: make([]models.TaskJSON, len(tasks.Tasks)),
	}

	reqTasks.TaskCount = taskCount

	for i, task := range tasks.Tasks {
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

		reqTasks.Tasks[i] = models.TaskJSON{
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
	}

	return reqTasks, nil
}

func (u *UseCaseTask) GetSendTask(UserId int) (*models.SendTasksJson, error) {
	tasks, err := u.Repo.GetSendTask(UserId)
	if err != nil {
		return nil, err
	}

	reqTasks := &models.SendTasksJson{
		Tasks: make([]models.SendTaskJson, len(tasks.Tasks)),
	}

	for i, task := range tasks.Tasks {
		reqTasks.Tasks[i] = models.SendTaskJson{
			ID:           task.ID,
			UserId:       task.UserId,
			TaskId:       task.TaskId,
			CheckTime:    task.CheckTime,
			BuildTime:    task.BuildTime,
			CheckResult:  task.CheckResult,
			CheckMessage: task.CheckMessage,
			TestsPassed:  task.TestsPassed,
			TestsTotal:   task.TestsTotal,
			LintSuccess:  task.LintSuccess,
			CodeText:     task.CodeText,
			Date:         task.Date,
		}
	}

	return reqTasks, nil
}

func (u *UseCaseTask) GetSendTaskByTaskId(UserId int, TaskId int) (*models.SendTasksJson, error) {
	tasks, err := u.Repo.GetSendTaskByTaskId(UserId, TaskId)
	if err != nil {
		return nil, err
	}

	reqTasks := &models.SendTasksJson{
		Tasks: make([]models.SendTaskJson, len(tasks.Tasks)),
	}

	for i, task := range tasks.Tasks {
		reqTasks.Tasks[i] = models.SendTaskJson{
			ID:           task.ID,
			UserId:       task.UserId,
			TaskId:       task.TaskId,
			CheckTime:    task.CheckTime,
			BuildTime:    task.BuildTime,
			CheckResult:  task.CheckResult,
			CheckMessage: task.CheckMessage,
			TestsPassed:  task.TestsPassed,
			TestsTotal:   task.TestsTotal,
			LintSuccess:  task.LintSuccess,
			CodeText:     task.CodeText,
			Date:         task.Date,
		}
	}

	return reqTasks, nil
}

func (u *UseCaseTask) GetCountTaskOfDate(id int) (*models.Days, error) {
	now := time.Now()
	days := &models.Days{}
	for i := 0; i < 365; i++ {
		task, _ := u.Repo.GetCountTaskOfDate(id, now)
		days.Days = append(days.Days, models.Day{Day: now, Count: task})
		now = now.Add(-24 * time.Hour)
	}

	return days, nil
}

func (u *UseCaseTask) GetShockMode(id int) (*models.ShockMode, error) {
	shockMode := &models.ShockMode{}
	now := time.Now()
	a := 0
	for i := 0; i < 365; i++ {
		count, err := u.Repo.GetCountTaskOfDate(id, now)
		if count == 0 {
			break
		}
		if err != nil {
			return nil, err
		}
		a += 1
		now = now.Add(-24 * time.Hour)
	}

	if a != 0 {
		shockMode.Today = true
		shockMode.ShockMode = a
		return shockMode, nil
	} else {
		fmt.Println()
		now = now.Add(-24 * time.Hour)
		for i := 0; i < 60; i++ {
			count, err := u.Repo.GetCountTaskOfDate(id, now)
			if err != nil {
				return nil, err
			}
			if count == 0 {
				break
			}
			a += 1
			now = now.Add(-24 * time.Hour)
		}
	}

	shockMode.Today = false
	shockMode.ShockMode = a

	return shockMode, nil
}

func (u *UseCaseTask) GetNotDoneTask(id int) (*models.DoneTask, error) {
	doneTask, err := u.Repo.GetDoneTask(id)
	if err != nil {
		return nil, err
	}
	fmt.Println("doneTask:", doneTask)

	allTask, err := u.Repo.GetAllUserTask(id)
	if err != nil {
		return nil, err
	}
	fmt.Println("allTask:", allTask)

	var notDoneTask []int

	for _, at := range *allTask {
		che := false
		for _, dt := range *doneTask {
			if at == dt {
				che = true
			}
		}
		if che == false {
			notDoneTask = append(notDoneTask, at)
		}
	}

	fmt.Println("notDoneTask:", notDoneTask)

	notDoneTaskResponse := &models.DoneTask{}
	notDoneTaskResponse.CountDoneTask = len(notDoneTask)

	for _, taskId := range notDoneTask {
		var buff models.TaskDB
		buff, err = u.Repo.GetTaskById(taskId)
		if err != nil {
			return nil, err
		}

		var tagsId []int
		var tagsEn []string
		var tagsRu []string

		if buff.CfTags.Elements[0].Int != 0 {
			for i := 0; i < len(buff.CfTags.Elements); i++ {
				tagsId = append(tagsId, int(buff.CfTags.Elements[i].Int))
				tagsEn = append(tagsEn, models.TagDict[tagsId[i]][0])
				tagsRu = append(tagsRu, models.TagDict[tagsId[i]][1])
			}
		}

		notDoneTaskResponse.DoneTask = append(notDoneTaskResponse.DoneTask, models.TaskJSON{
			Id:               buff.Id,
			Name:             buff.Name,
			Description:      buff.Description,
			PublicTests:      buff.PublicTests,
			Difficulty:       buff.Difficulty,
			CfContestId:      buff.CfContestId,
			CfIndex:          buff.CfIndex,
			CfPoints:         buff.CfPoints,
			CfRating:         buff.CfRating,
			CfTagsID:         tagsId,
			CfTagsRu:         tagsRu,
			CfTagsEN:         tagsEn,
			TimeLimit:        buff.TimeLimit,
			MemoryLimitBytes: buff.MemoryLimitBytes,
			Link:             buff.Link,
			ShortLink:        buff.ShortLink,
			NameRu:           buff.NameRu,
			TaskRu:           buff.TaskRu,
			Input:            buff.Input,
			Output:           buff.Output,
			Note:             buff.Note,
		})
	}

	return notDoneTaskResponse, nil
}

func (u *UseCaseTask) GetDoneTask(id int) (*models.DoneTask, error) {
	doneTask := &models.DoneTask{}
	tasks, err := u.Repo.GetDoneTask(id)
	if err != nil {
		return nil, err
	}

	doneTask.CountDoneTask = len(*tasks)

	for _, taskId := range *tasks {
		var buff models.TaskDB
		buff, err = u.Repo.GetTaskById(taskId)
		if err != nil {
			return nil, err
		}

		var tagsId []int
		var tagsEn []string
		var tagsRu []string

		if buff.CfTags.Elements[0].Int != 0 {
			for i := 0; i < len(buff.CfTags.Elements); i++ {
				tagsId = append(tagsId, int(buff.CfTags.Elements[i].Int))
				tagsEn = append(tagsEn, models.TagDict[tagsId[i]][0])
				tagsRu = append(tagsRu, models.TagDict[tagsId[i]][1])
			}
		}

		doneTask.DoneTask = append(doneTask.DoneTask, models.TaskJSON{
			Id:               buff.Id,
			Name:             buff.Name,
			Description:      buff.Description,
			PublicTests:      buff.PublicTests,
			Difficulty:       buff.Difficulty,
			CfContestId:      buff.CfContestId,
			CfIndex:          buff.CfIndex,
			CfPoints:         buff.CfPoints,
			CfRating:         buff.CfRating,
			CfTagsID:         tagsId,
			CfTagsRu:         tagsRu,
			CfTagsEN:         tagsEn,
			TimeLimit:        buff.TimeLimit,
			MemoryLimitBytes: buff.MemoryLimitBytes,
			Link:             buff.Link,
			ShortLink:        buff.ShortLink,
			NameRu:           buff.NameRu,
			TaskRu:           buff.TaskRu,
			Input:            buff.Input,
			Output:           buff.Output,
			Note:             buff.Note,
		})
	}

	return doneTask, nil
}

func (u *UseCaseTask) SetDifficultyTask(difficulty models.DifficultyJson) error {
	err := u.Repo.SetDifficultyTask(models.DifficultyDb{
		UserId:     difficulty.UserId,
		TaskId:     difficulty.TaskId,
		Difficulty: difficulty.Difficulty,
	})
	if err != nil {
		return err
	}
	return nil
}
