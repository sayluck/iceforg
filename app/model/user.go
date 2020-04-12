package model

func (*User) TableName() string {
	return "t_user"
}

type User struct {
	UserName string `json:"userName" gorm:"column:name" binding:"required"`
	Password string `json:"password" binding:"required"`
	NickName string `json:"nickName" gorm:"column:nick_name"`
}
