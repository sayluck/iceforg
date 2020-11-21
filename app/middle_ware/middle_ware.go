package middle_ware

import (
	"bytes"
	"fmt"
	"iceforg/app/common"
	. "iceforg/app/log"
	"iceforg/app/service/user"
	"iceforg/pkg/common/api"
	"iceforg/pkg/multilingual"
	"iceforg/pkg/utils"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	lenUUUD = 15
)

func Trace() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		reqID := utils.GenerateUUID(lenUUUD)
		ctx.Set(common.ReqID, reqID)

		IceLog.Logger.Debug(setStartBaseLog(ctx))

		ctx.Next()

		IceLog.Logger.Debug(setEndBaseLog(ctx))
	}
}

func RecordPanic() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.AbortWithStatus(http.StatusInternalServerError)
				resp := api.RespFailed(api.SystemPanic,
					multilingual.GetStrMsg(multilingual.SystemPanicError))
				resp.ReqID = c.GetString(common.ReqID)
				c.JSON(http.StatusInternalServerError, resp)
				IceLog.Errorf(c, "recover form panic:%v", err)
			}
		}()
		c.Next()
	}
}

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader(common.Authorization)
		u, err := user.ParseToken(token)
		if err != nil {
			resp := api.RespFailed(api.UserInvalidToken,
				multilingual.GetStrMsg(err))
			resp.ReqID = ctx.GetString(common.ReqID)
			ctx.JSON(http.StatusOK, resp)
			ctx.Abort()
			return
		}
		ctx.Set(common.UserName, u.UserName)
		ctx.Set(common.UserID, u.UserID)
		// TODO add team
		// ctx.Set(common.TEAMID, u.Password)

		ctx.Next()
	}
}

func setStartBaseLog(c *gin.Context) string {
	if c.Request.Method == common.MethodGet {
		return fmt.Sprintf("reqID:%s,URL:%s",
			c.GetString(common.ReqID),
			c.Request.RequestURI)
	}
	bodyStr := ""
	body, err := ioutil.ReadAll(c.Request.Body)
	if err == nil {
		bodyStr = string(body)
	}
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	return fmt.Sprintf("reqID:%s,URL:%s,Body:%+v",
		c.GetString(common.ReqID),
		c.Request.RequestURI,
		bodyStr)
}

func setEndBaseLog(c *gin.Context) string {
	return fmt.Sprintf("reqID:%s,Status:%d",
		c.GetString(common.ReqID),
		c.Writer.Status())
}
