package models

type TaskDbForChatGPT struct {
	Description    string `db:"description"`
	MasterSolution string `db:"master_solution"`
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
