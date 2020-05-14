package middle_ware

import (
	"iceforg/app/common"
	. "iceforg/app/log"
	"iceforg/app/service/user"
	"iceforg/pkg/common/api"
	"iceforg/pkg/multilingual"
	"iceforg/pkg/utils"
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

		Log.Debug(SetStartBaseLog(ctx))

		ctx.Next()

		Log.Debug(SetEndBaseLog(ctx))
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
				Log.Errorf("recover form panic:%v", err)
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
			resp := api.RespFailed(api.OperationErr,
				multilingual.GetStrMsg(err))
			resp.ReqID = ctx.GetString(common.ReqID)
			ctx.JSON(http.StatusOK, resp)
			ctx.Abort()
			return
		}
		ctx.Set(common.UserName, u.UserName)
		ctx.Set(common.UserID, u.ID)

		ctx.Next()
	}
}
