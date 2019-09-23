package schema

// HTTPError HTTP响应错误
type HTTPError struct {
	Error     HTTPErrorItem `json:"error" swaggo:"true,错误项"`
	RequestId string        `json:"request_id" swaggo:"true,请求ID"`
}

// HTTPErrorItem HTTP响应错误项
type HTTPErrorItem struct {
	Code    int    `json:"code" swaggo:"true,错误码"`
	Message string `json:"message" swaggo:"true,错误信息"`
}

// HTTPStatus HTTP响应状态
type HTTPStatus struct {
	Status    string `json:"status" swaggo:"true,状态(OK)"`
	RequestId string `json:"request_id" swaggo:"true,请求ID"`
}

type HTTPErrRes struct {
	Error     Error  `json:"error" swaggo:"true,错误信息"`
	RequestId string `json:"request_id" swaggo:"true,请求ID"`
}

type Error struct {
	Code    int    `json:"code" swaggo:"true,状态码"`
	Status  int    `json:"status" swaggo:"true,状态"`
	Message string `json:"message" swaggo:"错误描述"`
}

type HTTPSucRes struct {
	Data      *HTTPList `json:"result" swaggo:"true,返回数据"`
	RequestId string    `json:"data" swaggo:"true,请求ID"`
}

// HTTPList HTTP响应列表数据
type HTTPList struct {
	Result     interface{}     `json:"data"`
	Pagination *HTTPPagination `json:"pagination,omitempty"`
}

// HTTPPagination HTTP分页数据
type HTTPPagination struct {
	Total    int `json:"total"`
	Current  int `json:"current"`
	PageSize int `json:"pageSize"`
}

// PaginationParam 分页查询条件
type PaginationParam struct {
	PageIndex int // 页索引
	PageSize  int // 页大小
}

// PaginationResult 分页查询结果
type PaginationResult struct {
	Total int // 总数据条数
}
