package common

type SuccessRes struct {
	Result interface{} `json:"result"`
	Paging interface{} `json:"paging"`
	Filter interface{} `json:"filter"`
}

type FailureRes struct {
	Result  interface{} `json:"result"`
}

func NewSuccessResponseWithPaging(data, paging, filter interface{}) *SuccessRes {
	return &SuccessRes{Result: data, Paging: paging, Filter: filter}
}

func NewSuccessResponseNoPaging(data interface{}) *SuccessRes {
	return &SuccessRes{Result: data, Paging: nil, Filter: nil}
}

func NewFailureResponse(err *AppError) *FailureRes {
	return &FailureRes{Result: err}
}
