package controler

import (
	"fmt"
	"iceforg/pkg/config"

	"github.com/gin-gonic/gin"
)

var (
	defaultPort = "8080"
	apiVersion  = "/v1"
)

type Route struct {
	*config.App
}

func (r *Route) Router() {
	router := gin.Default()
	routerGrp := router.Group(apiVersion)

	// middleware
	userRouter(routerGrp)
	if r.App != nil && r.Port != "" {
		defaultPort = r.Port
	}
	router.Run(fmt.Sprintf(":%s", defaultPort))
}
