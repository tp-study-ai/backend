package task

import "github.com/tp-study-ai/backend/internal/app/models"

type UcaseTask struct {
	Repo Repository
}

func NewUcaseTask (TaskRepo Repository) *UcaseTask {
	return &UcaseTask{
		Repo: TaskRepo,
	}
}

func (u *UcaseTask) GetTask() (Task models.TaskResponse, err error) {
	Task, err = u.Repo.GetTask()

	if err != nil {
		return
	}
	return
}

func (u *UcaseTask) GetTaskById(id int) (Task models.TaskResponse, err error) {
	Task, err = u.Repo.GetTaskById(id)

	if err != nil {
		return
	}
	return
}