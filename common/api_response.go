package common

type SuccessRes struct {
	Result interface{} `json:"result"`
	Paging interface{} `json:"paging"`
	Filter interface{} `json:"filter"`
}

type FailureRes struct {
	Result  interface{} `json:"result"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
}

func NewSuccessResponseWithPaging(data, paging, filter interface{}) *SuccessRes {
	return &SuccessRes{Result: data, Paging: paging, Filter: filter}
}

func NewSuccessResponseNoPaging(data interface{}) *SuccessRes {
	return &SuccessRes{Result: data, Paging: nil, Filter: nil}
}

func NewFailureResponse(code int, message string) *FailureRes {
	return &FailureRes{Result: nil, Code: code, Message: message}
}
