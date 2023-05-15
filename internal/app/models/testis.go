package models

import "github.com/jackc/pgx/pgtype"

type TaskDBForTestis struct {
	PrivateTests []string `db:"private_tests"`
	TimeLimit    float64  `db:"time_limit"`
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

type CheckSolutionRequest struct {
	TaskId   int    `json:"task_id"`
	Solution string `json:"solution"`
}
