package models

import (
	"github.com/jackc/pgx/pgtype"
	"time"
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

type TaskCG struct {
	Description    string `db:"description"`
	MasterSolution string `db:"master_solution"`
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

type LikeTasks struct {
	CountTasks  int        `json:"count_tasks"`
	TasksIdList []int      `json:"tasks_id_list"`
	Tasks       []TaskJSON `json:"tasks"`
}

type SourceCode1 struct {
	Makefile string `json:"Makefile"`
	Main     string `json:"main.c"`
	Main1    string `json:"lib/sum.c"`
	Main2    string `json:"lib/sum.h"`
}

type SS123 struct {
	SourceCode   SourceCode1 `json:"sourceCode"`
	Tests        [][]string  `json:"tests"`
	BuildTimeout int         `json:"buildTimeout"`
	TestTimeout  int         `json:"testTimeout"`
}

type SourceCode struct {
	Makefile string `json:"Makefile"`
	Main     string `json:"main.cpp"`
}

type CheckSolution struct {
	SourceCode   SourceCode `json:"sourceCode"`
	Tests        [][]string `json:"tests"`
	BuildTimeout int        `json:"buildTimeout"`
	TestTimeout  float64    `json:"testTimeout"`
}

type CustomError struct {
	Number  int    `json:"number"`
	Comment string `json:"comment"`
	Error   string `json:"error"`
}

type CheckSolutionRequest struct {
	TaskId   int    `json:"task_id"`
	Solution string `json:"solution"`
}

type CheckSolutionUseCase struct {
	TaskId   int    `json:"task_id"`
	Solution string `json:"solution"`
}

type CheckSolutionUseCaseResponse struct {
	CheckTime    float64 `json:"checkTime"`
	BuildTime    float64 `json:"buildTime"`
	CheckResult  int     `json:"checkResult"`
	CheckMessage string  `json:"checkMessage"`
	TestsPassed  int     `json:"testsPassed"`
	TestsTotal   int     `json:"testsTotal"`
	LintSuccess  bool    `json:"lintSuccess"`
}

type SendTaskId uint64

type SendTask struct {
	ID           SendTaskId       `db:"id"`
	UserId       int              `db:"user_id"`
	TaskId       int              `db:"task_id"`
	CheckTime    float64          `db:"check_time"`
	BuildTime    float64          `db:"build_time"`
	CheckResult  int              `db:"check_result"`
	CheckMessage string           `db:"check_message"`
	TestsPassed  int              `db:"tests_passed"`
	TestsTotal   int              `db:"tests_total"`
	LintSuccess  bool             `db:"lint_success"`
	CodeText     string           `db:"code_text"`
	Date         pgtype.Timestamp `db:"date"`
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

type SimilarRequest struct {
	SourceCode string `json:"source_code"`
	ProblemUrl string `json:"problem_url"`
	Rating     int    `json:"rating"`
	Difficulty int    `json:"difficulty"`
	NRecs      int    `json:"n_recs"`
}

type SimilarResponse struct{}

type MlTaskResponse struct {
	ProblemUrl string  `json:"problem_url"`
	Rating     float64 `json:"rating"`
	Tags       []int   `json:"tags"`
}

type MlResponse struct {
	Tasks []MlTaskResponse
}

type LikeDb struct {
	Id     int    `db:"id"`
	UserId UserId `db:"user_id"`
	TaskId int    `db:"task_id"`
}

type LikesDb struct {
	Likes []LikeDb
}

type LikeJson struct {
	UserId UserId `json:"user_id"`
	TaskId int    `json:"task_id"`
}

type Message struct {
	Message string `json:"message"`
}

type Day struct {
	Day   time.Time `json:"day"`
	Count int       `json:"count_task"`
}

type Days struct {
	Days []Day `json:"days"`
}

type ShockMode struct {
	Today     bool `json:"today"`
	ShockMode int  `json:"chock_mode"`
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

type StoryItem struct {
	ProblemUrl      string `json:"problem_url"`
	Rating          int    `json:"rating"`
	Tags            []int  `json:"tags"`
	DifficultyMatch int    `json:"difficulty_match"`
	Solved          bool   `json:"solved"`
	NAttempts       int    `json:"n_attempts"`
}
type Story struct {
	UserId int         `json:"username"`
	Story  []StoryItem `json:"story"`
}

type StoryItem1 struct {
	ProblemUrl string `json:"problem_url"`
	Rating     int    `json:"rating"`
	Tags       []int  `json:"tags"`
	NAttempts  int    `json:"n_attempts"`
}
type Story1 struct {
	Solved  []StoryItem1 `json:"solved"`
	TooEasy []StoryItem1 `json:"too_easy"`
	TooHard []StoryItem1 `json:"too_hard"`
}

type Problems struct {
	ProblemUrl string `json:"problem_url"`
	Rating     int    `json:"rating"`
	Tags       []int  `json:"tags"`
}

type Recommended struct {
	RecommendedTag int        `json:"recommended_tag"`
	Priority       int        `json:"priority"`
	Problems       []Problems `json:"problems"`
}

type Rec struct {
	Rec []Recommended `json:"rec"`
}

type RecommendedResponse struct {
	RecommendedTag string     `json:"recommended_tag"`
	Priority       int        `json:"priority"`
	Problems       []TaskJSON `json:"problems"`
}

type RecResponse struct {
	Rec []RecommendedResponse `json:"rec"`
}

type Progress struct {
	Tag  int  `json:"tag"`
	Done bool `json:"done"`
}

type ColdStartML struct {
	Finished    bool       `json:"finished"`
	ProblemUrl  string     `json:"problem_url"`
	Tag         int        `json:"tag"`
	Progress    []Progress `json:"progress"`
	ProblemTags []int      `json:"problem_tags"`
	Rating      int        `json:"rating"`
}

type ColdStartResponse struct {
	Finished bool       `json:"finished"`
	Progress []Progress `json:"progress"`
	Task     TaskJSON   `json:"task"`
}

type ChatGPT struct {
	TaskId  int    `json:"task_id"`
	Message string `json:"message"`
	Code    string `json:"code"`
}

type ChatGPTRequest struct {
	UserMessage    string `json:"user_message"`
	Statement      string `json:"statement"`
	UserSolution   string `json:"user_solution"`
	MasterSolution string `json:"master_solution"`
}
