package taskHandler

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/tp-study-ai/backend/internal/app/middleware"
	"github.com/tp-study-ai/backend/internal/app/models"
	"github.com/tp-study-ai/backend/internal/app/task"
	"github.com/tp-study-ai/backend/tools"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

type HandlerTask struct {
	UseCase task.UseCase
}

func NewHandlerTask(useCase task.UseCase) *HandlerTask {
	return &HandlerTask{
		UseCase: useCase,
	}
}

func (h HandlerTask) GetTask(ctx echo.Context) error {
	task, err := h.UseCase.GetTask()
	if err != nil {
		return tools.CustomError(ctx, err, 1, "ошибка полчения задачи")
	}

	result, err := json.Marshal(task)
	if err != nil {
		return tools.CustomError(ctx, err, 1, "ошибка формирования задачи")
	}
	ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
	return ctx.JSONBlob(http.StatusOK, result)
}

func (h HandlerTask) GetTaskById(ctx echo.Context) error {
	id := ctx.QueryParam("id")
	fmt.Println("Param: ", id, " ", reflect.TypeOf(id))
	che, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return tools.CustomError(ctx, err, 0, "невалидный id задачи")
	}

	task, err := h.UseCase.GetTaskById(int(che))
	if err != nil {
		return tools.CustomError(ctx, err, 1, "ошибка получения задачи")
	}

	result, err := json.Marshal(task)
	if err != nil {
		return tools.CustomError(ctx, err, 1, "ошибка формирования задачи")
	}
	ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
	return ctx.JSONBlob(http.StatusOK, result)
}

func (h HandlerTask) GetTaskByLimit(ctx echo.Context) error {
	page := ctx.QueryParam("page")
	//fmt.Printf("Param: %s, %s\n", page, reflect.TypeOf(page))
	pageInt, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		return tools.CustomError(ctx, err, 1, "не валидный id задачи")
	}

	sort := ctx.QueryParam("sort")
	//fmt.Printf("sort: %s %s\n", sort, reflect.TypeOf(sort))
	if sort == "rating_up" || sort == "rating_down" || sort == "" {
		fmt.Println("sort params is OK")
	} else {
		sort = ""
	}

	minRating := ctx.QueryParam("min_rating")
	var requestMinRatig int64
	if minRating == "" {
		requestMinRatig = 0
	} else {
		requestMinRatig, err = strconv.ParseInt(minRating, 10, 64)
		if err != nil {
			requestMinRatig = 0
		}
	}

	maxRating := ctx.QueryParam("min_rating")
	var requestMaxRatig int64
	if minRating == "" {
		requestMaxRatig = 0
	} else {
		requestMaxRatig, err = strconv.ParseInt(maxRating, 10, 64)
		if err != nil {
			requestMaxRatig = 0
		}
	}

	tags := ctx.QueryParam("tags")
	//fmt.Println(tags)
	var tagsInt []int
	if len(tags) != 0 {
		//tags = tags[1 : len(tags)-1]
		if len(tags) == 1 || len(tags) == 2 {
			tagInt, _ := strconv.ParseInt(tags, 10, 64)
			tagsInt = append(tagsInt, int(tagInt))
		} else {
			tags1 := strings.Split(tags, ",")
			//fmt.Println("list tags string", tags1)
			for _, item := range tags1 {
				tagInt, _ := strconv.ParseInt(item, 10, 64)
				tagsInt = append(tagsInt, int(tagInt))
			}
			//fmt.Println(tagsInt)
		}
	}

	tasks, err := h.UseCase.GetTaskByLimit(int(pageInt), sort, tagsInt, int(requestMinRatig), int(requestMaxRatig))
	if err != nil {
		return tools.CustomError(ctx, err, 1, "ошибка полчения задач")
	}

	result, err := json.Marshal(tasks)
	if err != nil {
		return tools.CustomError(ctx, err, 2, "ошибка формирования задач")
	}
	ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
	return ctx.JSONBlob(http.StatusOK, result)
}

func (h HandlerTask) GetTags(ctx echo.Context) error {
	tags := models.TagsJson{}

	for _, item := range models.MyTags {
		tags.Tags = append(tags.Tags, models.TagJson{TagsId: item.Id, TagsEn: item.Eng, TagsRu: item.Rus})
	}

	result, err := json.Marshal(tags)
	if err != nil {
		return tools.CustomError(ctx, err, 1, "ошибка формирвания тегов")
	}

	ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
	return ctx.JSONBlob(http.StatusOK, result)
}

func (h HandlerTask) GetSendTasks(ctx echo.Context) error {
	user := middleware.GetUserFromCtx(ctx)
	if user == nil {
		return tools.CustomError(ctx, errors.Errorf("пользователь не в системе"), 0, "ошибка при запросе пользователя")
	}

	tasks, err := h.UseCase.GetSendTask(int(user.Id))
	if err != nil {
		return tools.CustomError(ctx, err, 1, "ошибка получения посылок")
	}

	result, err := json.Marshal(tasks)
	if err != nil {
		return tools.CustomError(ctx, err, 2, "ошибка формирования посылок")
	}

	ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
	return ctx.JSONBlob(http.StatusOK, result)
}

func (h HandlerTask) GetSendTaskByTaskId(ctx echo.Context) error {
	user := middleware.GetUserFromCtx(ctx)
	if user == nil {
		return tools.CustomError(ctx, errors.Errorf("пользователь не в системе"), 0, "ошибка при запросе пользователя")
	}

	id := ctx.QueryParam("id")
	fmt.Println("Param: ", id, " ", reflect.TypeOf(id))
	taskId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return tools.CustomError(ctx, err, 0, "не валидный id задачи")
	}

	tasks, err := h.UseCase.GetSendTaskByTaskId(int(user.Id), int(taskId))
	if err != nil {
		result, err1 := json.Marshal(models.Message{Message: "задача еще не решалась пользователем"})
		if err1 != nil {
			return tools.CustomError(ctx, err, 2, "ошибка формирования посылок")
		}

		ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
		return ctx.JSONBlob(http.StatusOK, result)
	}

	result, err := json.Marshal(tasks)
	if err != nil {
		return tools.CustomError(ctx, err, 2, "ошибка формирования посылок")
	}

	ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
	return ctx.JSONBlob(http.StatusOK, result)
}

func (h HandlerTask) GetDoneTask(ctx echo.Context) error {
	user := middleware.GetUserFromCtx(ctx)
	if user == nil {
		return tools.CustomError(ctx, errors.Errorf("пользователь не в системе"), 0, "ошибка при запросе пользователя")
	}

	response, err := h.UseCase.GetDoneTask(int(user.Id))
	if err != nil {
		return tools.CustomError(ctx, err, 0, "ошибка получения решенных задач")
	}

	result, err := json.Marshal(response)
	if err != nil {
		return tools.CustomError(ctx, err, 2, "ошибка формирования задач")
	}

	ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
	return ctx.JSONBlob(http.StatusOK, result)
}

func (h HandlerTask) GetNotDoneTask(ctx echo.Context) error {
	user := middleware.GetUserFromCtx(ctx)
	if user == nil {
		return tools.CustomError(ctx, errors.Errorf("пользователь не в системе"), 0, "ошибка при запросе пользователя")
	}

	response, err := h.UseCase.GetNotDoneTask(int(user.Id))
	if err != nil {
		return tools.CustomError(ctx, err, 1, "ошибка получения начатых задач")
	}

	result, err := json.Marshal(response)
	if err != nil {
		return tools.CustomError(ctx, err, 2, "ошибка формирования задач")
	}

	ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
	return ctx.JSONBlob(http.StatusOK, result)
}

func (h HandlerTask) SetDifficultyTask(ctx echo.Context) error {
	user := middleware.GetUserFromCtx(ctx)
	if user == nil {
		return tools.CustomError(ctx, errors.Errorf("пользователь не в системе"), 0, "ошибка при запросе пользователя")
	}

	var like models.DifficultyJson
	if err := ctx.Bind(&like); err != nil {
		return tools.CustomError(ctx, err, 1, "ошибка формирования запроса")
	}

	like.UserId = int(user.Id)

	err := h.UseCase.SetDifficultyTask(like)
	if err != nil {
		return tools.CustomError(ctx, err, 0, "ошибка оценки задачи")
	}

	result, err := json.Marshal(models.Message{
		Message: "оценка поставленна",
	})
	if err != nil {
		return tools.CustomError(ctx, err, 2, "ошибка формирования ответа")
	}

	ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
	return ctx.JSONBlob(http.StatusOK, result)
}
