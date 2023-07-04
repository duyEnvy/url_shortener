package common

type successRes struct {
	Data interface{} `json:"data"`
}

func NewSuccessResponse(data interface{}) *successRes {
	return &successRes{Data: data}
}

func SimpleSuccessResponse(data interface{}) *successRes {
	return NewSuccessResponse(data)
}
