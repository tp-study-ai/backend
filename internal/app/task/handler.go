package task

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/tp-study-ai/backend/internal/app/models"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
)

type HandlerTask struct {
	Ucase Ucase
}

func NewHandlerTask(ucase Ucase) *HandlerTask {
	return &HandlerTask{
		Ucase: ucase,
	}
}

func (h HandlerTask) GetTask(ctx echo.Context) error {
	task, err := h.Ucase.GetTask()
	if err != nil {
		fmt.Println(err.Error())
	}

	task1 := models.Task{
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
		CfTags:           task.CfTags,
		TimeLimit:        task.TimeLimit,
		MemoryLimitBytes: task.MemoryLimitBytes,
		Link:             task.Link,
		TaskRu:           task.TaskRu,
		Input:            task.Input,
		Output:           task.Output,
		Note:             task.Note,
	}

	result, _ := json.Marshal(task1)
	ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
	return ctx.JSONBlob(http.StatusOK, result)
}

func (h HandlerTask) GetTaskById(ctx echo.Context) error {
	id := ctx.QueryParam("id")
	fmt.Println("Param: ", id, " ", reflect.TypeOf(id))
	che, _ := strconv.ParseInt(id, 10, 64)

	task, err := h.Ucase.GetTaskById(int(che))
	//task, err := h.Ucase.GetTask()
	if err != nil {
		fmt.Println(err.Error())
	}

	task1 := models.Task{
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
		CfTags:           task.CfTags,
		TimeLimit:        task.TimeLimit,
		MemoryLimitBytes: task.MemoryLimitBytes,
		Link:             task.Link,
		TaskRu:           task.TaskRu,
		Input:            task.Input,
		Output:           task.Output,
		Note:             task.Note,
	}

	result, _ := json.Marshal(task1)
	ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
	return ctx.JSONBlob(http.StatusOK, result)
}

type Test struct {
	che []string
}

type SourceCode struct {
	Makefile string `json:"Makefile"`
	Main     string `json:"main.c"`
	Main1    string `json:"lib/sum.c"`
	Main2    string `json:"lib/sum.h"`
}

type SS123 struct {
	SourceCode   SourceCode `json:"sourceCode"`
	Tests        [][]string `json:"tests"`
	BuildTimeout int        `json:"buildTimeout"`
	TestTimeout  int        `json:"testTimeout"`
}

type CustomError struct {
	Number int    `json:"number"`
	Error  string `json:"error"`
}

func (h HandlerTask) SendSolution(ctx echo.Context) error {
	var solution SS123
	solution.Tests = make([][]string, 0)

	if err := ctx.Bind(&solution); err != nil {
		fmt.Println(solution)
		che := CustomError{
			Number: 1,
			Error:  err.Error(),
		}
		result, _ := json.Marshal(che)
		ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
		return ctx.JSONBlob(http.StatusInternalServerError, result)
	}

	result, err := json.Marshal(solution)
	if err != nil {
		che := CustomError{
			Number: 2,
			Error:  err.Error(),
		}
		result, _ := json.Marshal(che)
		ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
		return ctx.JSONBlob(http.StatusInternalServerError, result)
	}

	responseBody := bytes.NewBuffer(result)
	resp, err := http.Post("http://127.0.0.1:8080/check_solution?api_key=secret_key_here", "application/json", responseBody)
	if err != nil {
		che := CustomError{
			Number: 3,
			Error:  err.Error(),
		}
		result, _ := json.Marshal(che)
		ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
		return ctx.JSONBlob(http.StatusInternalServerError, result)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		che := CustomError{
			Number: 4,
			Error:  err.Error(),
		}
		result, _ := json.Marshal(che)
		ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
		return ctx.JSONBlob(http.StatusInternalServerError, result)
	}
	sb := string(body)
	fmt.Printf(sb)

	ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
	return ctx.JSONBlob(http.StatusOK, result)
}
