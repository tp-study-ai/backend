package models

import (
	"github.com/jackc/pgx/pgtype"
)

type TagJson struct {
	TagsId int    `json:"tags_id"`
	TagsEn string `json:"tags_en"`
	TagsRu string `json:"tags_ru"`
}

type TagsJson struct {
	Tags []TagJson `json:"tags"`
}

type TaskDB struct {
	Id               int              `db:"id"`
	Name             string           `db:"name"`
	Description      string           `db:"description"`
	PublicTests      []string         `db:"public_tests"`
	PrivateTests     []string         `db:"private_tests"`
	GeneratedTests   []string         `db:"generated_tests"`
	Difficulty       int              `db:"difficulty"`
	CfContestId      int              `db:"cf_contest_id"`
	CfIndex          string           `db:"cf_index"`
	CfPoints         float64          `db:"cf_points"`
	CfRating         int              `db:"cf_rating"`
	CfTags           pgtype.Int4Array `db:"cf_tags"`
	TimeLimit        float64          `db:"time_limit"`
	MemoryLimitBytes int              `db:"memory_limit_bytes"`
	Link             string           `db:"link"`
	ShortLink        string           `db:"short_link"`
	NameRu           string           `db:"name_ru"`
	TaskRu           string           `db:"task_ru"`
	Input            string           `db:"input"`
	Output           string           `db:"output"`
	Note             string           `db:"note"`
	MasterSolution   string           `db:"master_solution"`
}

type TasksResponse struct {
	Tasks []TaskDB
}

type TaskJSON struct {
	Id               int      `json:"id"`
	Name             string   `json:"name"`
	Description      string   `json:"description"`
	PublicTests      []string `json:"public_tests"`
	Difficulty       int      `json:"difficulty"`
	CfContestId      int      `json:"cf_contest_id"`
	CfIndex          string   `json:"cf_index"`
	CfPoints         float64  `json:"cf_points"`
	CfRating         int      `json:"cf_rating"`
	CfTagsID         []int    `json:"cf_tags_ID"`
	CfTagsRu         []string `json:"cf_tags_RU"`
	CfTagsEN         []string `json:"cf_tags_en"`
	TimeLimit        float64  `json:"time_limit"`
	MemoryLimitBytes int      `json:"memory_limit_bytes"`
	Link             string   `json:"link"`
	ShortLink        string   `json:"short_link"`
	NameRu           string   `json:"name_ru"`
	TaskRu           string   `json:"task_ru"`
	Input            string   `json:"input"`
	Output           string   `json:"output"`
	Note             string   `json:"note"`
}

type Tasks struct {
	Tasks []TaskJSON `json:"tasks"`
}

type TasksPagination struct {
	TaskCount int        `json:"task_count"`
	Tasks     []TaskJSON `json:"tasks"`
}

type SendTasks struct {
	Tasks []SendTask
}

type SendTaskJson struct {
	ID           SendTaskId       `json:"id"`
	UserId       int              `json:"user_id"`
	TaskId       int              `json:"task_id"`
	CheckTime    float64          `json:"check_time"`
	BuildTime    float64          `json:"build_time"`
	CheckResult  int              `json:"check_result"`
	CheckMessage string           `json:"check_message"`
	TestsPassed  int              `json:"tests_passed"`
	TestsTotal   int              `json:"tests_total"`
	LintSuccess  bool             `json:"lint_success"`
	CodeText     string           `json:"code_text"`
	Date         pgtype.Timestamp `json:"date"`
}

type SendTasksJson struct {
	Tasks []SendTaskJson `json:"tasks"`
}

type DoneTask struct {
	CountDoneTask int        `json:"count_done_task"`
	DoneTask      []TaskJSON `json:"done_task"`
}

type DifficultyDb struct {
	UserId     int `db:"user_id"`
	TaskId     int `db:"task_id"`
	Difficulty int `db:"difficulty"`
}

type DifficultyJson struct {
	UserId     int `json:"user_id"`
	TaskId     int `json:"task_id"`
	Difficulty int `json:"difficulty"`
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

var TagDict = map[int][]string{
	1:  {"*special", "*особая задача"},
	2:  {"2-sat", "2-sat"},
	3:  {"binary search", "бинарный поиск"},
	4:  {"bitmasks", "битмаски"},
	5:  {"brute force", "перебор"},
	6:  {"chinese remainder theorem", "китайская теорема об остатках"},
	7:  {"combinatorics", "комбинаторика"},
	8:  {"constructive algorithms", "конструктив"},
	9:  {"data structures", "структуры данных"},
	10: {"dfs and similar", "поиск в глубину и подобное"},
	11: {"divide and conquer", "разделяй и властвуй"},
	12: {"dp", "дп"},
	13: {"dsu", "системы непересекающихся множеств"},
	14: {"expression parsing", "разбор выражений"},
	15: {"fft", "быстрое преобразование Фурье"},
	16: {"flows", "потоки"},
	17: {"games", "игры"},
	18: {"geometry", "геометрия"},
	19: {"graph matchings", "паросочетания"},
	20: {"graphs", "графы"},
	21: {"greedy", "жадные алгоритмы"},
	22: {"hashing", "хэши"},
	23: {"implementation", "реализация"},
	24: {"interactive", "интерактив"},
	25: {"math", "математика"},
	26: {"matrices", "матрицы"},
	27: {"meet-in-the-middle", "meet-in-the-middle"},
	28: {"number theory", "теория чисел"},
	29: {"probabilities", "теория вероятностей"},
	30: {"schedules", "расписания"},
	31: {"shortest paths", "кратчайшие пути"},
	32: {"sortings", "сортировки"},
	33: {"string suffix structures", "строковые суфф. структуры"},
	34: {"strings", "строки"},
	35: {"ternary search", "тернарный поиск"},
	36: {"trees", "деревья"},
	37: {"two pointers", "два указателя"},
}
