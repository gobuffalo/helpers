package paths

import (
	"errors"
	"fmt"
	"path"
	"reflect"
	"strings"

	"github.com/gobuffalo/flect/name"
)

type Pathable interface {
	ToPath() string
}

type Paramable interface {
	ToParam() string
}

// PathFor takes an `interface{}`, or a `slice` of them,
// and tries to convert it to a `/foos/{id}` style URL path.
// Rules:
// * if `string` it is returned as is
// * if `Pathable` the `ToPath` method is returned
// * if `slice` or an `array` each element is run through the helper then joined
// * if `struct` the name of the struct, pluralized is used for the name
// * if `Paramable` the `ToParam` method is used to fill the `{id}` slot
// * if `struct.Slug` the slug is used to fill the `{id}` slot of the URL
// * if `struct.ID` the ID is used to fill the `{id}` slot of the URL
func PathFor(in interface{}) (string, error) {
	if in == nil {
		return "", errors.New("can not calculate path to nil")
	}

	if s, ok := in.(string); ok {
		return join(s), nil
	}

	if s, ok := in.(Pathable); ok {
		return join(s.ToPath()), nil
	}

	ni, err := name.Interface(in)
	if err != nil {
		return "", err
	}

	to := reflect.TypeOf(in)
	k := to.Kind()
	switch k {
	case reflect.Struct:
		v := reflect.ValueOf(in)
		f := v.FieldByName("Slug")
		if f.IsValid() {
			return byField(ni, f)
		}
		f = v.FieldByName("ID")
		if f.IsValid() {
			return byField(ni, f)
		}
	case reflect.Slice, reflect.Array:
		var paths []string
		v := reflect.ValueOf(in)
		for i := 0; i < v.Len(); i++ {
			rv := v.Index(i)
			s, err := PathFor(rv.Interface())
			if err != nil {
				return "", err
			}
			paths = append(paths, s)
		}
		return join(paths...), nil
	}

	if s, ok := in.(Paramable); ok {
		return join(ni.URL().String(), s.ToParam()), nil
	}

	return "", errors.New("could not convert to path")
}

func byField(ni name.Ident, f reflect.Value) (string, error) {
	ii := f.Interface()
	if ii == nil {
		return "", nil
	}

	zero := reflect.DeepEqual(ii, reflect.Zero(reflect.TypeOf(ii)).Interface())
	if zero {
		return join(ni.URL().String()), nil
	}
	return join(ni.URL().String(), fmt.Sprint(ii)), nil
}

func join(s ...string) string {
	p := path.Join(s...)
	if !strings.HasPrefix(p, "/") {
		p = "/" + p
	}

	return p
}
