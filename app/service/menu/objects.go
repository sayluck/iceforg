package menu

import (
	"context"
	"iceforg/app/service/common"
)

// request
type MenuReq struct {
	common.BaseReq

	PageNum int    `json:"pageNum" validate:"required"`
	Name    string `json:"name" validate:"required"`
	Sort    int    `json:"sort"`
	Route   string `json:"route" validate:"required"`

	context.Context `json:"-"`
}

// response
type MenuDetial struct {
	PageNum int    `json:"pageNum"`
	Name    string `json:"name"`
	Sort    int    `json:"sort"`
	Route   string `json:"route"`
}
