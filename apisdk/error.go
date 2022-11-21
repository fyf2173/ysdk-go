package apisdk

import (
	"errors"
	"fmt"
)

type XError struct {
	Code int   `json:"code"`
	Err  error `json:"err,omitempty"`
}

func NewXError(code int, err error) *XError {
	return &XError{Code: code, Err: err}
}

func NewXErrorMsg(code int, msg string) *XError {
	return &XError{Code: code, Err: errors.New(msg)}
}

func (ae *XError) Error() string {
	if ae.Err != nil {
		return ae.Err.Error()
	}
	return fmt.Sprintf("未定义错误码[%d]", ae.Code)
}

func (ae *XError) String() string {
	if ae.Err != nil {
		return fmt.Sprintf("[%d]%s", ae.Code, ae.Err.Error())
	}
	return fmt.Sprintf("未定义错误码[%d]", ae.Code)
}
