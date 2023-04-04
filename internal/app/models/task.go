package models

type TaskResponse struct {
	Id               int      `db:"id"`
	Name             string   `db:"name"`
	Description      string   `db:"description"`
	PublicTests      []string `db:"public_tests"`
	PrivateTests     []string `db:"private_tests"`
	GeneratedTests   []string `db:"generated_tests"`
	Difficulty       string   `db:"difficulty"`
	CfContestId      int      `db:"cf_contest_id"`
	CfIndex          string   `db:"cf_index"`
	CfPoints         float64  `db:"cf_points"`
	CfRating         int      `db:"cf_rating"`
	CfTags           []string `db:"cf_tags"`
	TimeLimit        string   `db:"time_limit"`
	MemoryLimitBytes int      `db:"memory_limit_bytes"`
	Link             string   `db:"link"`
	TaskRu           string   `db:"task_ru"`
	Input            string   `db:"input"`
	Output           string   `db:"output"`
	Note             string   `db:"note"`
}

type TasksResponse struct {
	Tasks []TaskResponse
}

type Task struct {
	Id               int      `json:"id"`
	Name             string   `json:"name"`
	Description      string   `json:"description"`
	PublicTests      []string `json:"public_tests"`
	PrivateTests     []string `json:"private_tests"`
	GeneratedTests   []string `json:"generated_tests"`
	Difficulty       string   `json:"difficulty"`
	CfContestId      int      `json:"cf_contest_id"`
	CfIndex          string   `json:"cf_index"`
	CfPoints         float64  `json:"cf_points"`
	CfRating         int      `json:"cf_rating"`
	CfTags           []string `json:"cf_tags"`
	TimeLimit        string   `json:"time_limit"`
	MemoryLimitBytes int      `json:"memory_limit_bytes"`
	Link             string   `json:"link"`
	TaskRu           string   `json:"task_ru"`
	Input            string   `json:"input"`
	Output           string   `json:"output"`
	Note             string   `json:"note"`
}

type Tasks struct {
	Tasks []Task `json:"tasks"`
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
	TestTimeout  int        `json:"testTimeout"`
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
