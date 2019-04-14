package meta

import (
	"reflect"

	"github.com/gobuffalo/helpers/hctx"
)

func New() hctx.Map {
	return hctx.Map{
		"len": Len,
	}
}

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
