package entity

type User struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Mail     string `json:"mail"`
	Password string `json:"password"`
}
