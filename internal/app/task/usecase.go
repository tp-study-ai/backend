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

	var Req = models.SourceCode{
		Makefile: "solution: main.c\n\tgcc main.c -o solution\nrun: solution\n\t./solution",
		Main:     UseCaseSolution.Solution,
	}

	che := make([][]string, 2)
	che[0] = make([]string, 2)
	che[0][0] = "1 2"
	che[0][1] = "3"

	var SolutionReq = models.CheckSolution{
		SourceCode:   Req,
		Tests:        che,
		BuildTimeout: 2,
		TestTimeout:  6,
	}

	result, err := json.Marshal(SolutionReq)
	if err != nil {
		return models.CheckSolutionUseCaseResponse{}, err
	}

	responseBody := bytes.NewBuffer(result)
	resp, err := http.Post("http://95.163.214.80:8080/check_solution?api_key=secret_key_here", "application/json", responseBody)
	if err != nil {
		return models.CheckSolutionUseCaseResponse{}, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return models.CheckSolutionUseCaseResponse{}, err
	}
	sb := string(body)
	fmt.Printf(sb)

	err = json.Unmarshal(body, cheche)
	if err != nil {
		return models.CheckSolutionUseCaseResponse{}, err
	}

	return cheche, nil
}
