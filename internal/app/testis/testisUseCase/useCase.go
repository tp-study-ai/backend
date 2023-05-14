package testisUseCase

import (
	"bytes"
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/tp-study-ai/backend/internal/app/models"
	che "github.com/tp-study-ai/backend/internal/app/testis"
	"io/ioutil"
	"net/http"
)

type UseCaseTestis struct {
	Repo    che.Repository
	Secret1 string
	Secret2 string
	Secret3 string
	Secret4 string
	Secret5 string
}

func NewUseCaseTestis(TaskRepo che.Repository, secret string, secret1 string, secret2 string, secret3 string, secret4 string) *UseCaseTestis {
	return &UseCaseTestis{
		Repo:    TaskRepo,
		Secret1: secret,
		Secret2: secret1,
		Secret3: secret2,
		Secret4: secret3,
		Secret5: secret4,
	}
}

func (u *UseCaseTestis) CheckSolution(solution models.CheckSolutionRequest, userId int) (*models.CheckSolutionUseCaseResponse, error) {
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
