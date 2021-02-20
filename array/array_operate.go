package array

import (
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
