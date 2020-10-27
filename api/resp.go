package api

type CommPageResp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

type PageData struct {
	TotalCount int         `json:"total_count"`
	Items      interface{} `json:"items"`
}

func ApiReturn(args ...interface{}) CommPageResp {
	var (
		code = 0
		msg  = "OK"
		data interface{}
	)
	if len(args) >= 2 && args[0] != 0 {
		msg = args[1].(string)
	}
	if len(args) >= 3 {
		data = args[2]
	}
	return CommPageResp{code, msg, data}
}

func EchoReturn(code int, msg string, data interface{}) *CommPageResp {
	var ar = &CommPageResp{}
	ar.Code = code
	ar.Msg = msg
	ar.Data = data
	return ar
}
