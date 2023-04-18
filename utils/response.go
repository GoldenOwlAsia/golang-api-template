package utils

const (
	ResponseStatusCodeSuccess = "success"
	ResponseStatusCodeFailed  = "error"
)

type ResponseSuccess struct {
	Code       int         `json:"code"`
	StatusCode string      `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

type RespSuccess struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseFailed struct {
	Code       int    `json:"code"`
	StatusCode string `json:"status_code"`
	Message    string `json:"message"`
}

type RespFailed struct {
	Message string `json:"message"`
}

type StatusResp string

const (
	StatusError   StatusResp = "error"
	StatusSuccess StatusResp = "success"
)

type Response struct {
	Status  StatusResp  `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func GetRespError(msg string, data any) *Response {
	resp := &Response{
		Status:  StatusError,
		Message: msg,
		Data:    map[string]interface{}{},
	}
	if data != nil {
		resp.Data = data
	}
	return resp
}

func GetRespSuccess(msg string, data any) *Response {
	return &Response{
		Status:  StatusSuccess,
		Message: msg,
		Data:    data,
	}
}
