package models

type UserId uint64

type UserJson struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserDB struct {
	Username string `db:"username"`
	Password string `db:"password"`
}

type User struct {
	Id       int    `db:"id"`
	Username string `db:"username"`
	Password string `db:"password"`
}

type ResponseUserDb struct {
	Id       int    `db:"id"`
	Username string `db:"username"`
}

type ResponseUserJson struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}
