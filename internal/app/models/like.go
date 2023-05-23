package models

type LikeJson struct {
	UserId UserId `json:"user_id"`
	TaskId int    `json:"task_id"`
}

type LikeDb struct {
	Id     int    `db:"id"`
	UserId UserId `db:"user_id"`
	TaskId int    `db:"task_id"`
}

type LikesDb struct {
	Likes []LikeDb
}

type LikeTasks struct {
	CountTasks  int        `json:"count_tasks"`
	TasksIdList []int      `json:"tasks_id_list"`
	Tasks       []TaskJSON `json:"tasks"`
}
