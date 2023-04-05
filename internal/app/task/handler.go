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
	"strings"
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

	result, err := json.Marshal(task)
	if err != nil {
		return tools.CustomError(ctx, err, 1, "GetTask")
	}
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

	result, err := json.Marshal(task)
	if err != nil {
		return tools.CustomError(ctx, err, 1, "GetTaskById")
	}
	ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
	return ctx.JSONBlob(http.StatusOK, result)
}

func (h HandlerTask) GetTaskByLimit(ctx echo.Context) error {
	page := ctx.QueryParam("page")
	fmt.Printf("Param: %s, %s\n", page, reflect.TypeOf(page))
	pageInt, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		return tools.CustomError(ctx, err, 1, "Parse int")
	}

	sort := ctx.QueryParam("sort")
	fmt.Printf("sort: %s %s\n", sort, reflect.TypeOf(sort))
	if sort == "rating_up" || sort == "rating_down" || sort == "" {
		fmt.Println("sort params is OK")
	} else {
		sort = ""
	}

	tags := ctx.QueryParam("tags")
	fmt.Println(tags)
	var tagsInt []int
	if len(tags) != 0 {
		tags = tags[1 : len(tags)-1]
		if len(tags) == 1 || len(tags) == 2 {
			tagInt, _ := strconv.ParseInt(tags, 10, 64)
			tagsInt = append(tagsInt, int(tagInt))
		} else {
			tags1 := strings.Split(tags, ",")
			fmt.Println("list tags string", tags1)
			for _, item := range tags1 {
				tagInt, _ := strconv.ParseInt(item, 10, 64)
				tagsInt = append(tagsInt, int(tagInt))
			}
			fmt.Println(tagsInt)
		}
	}

	tasks, err := h.UseCase.GetTaskByLimit(int(pageInt), sort, tagsInt)
	if err != nil {
		return tools.CustomError(ctx, err, 1, "что-то сломалось GetTaskByLimit")
	}

	result, err := json.Marshal(tasks)
	if err != nil {
		return tools.CustomError(ctx, err, 1, "Marshal")
	}
	ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
	return ctx.JSONBlob(http.StatusOK, result)
}

func (h HandlerTask) CheckSolution(ctx echo.Context) error {
	var solution models.CheckSolutionRequest
	if err := ctx.Bind(&solution); err != nil {
		return tools.CustomError(ctx, err, 1, "CheckSolution Bind")
	}

	testisResponse, err := h.UseCase.CheckSolution(solution)

	result, err := json.Marshal(testisResponse)
	if err != nil {
		return tools.CustomError(ctx, err, 2, "CheckSolution Bind")
	}

	ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
	return ctx.JSONBlob(http.StatusOK, result)
}

type Tags struct {
	Id int
	Ru string
	En string
}

func (h HandlerTask) GetTags(ctx echo.Context) error {
	tags := models.TagsJson{}

	MyTags := []Tags{
		Tags{Id: 1, Ru: "*special", En: "*особая задача"},
		Tags{Id: 2, Ru: "2-sat", En: "2-sat"},
		Tags{Id: 3, Ru: "binary search", En: "бинарный поиск"},
		Tags{Id: 4, Ru: "bitmasks", En: "битмаски"},
		Tags{Id: 5, Ru: "brute force", En: "перебор"},
		Tags{Id: 6, Ru: "chinese remainder theorem", En: "китайская теорема об остатках"},
		Tags{Id: 7, Ru: "combinatorics", En: "комбинаторика"},
		Tags{Id: 8, Ru: "constructive algorithms", En: "конструктив"},
		Tags{Id: 9, Ru: "data structures", En: "структуры данных"},
		Tags{Id: 10, Ru: "dfs and similar", En: "поиск в глубину и подобное"},
		Tags{Id: 11, Ru: "divide and conquer", En: "разделяй и властвуй"},
		Tags{Id: 12, Ru: "dp", En: "дп"},
		Tags{Id: 13, Ru: "dsu", En: "системы непересекающихся множеств"},
		Tags{Id: 14, Ru: "expression parsing", En: "разбор выражений"},
		Tags{Id: 15, Ru: "fft", En: "быстрое преобразование Фурье"},
		Tags{Id: 16, Ru: "flows", En: "потоки"},
		Tags{Id: 17, Ru: "games", En: "игры"},
		Tags{Id: 18, Ru: "geometry", En: "геометрия"},
		Tags{Id: 19, Ru: "graph matchings", En: "паросочетания"},
		Tags{Id: 20, Ru: "graphs", En: "графы"},
		Tags{Id: 21, Ru: "greedy", En: "жадные алгоритмы"},
		Tags{Id: 22, Ru: "hashing", En: "хэши"},
		Tags{Id: 23, Ru: "implementation", En: "реализация"},
		Tags{Id: 24, Ru: "interactive", En: "интерактив"},
		Tags{Id: 25, Ru: "math", En: "математика"},
		Tags{Id: 26, Ru: "matrices", En: "матрицы"},
		Tags{Id: 27, Ru: "meet-in-the-middle", En: "meet-in-the-middle"},
		Tags{Id: 28, Ru: "number theory", En: "теория чисел"},
		Tags{Id: 29, Ru: "probabilities", En: "теория вероятностей"},
		Tags{Id: 30, Ru: "schedules", En: "расписания"},
		Tags{Id: 31, Ru: "shortest paths", En: "кратчайшие пути"},
		Tags{Id: 32, Ru: "sortings", En: "сортировки"},
		Tags{Id: 33, Ru: "string suffix structures", En: "строковые суфф. структуры"},
		Tags{Id: 34, Ru: "strings", En: "строки"},
		Tags{Id: 35, Ru: "ternary search", En: "тернарный поиск"},
		Tags{Id: 36, Ru: "trees", En: "деревья"},
		Tags{Id: 37, Ru: "two pointers", En: "два указателя"},
	}

	for _, item := range MyTags {
		tags.Tags = append(tags.Tags, models.TagJson{TagsId: item.Id, TagsRu: item.Ru, TagsEn: item.En})
	}

	result, err := json.Marshal(tags)
	if err != nil {
		return tools.CustomError(ctx, err, 1, "теги отвалились")
	}

	ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
	return ctx.JSONBlob(http.StatusOK, result)
}
