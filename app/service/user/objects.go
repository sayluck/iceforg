package user

type UserRegister struct {
	UserName string `json:"userName" validate:"required"`
	Password string `json:"password" validate:"required"`
	NickName string `json:"nickName"`
}

type UserLogin struct {
	ID       int64  `json:"id"`
	UserName string `json:"userName" validate:"required"`
	Password string `json:"password" validate:"required"`
}
