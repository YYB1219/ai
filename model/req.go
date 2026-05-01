package model

// AskRequest 定义前端传来的请求参数
type AskRequest struct {
	Question string `json:"question"`
}

// AskResponse 定义接口返回结果
type AskResponse struct {
	Question string `json:"question"`
	SQL      string `json:"sql"`
}
