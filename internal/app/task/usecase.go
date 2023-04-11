package task

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/tp-study-ai/backend/internal/app/models"
	"io/ioutil"
	"net/http"
)

type UseCaseTask struct {
	Repo Repository
}

func NewUseCaseTask(TaskRepo Repository) *UseCaseTask {
	return &UseCaseTask{
		Repo: TaskRepo,
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
	var tagsRu []string
	var tagsEn []string

	for i := 0; i < len(Task.CfTags.Elements); i++ {
		tagsId = append(tagsId, int(Task.CfTags.Elements[i].Int))
		tagsRu = append(tagsRu, TagDict[tagsId[i]][0])
		tagsEn = append(tagsEn, TagDict[tagsId[i]][1])
	}

	task = models.TaskJSON{
		Id:               Task.Id,
		Name:             Task.Name,
		Description:      Task.Description,
		PublicTests:      Task.PublicTests,
		PrivateTests:     Task.PrivateTests,
		GeneratedTests:   Task.GeneratedTests,
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
	var tagsRu []string
	var tagsEn []string

	for i := 0; i < len(Task.CfTags.Elements); i++ {
		tagsId = append(tagsId, int(Task.CfTags.Elements[i].Int))
		tagsRu = append(tagsRu, TagDict[tagsId[i]][0])
		tagsEn = append(tagsEn, TagDict[tagsId[i]][1])
	}

	task = models.TaskJSON{
		Id:               Task.Id,
		Name:             Task.Name,
		Description:      Task.Description,
		PublicTests:      Task.PublicTests,
		PrivateTests:     Task.PrivateTests,
		GeneratedTests:   Task.GeneratedTests,
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

func (u *UseCaseTask) GetTaskByLimit(id int, sort string, tag []int) (*models.Tasks, error) {
	tasks, err := u.Repo.GetTaskByLimit(id, sort, tag)
	if err != nil {
		return nil, err
	}

	reqTasks := &models.Tasks{
		Tasks: make([]models.TaskJSON, len(tasks.Tasks)),
	}

	for i, task := range tasks.Tasks {
		var tagsId []int
		var tagsRu []string
		var tagsEn []string

		for i := 0; i < len(task.CfTags.Elements); i++ {
			tagsId = append(tagsId, int(task.CfTags.Elements[i].Int))
			tagsRu = append(tagsRu, TagDict[tagsId[i]][0])
			tagsEn = append(tagsEn, TagDict[tagsId[i]][1])
		}

		reqTasks.Tasks[i] = models.TaskJSON{
			Id:               task.Id,
			Name:             task.Name,
			Description:      task.Description,
			PublicTests:      task.PublicTests,
			PrivateTests:     task.PrivateTests,
			GeneratedTests:   task.GeneratedTests,
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

func (u *UseCaseTask) CheckSolution(solution models.CheckSolutionRequest) (*models.CheckSolutionUseCaseResponse, error) {
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
		TestTimeout:  1,
	}

	fmt.Println(SolutionReq)

	result, err := json.Marshal(SolutionReq)
	if err != nil {
		return nil, err
	}

	req := bytes.NewBuffer(result)
	resp, err := http.Post("http://146.185.208.233:8080/check_solution?api_key=secret_key_here", "application/json", req)
	if err != nil {
		return nil, err
	}

	fmt.Println(resp)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	fmt.Printf(string(body))
	TestisResponse := &models.CheckSolutionUseCaseResponse{}

	err = json.Unmarshal(body, &TestisResponse)
	if err != nil {
		return nil, err
	}

	fmt.Println(TestisResponse)

	_, err = u.Repo.SendTask(&models.SendTask{
		ID:           0,
		UserId:       0,
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
