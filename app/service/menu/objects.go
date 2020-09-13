package menu

// request
type MenuReq struct {
	PageNum int    `json:"pageNum";validate:"required"`
	Name    string `json:"name";validate:"required"`
	Sort    int    `json:"sort"`
	Route   string `json:"route";validate:"required"`
}

// response
type MenuDetial struct {
	PageNum int    `json:"pageNum"`
	Name    string `json:"name"`
	Sort    int    `json:"sort"`
	Route   string `json:"route"`
}
