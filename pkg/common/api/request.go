package api

type BasePageReq struct {
	PageNum  int `json:"page_num"  binding:"gt=0"`
	PageSize int `json:"page_size" binding:"gt=0"`
}
