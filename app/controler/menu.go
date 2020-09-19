package controler

import (
	"iceforg/app/common"
	"iceforg/app/service/menu"
	. "iceforg/app/validate"
	"iceforg/pkg/common/api"
	"iceforg/pkg/multilingual"

	"strings"

	"github.com/gin-gonic/gin"
)

func menuRouterGroup(r *gin.RouterGroup) {
	menu := r.Group("menu")
	{
		menu.POST("/", addMenu)
	}
}

func addMenu(c *gin.Context) {
	var (
		m      menu.MenuReq
		err    error
		userID string
	)
	if err = c.ShouldBindJSON(&m); err != nil {
		resp(c, api.RespFailed(api.ParamsErr, err.Error()))
		return
	}
	errs := ValidateStruct(c, &m)
	if len(errs) != 0 {
		resp(c, api.RespFailed(api.ParamsErr,
			multilingual.GetStrMsgs(errs)...))
		return
	}

	if userID, err = menu.AddMenu(&m); err != nil {
		if strings.Contains(err.Error(), common.DuplicateEntry) {
			resp(c, api.RespFailed(api.OperationErr,
				multilingual.GetStrMsg(multilingual.MenuAlreadyExisted)))
			return
		}
		resp(c, api.RespFailed(api.SystemErr, err.Error()))
		return
	}

	resp(c, api.RespSucc(userID))
}
