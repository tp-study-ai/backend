package models

type TaskResponse struct {
	Id 					int			`db:"id"`
	Name 				string		`db:"name"`
	Description 		string		`db:"description"`
	PublicTests 		[]string	`db:"public_tests"`
	PrivateTests 		[]string	`db:"private_tests"`
	GeneratedTests 		[]string	`db:"generated_tests"`
	Difficulty  		string		`db:"difficulty"`
	CfContestId 		string		`db:"cf_contest_id"`
	CfIndex   			string		`db:"cf_index"`
	CfPoints 			string		`db:"cf_points"`
	CfRating   			string		`db:"cf_rating"`
	CfTags            	string		`db:"cf_tags"`
	TimeLimit        	string		`db:"time_limit"`
	MemoryLimitBytes 	string		`db:"memory_limit_bytes"`
	Link   				string		`db:"link"`
	TaskRu 				string		`db:"task_ru"`
	Input  				string		`db:"input"`
	Output 				string		`db:"output"`
	Note 				string		`db:"note"`
}