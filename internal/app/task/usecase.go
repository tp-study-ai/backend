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

func (u *UseCaseTask) GetTask() (Task models.TaskResponse, err error) {
	Task, err = u.Repo.GetTask()

	if err != nil {
		return
	}
	return
}

func (u *UseCaseTask) GetTaskById(id int) (Task models.TaskResponse, err error) {
	Task, err = u.Repo.GetTaskById(id)

	if err != nil {
		return
	}
	return
}

func (u *UseCaseTask) CheckSolution(solution models.CheckSolutionRequest) (cheche models.CheckSolutionUseCaseResponse, err error) {
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

	fmt.Println(PrivateTestsBuffer)

	che := make([][]string, 5)

	for i := 0; i < 5; i++ {
		che[i] = make([]string, 2)
		che[i][0] = PrivateTestsBuffer[i*2]
		che[i][1] = PrivateTestsBuffer[i*2+1]
	}

	fmt.Println(che)

	var Req = models.SourceCode{
		Makefile: "solution: main.cpp\n\tg++ main.cpp -o solution\n\nrun: solution\n\t./solution",
		Main:     UseCaseSolution.Solution,
	}

	//che := make([][]string, 1)
	//che[0] = make([]string, 2)
	//che[0][0] = "1 2"
	//che[0][1] = "3"

	var SolutionReq = models.CheckSolution{
		SourceCode:   Req,
		Tests:        che,
		BuildTimeout: 10,
		TestTimeout:  6,
	}

	result, err := json.Marshal(SolutionReq)
	if err != nil {
		return models.CheckSolutionUseCaseResponse{}, err
	}

	responseBody := bytes.NewBuffer(result)
	//fmt.Println(responseBody)
	resp, err := http.Post("http://95.163.214.80:8080/check_solution?api_key=secret_key_here", "application/json", responseBody)
	if err != nil {
		return models.CheckSolutionUseCaseResponse{}, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return models.CheckSolutionUseCaseResponse{}, err
	}
	fmt.Printf(string(body))

	err = json.Unmarshal(body, &cheche)
	if err != nil {
		return models.CheckSolutionUseCaseResponse{}, err
	}

	return cheche, nil
}
