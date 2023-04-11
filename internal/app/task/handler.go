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
	if err != nil {
		return tools.CustomError(ctx, err, 2, "CheckSolution usecase")
	}

	result, err := json.Marshal(testisResponse)
	if err != nil {
		return tools.CustomError(ctx, err, 3, "CheckSolution Bind")
	}

	ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
	return ctx.JSONBlob(http.StatusOK, result)
}

type Tags struct {
	Id  int
	Eng string
	Rus string
}

var MyTags = []Tags{
	Tags{Id: 1, Eng: "*special", Rus: "*особая задача"},
	Tags{Id: 2, Eng: "2-sat", Rus: "2-sat"},
	Tags{Id: 3, Eng: "binary search", Rus: "бинарный поиск"},
	Tags{Id: 4, Eng: "bitmasks", Rus: "битмаски"},
	Tags{Id: 5, Eng: "brute force", Rus: "перебор"},
	Tags{Id: 6, Eng: "chinese remainder theorem", Rus: "китайская теорема об остатках"},
	Tags{Id: 7, Eng: "combinatorics", Rus: "комбинаторика"},
	Tags{Id: 8, Eng: "constructive algorithms", Rus: "конструктив"},
	Tags{Id: 9, Eng: "data structures", Rus: "структуры данных"},
	Tags{Id: 10, Eng: "dfs and similar", Rus: "поиск в глубину и подобное"},
	Tags{Id: 11, Eng: "divide and conquer", Rus: "разделяй и властвуй"},
	Tags{Id: 12, Eng: "dp", Rus: "дп"},
	Tags{Id: 13, Eng: "dsu", Rus: "системы непересекающихся множеств"},
	Tags{Id: 14, Eng: "expression parsing", Rus: "разбор выражений"},
	Tags{Id: 15, Eng: "fft", Rus: "быстрое преобразование Фурье"},
	Tags{Id: 16, Eng: "flows", Rus: "потоки"},
	Tags{Id: 17, Eng: "games", Rus: "игры"},
	Tags{Id: 18, Eng: "geometry", Rus: "геометрия"},
	Tags{Id: 19, Eng: "graph matchings", Rus: "паросочетания"},
	Tags{Id: 20, Eng: "graphs", Rus: "графы"},
	Tags{Id: 21, Eng: "greedy", Rus: "жадные алгоритмы"},
	Tags{Id: 22, Eng: "hashing", Rus: "хэши"},
	Tags{Id: 23, Eng: "implementation", Rus: "реализация"},
	Tags{Id: 24, Eng: "interactive", Rus: "интерактив"},
	Tags{Id: 25, Eng: "math", Rus: "математика"},
	Tags{Id: 26, Eng: "matrices", Rus: "матрицы"},
	Tags{Id: 27, Eng: "meet-in-the-middle", Rus: "meet-in-the-middle"},
	Tags{Id: 28, Eng: "number theory", Rus: "теория чисел"},
	Tags{Id: 29, Eng: "probabilities", Rus: "теория вероятностей"},
	Tags{Id: 30, Eng: "schedules", Rus: "расписания"},
	Tags{Id: 31, Eng: "shortest paths", Rus: "кратчайшие пути"},
	Tags{Id: 32, Eng: "sortings", Rus: "сортировки"},
	Tags{Id: 33, Eng: "string suffix structures", Rus: "строковые суфф. структуры"},
	Tags{Id: 34, Eng: "strings", Rus: "строки"},
	Tags{Id: 35, Eng: "ternary search", Rus: "тернарный поиск"},
	Tags{Id: 36, Eng: "trees", Rus: "деревья"},
	Tags{Id: 37, Eng: "two pointers", Rus: "два указателя"},
}

func (h HandlerTask) GetTags(ctx echo.Context) error {
	tags := models.TagsJson{}

	for _, item := range MyTags {
		tags.Tags = append(tags.Tags, models.TagJson{TagsId: item.Id, TagsEn: item.Eng, TagsRu: item.Rus})
	}

	result, err := json.Marshal(tags)
	if err != nil {
		return tools.CustomError(ctx, err, 1, "теги отвалились")
	}

	ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
	return ctx.JSONBlob(http.StatusOK, result)
}
