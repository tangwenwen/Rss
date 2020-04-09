package usersType



type LoginResp struct {
	Token    string `json:"token"`
	UserType int    `json:"user_type"`
}

type GetInfoResp struct {
	UserName string `json:"userName"`
	UserType int `json:"user_type"`
}


type Users struct{
	ID int `xorm:"id"`
	Username string `xorm:"user_name"`
	Password string `xorm:"password"`
	Age string `xorm:"age"`
	Gender string `xorm:"gender"`
	Occupation string `xorm:"occupation"`
	CreatedTime string `xorm:"created_time"`
	Status string `xorm:"status"`
	UserType string `xorm:"user_type"`
}

type TokenValue struct {
	UserName string `json:"userName"`
	UserType string `json:"user_type"`
}
