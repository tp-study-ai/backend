package task

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"reflect"
	"strconv"
)

type HandlerTask struct {
	Ucase         Ucase
}

func NewHandlerTask(ucase Ucase) *HandlerTask {
	return &HandlerTask{
		Ucase:         ucase,
	}
}

func (h HandlerTask) GetTask(ctx echo.Context) error {
	task, err := h.Ucase.GetTask()
	if err != nil {
		fmt.Println(err.Error())
	}

	result, _ := json.Marshal(task)
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

	result, _ := json.Marshal(task)
	ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
	return ctx.JSONBlob(http.StatusOK, result)
}