package common

type BaseReq struct {
	Code     string `json:"code"`
	Creator  string `json:"creator"`
	Modifier string `json:"modifier"`
}
