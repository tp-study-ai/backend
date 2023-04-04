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
