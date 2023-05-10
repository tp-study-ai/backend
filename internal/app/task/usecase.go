package task

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/tp-study-ai/backend/internal/app/models"
	"io/ioutil"
	"net/http"
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
			tagsEn = append(tagsEn, TagDict[tagsId[i]][0])
			tagsRu = append(tagsRu, TagDict[tagsId[i]][1])
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
			tagsEn = append(tagsEn, TagDict[tagsId[i]][0])
			tagsRu = append(tagsRu, TagDict[tagsId[i]][1])
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
				tagsEn = append(tagsEn, TagDict[tagsId[j]][0])
				tagsRu = append(tagsRu, TagDict[tagsId[j]][1])
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

func (u *UseCaseTask) CheckSolution(solution models.CheckSolutionRequest, userId int) (*models.CheckSolutionUseCaseResponse, error) {
	Task, err := u.Repo.GetTaskById(solution.TaskId)
	if err != nil {
		return nil, err
	}

	PrivateTestsLength := len(Task.PrivateTests) / 4
	PrivateTestsBuffer := make([]string, 0)
	for _, value := range Task.PrivateTests {
		if value != "input" && value != "output" {
			PrivateTestsBuffer = append(PrivateTestsBuffer, value)
		}
	}

	if float64(PrivateTestsLength)*(Task.TimeLimit+1) > 300 {
		PrivateTestsLength = int(300/Task.TimeLimit + 1)
	}

	tests := make([][]string, PrivateTestsLength)

	for i := 0; i < PrivateTestsLength; i++ {
		tests[i] = make([]string, 2)
		tests[i][0] = PrivateTestsBuffer[i*2]
		tests[i][1] = PrivateTestsBuffer[i*2+1]
	}

	var SolutionReq = models.CheckSolution{
		SourceCode: models.SourceCode{
			Makefile: "solution: main.cpp\n\tg++ main.cpp -o solution\n\nrun: solution\n\t./solution",
			Main:     solution.Solution,
		},
		Tests:        tests,
		BuildTimeout: 10,
		TestTimeout:  Task.TimeLimit + 1,
	}

	result, err := json.Marshal(SolutionReq)
	if err != nil {
		return nil, err
	}

	req := bytes.NewBuffer(result)
	resp, err := http.Post(u.Secret1, "application/json", req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.Status == "401 UNAUTHORIZED" {
		return nil, errors.Errorf(string(body))
	}

	TestisResponse := &models.CheckSolutionUseCaseResponse{}

	err = json.Unmarshal(body, &TestisResponse)
	if err != nil {
		return nil, err
	}

	_, err = u.Repo.SendTask(&models.SendTask{
		ID:           0,
		UserId:       userId,
		TaskId:       solution.TaskId,
		CheckTime:    TestisResponse.CheckTime,
		BuildTime:    TestisResponse.BuildTime,
		CheckResult:  TestisResponse.CheckResult,
		CheckMessage: TestisResponse.CheckMessage,
		TestsPassed:  TestisResponse.TestsPassed,
		TestsTotal:   TestisResponse.TestsTotal,
		LintSuccess:  TestisResponse.LintSuccess,
		CodeText:     solution.Solution,
	})
	if err != nil {
		return nil, err
	}

	return TestisResponse, nil
}

func (u *UseCaseTask) GetSimilar(solution models.SimilarRequest) (*models.Tasks, error) {
	solution.NRecs = 6

	fmt.Println(solution)

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

	fmt.Println(string(body))

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
				tagsEn = append(tagsEn, TagDict[tagsId[j]][0])
				tagsRu = append(tagsRu, TagDict[tagsId[j]][1])
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

func (u *UseCaseTask) LikeTask(like models.LikeJson) (err error) {
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

func (u *UseCaseTask) DeleteLike(like models.LikeJson) (err error) {
	_, err = u.Repo.GetLike(models.LikeDb{UserId: like.UserId, TaskId: like.TaskId})
	if err != nil {
		return errors.Errorf("такого лайк нет")
	}
	err = u.Repo.DeleteLike(models.LikeDb{UserId: like.UserId, TaskId: like.TaskId})
	return
}

func (u *UseCaseTask) GetLikeTask(UserId models.UserId) (*models.LikeTasks, error) {
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
				tagsEn = append(tagsEn, TagDict[tagsId[i]][0])
				tagsRu = append(tagsRu, TagDict[tagsId[i]][1])
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
				tagsEn = append(tagsEn, TagDict[tagsId[i]][0])
				tagsRu = append(tagsRu, TagDict[tagsId[i]][1])
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

func (u *UseCaseTask) Recommendations(UserId int) (*models.RecResponse, error) {
	difficultyTask, err := u.Repo.GetSetDifficultyTasks(UserId)
	fmt.Println("difficultyTask", difficultyTask)

	submissionTask, err := u.Repo.GetAllUserTask(UserId)
	fmt.Println("submissionTask", submissionTask)

	var allTasks []int
	if len(*difficultyTask) != 0 && len(*submissionTask) != 0 {
		for _, item1 := range *difficultyTask {
			allTasks = append(allTasks, item1)
		}

		for _, at := range *submissionTask {
			che := false
			for _, dt := range *difficultyTask {
				if at == dt {
					che = true
				}
			}
			if che == false {
				allTasks = append(allTasks, at)
			}
		}
	} else {
		if len(*difficultyTask) != 0 {
			for _, item1 := range *difficultyTask {
				allTasks = append(allTasks, item1)
			}
		}
		if len(*submissionTask) != 0 {
			for _, item1 := range *submissionTask {
				allTasks = append(allTasks, item1)
			}
		}
	}

	fmt.Println(difficultyTask)
	fmt.Println(submissionTask)
	fmt.Println(allTasks)

	story := &models.Story{}
	story.UserId = UserId

	if len(allTasks) != 0 {
		for _, item := range allTasks {
			//fmt.Println("nice")
			//fmt.Println(item)
			var buff models.StoryItem
			task, err1 := u.Repo.GetTaskById(item)
			if err1 != nil {
				return nil, err1
			}

			buff.ProblemUrl = task.ShortLink
			buff.Rating = task.CfRating

			tagsId := make([]int, 0)
			if task.CfTags.Elements[0].Int != 0 {
				for i := 0; i < len(task.CfTags.Elements); i++ {
					tagsId = append(tagsId, int(task.CfTags.Elements[i].Int))
				}
			}

			buff.Tags = tagsId

			diff := false
			for _, i := range *difficultyTask {
				if i == item {
					diff = true
					break
				}
			}

			if diff {
				diffTask, err2 := u.Repo.GetSetDifficultyTask(UserId, item)
				//fmt.Println(diffTask)
				if err2 != nil {
					return nil, err2
				}
				buff.DifficultyMatch = diffTask.Difficulty
			} else {
				buff.DifficultyMatch = 0
			}

			submis := false
			for _, i := range *submissionTask {
				if i == item {
					submis = true
					break
				}
			}

			if submis {
				submisTask, err2 := u.Repo.GetSendTaskByTaskId(UserId, item)
				//fmt.Println(submisTask)
				if err2 != nil {
					return nil, err2
				}

				counterAttention := 0
				solveCheck := false

				for _, i := range submisTask.Tasks {
					if i.TestsTotal == i.TestsPassed && i.TestsTotal != 0 {
						solveCheck = true
						break
					} else {
						counterAttention += 1
					}
				}

				buff.Solved = solveCheck
				buff.NAttempts = counterAttention

			} else {
				buff.Solved = false
				buff.NAttempts = 0
			}

			story.Story = append(story.Story, buff)
		}
	}
	if story.Story == nil {
		story.Story = make([]models.StoryItem, 0)
	}

	result, err := json.Marshal(story)
	if err != nil {
		return nil, err
	}

	req := bytes.NewBuffer(result)
	resp, err := http.Post(u.Secret3, "application/json", req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	//fmt.Println(string(body))

	var MlResponse models.Rec

	err = json.Unmarshal(body, &MlResponse.Rec)
	if err != nil {
		return nil, errors.Errorf("876 " + err.Error() + " " + string(body) + " " + string(result) + " " + fmt.Sprint(difficultyTask) + " " + fmt.Sprint(submissionTask) + " " + fmt.Sprint(allTasks))
	}

	RecommendationResponse := &models.RecResponse{}

	for _, itemRec := range MlResponse.Rec {
		var buff models.RecommendedResponse
		buff.RecommendedTag = TagDict[itemRec.RecommendedTag][1]
		buff.Priority = itemRec.Priority

		for _, itemRecTask := range itemRec.Problems {
			task, err1 := u.Repo.GetTaskByLink("https://codeforces.com" + itemRecTask.ProblemUrl + "?locale=ru")
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
					tagsEn = append(tagsEn, TagDict[tagsId[j]][0])
					tagsRu = append(tagsRu, TagDict[tagsId[j]][1])
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

func (u *UseCaseTask) Recommendations1(UserId int) (*models.RecResponse, error) {
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
		buff.RecommendedTag = TagDict[itemRec.RecommendedTag][1]
		buff.Priority = itemRec.Priority

		for _, itemRecTask := range itemRec.Problems {
			task, err1 := u.Repo.GetTaskByLink("https://codeforces.com" + itemRecTask.ProblemUrl + "?locale=ru")
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
					tagsEn = append(tagsEn, TagDict[tagsId[j]][0])
					tagsRu = append(tagsRu, TagDict[tagsId[j]][1])
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

func (u *UseCaseTask) ColdStart(UserId int) (*models.ColdStartResponse, error) {
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
		return nil, errors.Errorf("1409 " + err.Error() + " " + string(body) + " " + string(result) + " " + fmt.Sprint(doneTask) + " " + fmt.Sprint(newEasyTask) + " " + fmt.Sprint(newHardTask))
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
		return nil, errors.Errorf(err.Error() + " 1423 " + ColdStartML.ProblemUrl)
	}

	var tagsId []int
	var tagsEn []string
	var tagsRu []string

	if task.CfTags.Elements[0].Int != 0 {
		for i := 0; i < len(task.CfTags.Elements); i++ {
			tagsId = append(tagsId, int(task.CfTags.Elements[i].Int))
			tagsEn = append(tagsEn, TagDict[tagsId[i]][0])
			tagsRu = append(tagsRu, TagDict[tagsId[i]][1])
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

func (u *UseCaseTask) Chat(Message models.ChatGPT) (*models.Message, error) {
	fmt.Println(Message.TaskId)
	task, err := u.Repo.GetTaskById(Message.TaskId)
	if err != nil {
		return nil, err
	}

	MessageRequest := &models.ChatGPTRequest{
		UserMessage:  Message.Message,
		Statement:    task.Description,
		UserSolution: Message.Code,
	}

	result, err := json.Marshal(MessageRequest)
	if err != nil {
		return nil, err
	}

	req := bytes.NewBuffer(result)
	resp, err := http.Post(u.Secret5, "application/json", req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(body))

	var ChatGPTResponse models.Message

	err = json.Unmarshal(body, &ChatGPTResponse)
	if err != nil {
		return nil, errors.Errorf("1511 " + err.Error())
	}

	return &ChatGPTResponse, nil
}
