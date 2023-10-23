package models

type LoginReq struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginRes struct {
	Token    string `json:"token"`
}

type User struct {
	Id       int64
	UserName string
	Password string
}

type UpdateUser struct {
	PhoneNumber string
	Password    string
}


