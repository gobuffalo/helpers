package debug

import (
	"fmt"
)

func Inspect(v interface{}) string {
	return fmt.Sprintf("%+v", v)
}
