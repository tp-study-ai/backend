package task

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/tp-study-ai/backend/internal/app/models"
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
