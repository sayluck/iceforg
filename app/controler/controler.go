package controler

import (
	"fmt"
	"iceforg/app/common"
	"iceforg/app/middle_ware"
	"iceforg/pkg/common/api"
	"iceforg/pkg/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	defaultPort = "8080"
	apiVersion  = "/api/v1"
)

type Route struct {
	*config.App
}

func (r *Route) Router() {
	router := gin.New()

	// middle ware
	router.Use(middle_ware.Trace())
	router.Use(middle_ware.RecordPanic())
	routerGrp := router.Group(apiVersion)

	// user register & login
	userLoginRouter(routerGrp)

	routerGrp.Use(middle_ware.Auth())
	// menu
	menuRouterGroup(routerGrp)
	// user
	userRouter(routerGrp)

	if r.App != nil && r.Port != "" {
		defaultPort = r.Port
	}
	router.Run(fmt.Sprintf(":%s", defaultPort))
}

func resp(c *gin.Context, obj *api.Resp) {
	obj.ReqID = c.GetString(common.ReqID)
	c.JSON(http.StatusOK, obj)
}
