package controler

import (
	"iceforg/app/common"
	. "iceforg/app/log"
	"iceforg/app/service/user"
	. "iceforg/app/validate"
	"iceforg/modules/userCenter"
	"iceforg/pkg/common/api"
	"iceforg/pkg/multilingual"
	"strings"

	"github.com/gin-gonic/gin"
)

func userLoginRouter(r *gin.RouterGroup) {
	user := r.Group("user")
	{
		// user register & login
		user.POST("register", register)
		user.POST("login", login)
	}
}

func userRouter(r *gin.RouterGroup) {
	user := r.Group("user")
	{
		// user
		user.GET("/", detail)
		user.GET("/current", currentUser)
	}
}

func register(c *gin.Context) {
	var (
		u      userCenter.UserRegister
		err    error
		userID string
	)
	if err = c.ShouldBindJSON(&u); err != nil {
		resp(c, api.RespFailed(api.ParamsErr, err.Error()))
		return
	}
	errs := ValidateStruct(c, &u)
	if len(errs) != 0 {
		resp(c, api.RespFailed(api.ParamsErr,
			multilingual.GetStrMsgs(errs)...))
		return
	}

	if userID, err = user.Register(&u); err != nil {
		if strings.Contains(err.Error(), common.DuplicateEntry) {
			resp(c, api.RespFailed(api.OperationErr,
				multilingual.GetStrMsg(multilingual.UserAlreadyExisted)))
			return
		}
		resp(c, api.RespFailed(api.SystemErr, err.Error()))
		return
	}

	resp(c, api.RespSucc(userID))
}

func detail(c *gin.Context) {
	var (
		name string
		err  error
		u    interface{}
	)
	name = c.Query("name")
	if u, err = user.Detail(name); err != nil {
		resp(c, api.RespFailed(api.OperationErr, multilingual.GetStrMsg(err)))
		return
	}
	resp(c, api.RespSucc(u))
}

func currentUser(c *gin.Context) {
	var (
		err  error
		name string
		u    *userCenter.UserDetail
	)
	name = c.GetString(user.UserName)
	if u, err = user.Detail(name); err != nil {
		resp(c, api.RespFailed(api.OperationErr, multilingual.GetStrMsg(err)))
		return
	}
	resp(c, api.RespSucc(u))
}

func login(c *gin.Context) {
	var (
		u     userCenter.UserLogin
		err   error
		token string
	)

	if err := c.ShouldBindJSON(&u); err != nil {
		IceLog.Errorf(c, "bind user error:%v", err)
		resp(c, api.RespFailed(api.ParamsErr, err.Error()))
		return
	}

	token, err = user.Login(&u)
	if err != nil {
		IceLog.Errorf(c, "login error:%s", err.Error())
		resp(c, api.RespFailed(api.SystemErr, multilingual.GetStrMsg(err)))
		return
	}
	resp(c, api.RespSucc(token))
}
