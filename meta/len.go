package meta

import (
	"reflect"
)

func Len(v interface{}) int {
	if v == nil {
		return 0
	}
	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}
	return rv.Len()
}
