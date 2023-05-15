package models

type TaskDbForChatGPT struct {
	Description    string `db:"description"`
	MasterSolution string `db:"master_solution"`
}
