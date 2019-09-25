package schema

// ---------------------------------------------------------------
type HTTPErrRes struct {
	Error     ErrorItem `json:"error" swaggo:"true,错误信息"`
	RequestId string    `json:"request_id" swaggo:"true,请求ID"`
}

type ErrorItem struct {
	Code    int    `json:"code" swaggo:"true,状态码"`
	Status  string `json:"status" swaggo:"true,状态"`
	Message string `json:"message" swaggo:"错误描述"`
}

type HTTPSucRes struct {
	Result    *HTTPList `json:"result" swaggo:"true,返回数据"`
	RequestId string    `json:"data" swaggo:"true,请求ID"`
}

// HTTPList HTTP响应列表数据
type HTTPList struct {
	Data       interface{}     `json:"data"`
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
