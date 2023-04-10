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

func (u *UseCaseTask) GetTask() (task models.TaskJSON, err error) {
	Task, err := u.Repo.GetTask()

	var che []int32

	for i := 0; i < len(Task.CfTags.Elements); i++ {
		che = append(che, Task.CfTags.Elements[i].Int)
	}

	fmt.Println(che)

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
		CfTags:           che,
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

	var che []int32

	for i := 0; i < len(Task.CfTags.Elements); i++ {
		che = append(che, Task.CfTags.Elements[i].Int)
	}

	fmt.Println(che)

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
		CfTags:           che,
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
		var che []int32

		for i := 0; i < len(task.CfTags.Elements); i++ {
			che = append(che, task.CfTags.Elements[i].Int)
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
			CfTags:           che,
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
