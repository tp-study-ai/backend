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

func (u *UseCaseTask) CheckSolution(solution models.CheckSolutionRequest) (models.CheckSolutionUseCaseResponse, error) {
	var UseCaseSolution = models.CheckSolutionUseCase{
		TaskId:   solution.TaskId,
		Solution: solution.Solution,
	}

	Task, err := u.Repo.GetTaskById(solution.TaskId)
	fmt.Println(Task.PrivateTests)
	PrivateTestsLength := len(Task.PrivateTests) / 4
	fmt.Println(PrivateTestsLength)
	PrivateTestsBuffer := make([]string, 0)
	for _, value := range Task.PrivateTests {
		if value != "input" && value != "output" {
			PrivateTestsBuffer = append(PrivateTestsBuffer, value)
		}
	}

	che := make([][]string, 1)

	for i := 0; i < 1; i++ {
		che[i] = make([]string, 2)
		che[i][0] = PrivateTestsBuffer[i*2]
		che[i][1] = PrivateTestsBuffer[i*2+1]
	}

	fmt.Println(che)

	var Req = models.SourceCode{
		Makefile: "solution: main.cpp\n\tg++ main.cpp -o solution\n\nrun: solution\n\t./solution",
		Main:     UseCaseSolution.Solution,
	}

	var SolutionReq = models.CheckSolution{
		SourceCode:   Req,
		Tests:        che,
		BuildTimeout: 10,
		TestTimeout:  10,
	}

	fmt.Println(SolutionReq)

	result, err := json.Marshal(SolutionReq)
	if err != nil {
		return models.CheckSolutionUseCaseResponse{}, err
	}

	responseBody := bytes.NewBuffer(result)
	//fmt.Println(responseBody)
	resp, err := http.Post("http://146.185.208.233:8080/check_solution?api_key=secret_key_here", "application/json", responseBody)
	if err != nil {
		return models.CheckSolutionUseCaseResponse{}, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return models.CheckSolutionUseCaseResponse{}, err
	}
	fmt.Printf(string(body))

	var cheche models.CheckSolutionUseCaseResponse

	err = json.Unmarshal(body, &cheche)
	if err != nil {
		return models.CheckSolutionUseCaseResponse{}, err
	}

	//send := &models.SendTask{
	//		ID:           0,
	//		UserId:       0,
	//		TaskId:       solution.TaskId,
	//		CheckTime:    cheche.CheckTime,
	//		BuildTime:    cheche.BuildTime,
	//		CheckResult:  cheche.CheckResult,
	//		CheckMessage: cheche.CheckMessage,
	//		TestsPassed:  cheche.TestsPassed,
	//		TestsTotal:   cheche.TestsTotal,
	//		LintSuccess:  cheche.LintSuccess,
	//		CodeText:     solution.Solution,
	//	}

	response1, err := u.Repo.SendTask(models.SendTask{
		ID:           0,
		UserId:       0,
		TaskId:       solution.TaskId,
		CheckTime:    cheche.CheckTime,
		BuildTime:    cheche.BuildTime,
		CheckResult:  cheche.CheckResult,
		CheckMessage: cheche.CheckMessage,
		TestsPassed:  cheche.TestsPassed,
		TestsTotal:   cheche.TestsTotal,
		LintSuccess:  cheche.LintSuccess,
		CodeText:     solution.Solution,
	})

	if solution.TaskId != response1.TaskId || solution.Solution != response1.CodeText {
		return models.CheckSolutionUseCaseResponse{}, errors.Errorf("")
	}

	return cheche, nil
}
