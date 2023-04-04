package task

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/tp-study-ai/backend/internal/app/models"
	"github.com/tp-study-ai/backend/tools"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
)

type HandlerTask struct {
	UseCase UseCase
}

func NewHandlerTask(useCase UseCase) *HandlerTask {
	return &HandlerTask{
		UseCase: useCase,
	}
}

func (h HandlerTask) GetTask(ctx echo.Context) error {
	task, err := h.UseCase.GetTask()
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

	task, err := h.UseCase.GetTaskById(int(che))
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

func (h HandlerTask) GetTaskByLimit(ctx echo.Context) error {
	id := ctx.QueryParam("id")
	fmt.Println("Param: ", id, " ", reflect.TypeOf(id))
	che, _ := strconv.ParseInt(id, 10, 64)

	tasks, err := h.UseCase.GetTaskByLimit(int(che))
	if err != nil {
		return tools.CustomError(ctx, err, 1, "что-то сломалось")
	}

	result, _ := json.Marshal(tasks)
	ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
	return ctx.JSONBlob(http.StatusOK, result)
}

func (h HandlerTask) CheckSolution(ctx echo.Context) error {
	var solution models.CheckSolutionRequest
	if err := ctx.Bind(&solution); err != nil {
		//fmt.Println(solution)
		//che := models.CustomError{
		//	Number: 1,
		//	Error:  err.Error(),
		//}
		//result, _ := json.Marshal(che)
		//ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
		//return ctx.JSONBlob(http.StatusInternalServerError, result)
		return tools.CustomError(ctx, err, 1, "")
	}

	cheche, err := h.UseCase.CheckSolution(solution)

	result, err := json.Marshal(cheche)
	if err != nil {
		che := models.CustomError{
			Number: 2,
			Error:  err.Error(),
		}
		result, _ := json.Marshal(che)
		ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
		return ctx.JSONBlob(http.StatusInternalServerError, result)
	}

	ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
	return ctx.JSONBlob(http.StatusOK, result)
}

func (h HandlerTask) MySendSolution(ctx echo.Context) error {
	var solution models.SS123
	solution.Tests = make([][]string, 0)

	if err := ctx.Bind(&solution); err != nil {
		fmt.Println(solution)
		che := models.CustomError{
			Number: 1,
			Error:  err.Error(),
		}
		result, _ := json.Marshal(che)
		ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
		return ctx.JSONBlob(http.StatusInternalServerError, result)
	}

	result, err := json.Marshal(solution)
	if err != nil {
		che := models.CustomError{
			Number: 2,
			Error:  err.Error(),
		}
		result, _ := json.Marshal(che)
		ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
		return ctx.JSONBlob(http.StatusInternalServerError, result)
	}

	responseBody := bytes.NewBuffer(result)
	resp, err := http.Post("http://95.163.214.80:8080/check_solution?api_key=secret_key_here", "application/json", responseBody)
	if err != nil {
		che := models.CustomError{
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
		che := models.CustomError{
			Number: 4,
			Error:  err.Error(),
		}
		result, _ := json.Marshal(che)
		ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
		return ctx.JSONBlob(http.StatusInternalServerError, result)
	}
	sb := string(body)
	fmt.Printf(sb)

	ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(body)))
	return ctx.JSONBlob(http.StatusOK, body)
}
