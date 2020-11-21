package menu

import (
	"context"
	"iceforg/app/service/common"
)

// request
type MenuAddReq struct {
	common.BaseReq

	PageNum int    `json:"pageNum" validate:"required"`
	Name    string `json:"name" validate:"required"`
	Sort    int    `json:"sort"`
	Level   int    `json:"level" validate:"required,min=1"`
	SupCode string `json:"subCode"`
	Route   string `json:"route" validate:"required"`

	context.Context `json:"-"`
}

// response
type MenuDetial struct {
	PageNum int    `json:"pageNum"`
	Name    string `json:"name"`
	Sort    int    `json:"sort"`
	Level   int    `json:"level"`
	Route   string `json:"route"`
}

type MenuTree struct {
	Code    string        `json:"code"`
	PageNum int           `json:"pageNum"`
	Name    string        `json:"name"`
	Sort    int           `json:"sort"`
	Level   int           `json:"level"`
	Route   string        `json:"route"`
	SupCode string        `json:"supCode"`
	SubMenu []interface{} `json:"sub_menu"`
}
