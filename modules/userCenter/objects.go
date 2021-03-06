package userCenter

// request
type UserRegister struct {
	UserName string `json:"userName" validate:"required"`
	Password string `json:"password" validate:"required"`
	NickName string `json:"nickName"`
}

type UserLogin struct {
	UserName string `json:"userName" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// response
type UserDetail struct {
	Code     string `json:"code"`
	UserName string `json:"userName"`
	NickName string `json:"nickName"`
}

// model
type User struct {
	Code     string `gorm:"column:m_code" json:"code"`
	UserName string `gorm:"column:name"`
	Password string `gorm:"column:password"`
	NickName string `gorm:"column:nick_name"`
}
