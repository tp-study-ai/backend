package models

type UserId uint64

type UserJson struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

//type UserDB struct {
//	Username string `db:"username"`
//	Password string `db:"password"`
//}

type UserDB struct {
	Id        UserId `db:"id"`
	Username  string `db:"username"`
	Password  string `db:"password"`
	ColdStart bool   `json:"cold_start"`
}

type UpdateUsernameDb struct {
	Id          UserId `db:"id"`
	Username    string `db:"username"`
	NewUsername string `db:"username"`
}

type UpdatePasswordDb struct {
	Id          UserId `db:"id"`
	Username    string `db:"username"`
	OldPassword string `db:"password"`
	NewPassword string `db:"password"`
}

type ResponseUserJson struct {
	Id        UserId `json:"id"`
	Username  string `json:"username"`
	ColdStart bool   `json:"cold_start"`
}

//type UpdateUsernameJson struct {
//	Id          UserId `json:"id"`
//	Username    string `json:"username"`
//	NewUsername string `json:"new_username"`
//}

type UpdateJson struct {
	NewUsername string `json:"new_username"`
	NewPassword string `json:"new_password"`
}
