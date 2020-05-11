package user

type UserRegister struct {
	UserName string `json:"userName"binding:"required"`
	Password string `json:"password",binding:"required"`
	NickName string `json:"nickName"`
}

type UserLogin struct {
	ID       int64  `json:"id"`
	UserName string `json:"userName",binding:"required"`
	Password string `json:"password",binding:"required"`
}
