package user

// request
type UserRegister struct {
	UserName string `json:"userName" validate:"required"`
	Password string `json:"password" validate:"required"`
	NickName string `json:"nickName"`
}

type UserLogin struct {
	UserID   string `json:"UserID"`
	UserName string `json:"UserName" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// response
type UserDetail struct {
	UserID   string `json:"userid"`
	UserName string `json:"userName"`
	NickName string `json:"nickName"`
}
