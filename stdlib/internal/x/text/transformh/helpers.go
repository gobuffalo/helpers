package transformh

import (
	"internal/x/text/transform"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	AppendKey = "Append"

	BytesKey = "Bytes"

	ChainKey = "Chain"

	NewReaderKey = "NewReader"

	NewWriterKey = "NewWriter"

	RemoveFuncKey = "RemoveFunc"

	StringKey = "String"
)

func New() hctx.Map {
	return hctx.Map{

		AppendKey: Append,

		BytesKey: Bytes,

		ChainKey: Chain,

		NewReaderKey: NewReader,

		NewWriterKey: NewWriter,

		RemoveFuncKey: RemoveFunc,

		StringKey: String,
	}
}

// Append appends the result of converting src[:n] using t to dst, where
// n &lt;= len(src), If err == nil, n will be len(src). It calls Reset on t.
var Append = transform.Append

// Bytes returns a new byte slice with the result of converting b[:n] using t,
// where n &lt;= len(b). If err == nil, n will be len(b). It calls Reset on t.
var Bytes = transform.Bytes

// Chain returns a Transformer that applies t in sequence.
var Chain = transform.Chain

// NewReader returns a new Reader that wraps r by transforming the bytes read
// via t. It calls Reset on t.
var NewReader = transform.NewReader

// NewWriter returns a new Writer that wraps w by transforming the bytes written
// via t. It calls Reset on t.
var NewWriter = transform.NewWriter

// Deprecated: use runes.Remove instead.
var RemoveFunc = transform.RemoveFunc

// String returns a string with the result of converting s[:n] using t, where
// n &lt;= len(s). If err == nil, n will be len(s). It calls Reset on t.
var String = transform.String
