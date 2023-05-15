package models

type CustomError struct {
	Number  int    `json:"number"`
	Comment string `json:"comment"`
	Error   string `json:"error"`
}
