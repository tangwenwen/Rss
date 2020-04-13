package usersType

import "time"

type LoginResp struct {
	Token    string `json:"token"`
	UserType int    `json:"user_type"`
}

type GetInfoResp struct {
	UserName string `json:"userName"`
	UserType int    `json:"user_type"`
}

type Users struct {
	ID          int       `xorm:"id"`
	Username    string    `xorm:"user_name"`
	Password    string    `xorm:"password"`
	Age         int    `xorm:"age"`
	Gender      string    `xorm:"gender"`
	Occupation  string    `xorm:"occupation"`
	CreatedTime time.Time `xorm:"created_time"`
	Status      int    `xorm:"status"`
	UserType    int    `xorm:"user_type"`
}

type TokenValue struct {
	UserName string `json:"userName"`
	UserType string `json:"user_type"`
}

type AddUserReq struct {
	UserName   string `json:"userName",xorm:"user_name"`
	Password   string `json:"passWord",xorm:"password"`
	Age        string `json:"age",xorm:"age"`
	Gender     string `json:"gender",xorm:"gender"`
	Occupation string `json:"occupation",xorm:"occupation"`
}
