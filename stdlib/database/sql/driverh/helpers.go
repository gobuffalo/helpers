package driverh

import (
	"database/sql/driver"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	IsScanValueKey = "IsScanValue"

	IsValueKey = "IsValue"
)

func New() hctx.Map {
	return hctx.Map{

		IsScanValueKey: IsScanValue,

		IsValueKey: IsValue,
	}
}

// IsScanValue is equivalent to IsValue.
// It exists for compatibility.
var IsScanValue = driver.IsScanValue

// IsValue reports whether v is a valid Value parameter type.
var IsValue = driver.IsValue
