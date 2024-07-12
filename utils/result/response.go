package result

type HttpResult struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (r *HttpResult) Success(data interface{}) {
	r.Code = HTTP_STATUS_SUCCESS
	r.Msg = "success"
	r.Data = data
}

func (r *HttpResult) Error(msg string) {
	r.Code = HTTP_STATUS_ERROR
	r.Msg = msg
	r.Data = nil
}
