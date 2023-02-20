package model

type User struct {
	ID       int64  `json:"id"`
	UserName string `json:"username"`
	Pws      string `json:"pws"`
}

type UserCreate struct {
	UserCreate User `json:"user"` 
}
