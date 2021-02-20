package util

import (
	"fmt"
	"reflect"
)

func CallFuncs(fc interface{}, params ...interface{}) (result interface{}, err error) {
	f := reflect.ValueOf(fc)
	if f.Kind() != reflect.Func {
		err = fmt.Errorf("fc is not func")
		return
	}
	if len(params) != f.Type().NumIn() {
		err = fmt.Errorf("the number of params is not adapted")
		return
	}

	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	resp := f.Call(in)
	if len(resp) > 0 {
		result = reflect.ValueOf(resp[0].Interface()).Interface()
		return
	}
	result = nil
	return
}
