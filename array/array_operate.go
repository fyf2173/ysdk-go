package array

import (
	"fmt"
	"reflect"
)

func GetColumns(haystack interface{}, field string) interface{} {
	oldArr := reflect.ValueOf(haystack)
	if oldArr.Kind() != reflect.Slice {
		panic("haystack not a slice")
	}
	isPtr := false
	if oldArr.Type().Elem().Kind() == reflect.Ptr {
		isPtr = true
	}
	var rslice reflect.Value
	for i := 0; i <= oldArr.Len()-1; i++ {
		var val reflect.Value
		if isPtr {
			val = oldArr.Index(i).Elem().FieldByName(field)
		} else {
			val = oldArr.Index(i).FieldByName(field)
		}
		if !rslice.IsValid() {
			rslice = reflect.MakeSlice(reflect.SliceOf(val.Type()), 0, oldArr.Len())
		}
		rslice = reflect.Append(rslice, val)
	}
	return rslice.Interface()
}

// ArrayPluck 按长度分组
func ArrayPluck(src []interface{}, length int) (dst [][]interface{}) {
	step := 0
	for {
		nextStep := step + length
		if nextStep >= len(src) {
			dst = append(dst, src[step:])
			break
		}
		dst = append(dst, src[step:nextStep])
		step = nextStep
	}
	return dst
}

// ArrayPluckByT 按长度分组
func ArrayPluckByT[T any](src []T, length int) (dst [][]T) {
	step := 0
	for {
		nextStep := step + length
		if nextStep >= len(src) {
			dst = append(dst, src[step:])
			break
		}
		dst = append(dst, src[step:nextStep])
		step = nextStep
	}
	return dst
}

func ArrayUniq[T any](src []T) (dst []T) {
	var tmpMap = make(map[string]T)
	for _, v := range src {
		if _, ok := tmpMap[fmt.Sprintf("%v", v)]; ok {
			continue
		}
		tmpMap[fmt.Sprintf("%v", v)] = v
	}

	for _, v := range tmpMap {
		dst = append(dst, v)
	}
	return dst
}
