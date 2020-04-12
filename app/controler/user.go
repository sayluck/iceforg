package controler

import (
	"iceforg/app/model"
	"iceforg/app/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func userRouter(r *gin.RouterGroup) {
	user := r.Group("user")
	{
		// user register
		user.POST("register", register)
		//user.POST("uploadHeaderImg", v1.UploadHeaderImg)   //上传头像
		//user.POST("getUserList", v1.GetUserList)           // 分页获取用户列表
		//user.POST("setUserAuthority", v1.SetUserAuthority) //设置用户权限
	}
}

func register(c *gin.Context) {
	var (
		user model.User
		err  error
	)
	if err = c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusOK, "params error")
		return
	}

	if err = service.Register(&user); err != nil {
		c.JSON(http.StatusOK, "inter error:"+err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
}
