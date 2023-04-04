package task

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/tp-study-ai/backend/internal/app/models"
	"github.com/tp-study-ai/backend/tools"
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
		return tools.CustomError(ctx, err, 1, "GetTask")
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
	che, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return tools.CustomError(ctx, err, 0, "ParseInt")
	}

	task, err := h.UseCase.GetTaskById(int(che))
	if err != nil {
		return tools.CustomError(ctx, err, 1, "GetTaskById")
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
	page := ctx.QueryParam("page")
	fmt.Println("Param: ", page, " ", reflect.TypeOf(page))
	che, _ := strconv.ParseInt(page, 10, 64)

	tasks, err := h.UseCase.GetTaskByLimit(int(che))
	if err != nil {
		return tools.CustomError(ctx, err, 1, "что-то сломалось GetTaskByLimit")
	}

	result, _ := json.Marshal(tasks)
	ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
	return ctx.JSONBlob(http.StatusOK, result)
}

func (h HandlerTask) CheckSolution(ctx echo.Context) error {
	var solution models.CheckSolutionRequest
	if err := ctx.Bind(&solution); err != nil {
		return tools.CustomError(ctx, err, 1, "")
	}

	testisResponse, err := h.UseCase.CheckSolution(solution)

	result, err := json.Marshal(testisResponse)
	if err != nil {
		return tools.CustomError(ctx, err, 2, "CheckSolution")
	}

	ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
	return ctx.JSONBlob(http.StatusOK, result)
}
