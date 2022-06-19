package api

import (
	"errors"
	"fmt"
)

type ApiError struct {
	Code int   `json:"code"`
	Err  error `json:"err,omitempty"`
}

func NewError(code int, err error) *ApiError {
	return &ApiError{Code: code, Err: err}
}

func NewErrorMsg(code int, msg string) *ApiError {
	return &ApiError{Code: code, Err: errors.New(msg)}
}

func (ae *ApiError) Error() string {
	if ae.Err != nil {
		return ae.Err.Error()
	}
	return fmt.Sprintf("未定义错误码[%d]", ae.Code)
}

func (ae *ApiError) String() string {
	if ae.Err != nil {
		return fmt.Sprintf("[%d]%s", ae.Code, ae.Err.Error())
	}
	return fmt.Sprintf("未定义错误码[%d]", ae.Code)
}
