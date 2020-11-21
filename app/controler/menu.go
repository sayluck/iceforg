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
	menuGroup := r.Group("menu")
	{
		menuGroup.POST("", addMenu)
		menuGroup.GET("", list)
	}
}

func addMenu(c *gin.Context) {
	var (
		m      menu.MenuAddReq
		err    error
		menuID string
	)
	if err = c.ShouldBindJSON(&m); err != nil {
		resp(c, api.RespFailed(api.ParamsErr, err.Error()))
		return
	}
	m.SetContext(c)
	errs := ValidateStruct(c, &m)
	if len(errs) != 0 {
		resp(c, api.RespFailed(api.ParamsErr,
			multilingual.GetStrMsgs(errs)...))
		return
	}
	if menuID, err = menu.AddMenu(&m); err != nil {
		if strings.Contains(err.Error(), common.DuplicateEntry) {
			resp(c, api.RespFailed(api.OperationErr,
				multilingual.GetStrMsg(multilingual.MenuAlreadyExisted)))
			return
		}
		resp(c, api.RespFailed(api.SystemErr, err.Error()))
		return
	}

	resp(c, api.RespSucc(menuID))
}

func list(c *gin.Context) {
	var (
		err     error
		pageNum string
	)
	if pageNum = c.Query("pageNum"); pageNum == "" {
		resp(c, api.RespFailed(api.ParamsErr, err.Error()))
		return
	}

	var data interface{}
	if data, err = menu.List(pageNum); err != nil {
		resp(c, api.RespFailed(api.SystemErr, err.Error()))
		return
	}
	resp(c, api.RespSucc(data))
}
