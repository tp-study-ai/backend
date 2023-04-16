package task

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/tp-study-ai/backend/internal/app/models"
	"io/ioutil"
	"net/http"
)

type UseCaseTask struct {
	Repo    Repository
	Secret1 string
	Secret2 string
}

func NewUseCaseTask(TaskRepo Repository, secret string, secret1 string) *UseCaseTask {
	return &UseCaseTask{
		Repo:    TaskRepo,
		Secret1: secret,
		Secret2: secret1,
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
		TaskRu:           Task.TaskRu,
		Input:            Task.Input,
		Output:           Task.Output,
		Note:             Task.Note,
	}

	if err != nil {
		return
	}
	return
}

func (u *UseCaseTask) GetTaskById(id int) (task models.TaskJSON, err error) {
	Task, err := u.Repo.GetTaskById(id)

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
		TaskRu:           Task.TaskRu,
		Input:            Task.Input,
		Output:           Task.Output,
		Note:             Task.Note,
	}

	if err != nil {
		return
	}
	return
}

func (u *UseCaseTask) GetTaskByLimit(id int, sort string, tag []int) (*models.TasksPagination, error) {
	tasks, taskCount, err := u.Repo.GetTaskByLimit(id, sort, tag)
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

	if float64(PrivateTestsLength)*Task.TimeLimit > 300 {
		PrivateTestsLength = int(300 / Task.TimeLimit)
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
		TestTimeout:  Task.TimeLimit,
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

	//fmt.Println(body)

	var TestisResponse []models.MlTaskResponse

	err = json.Unmarshal(body, &TestisResponse)
	if err != nil {
		return nil, err
	}

	//fmt.Println(TestisResponse)

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
			TaskRu:           task.TaskRu,
			Input:            task.Input,
			Output:           task.Output,
			Note:             task.Note,
		})
	}

	return tasks, nil
}
