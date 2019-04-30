package web2h

import (
	"cmd/go/internal/web2"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	BodyKey = "Body"

	CopyHeaderKey = "CopyHeader"

	DecodeJSONKey = "DecodeJSON"

	GetKey = "Get"

	HeaderKey = "Header"

	Non200OKKey = "Non200OK"

	ReadAllBodyKey = "ReadAllBody"

	SetHTTPDoForTestingKey = "SetHTTPDoForTesting"
)

func New() hctx.Map {
	return hctx.Map{

		BodyKey: Body,

		CopyHeaderKey: CopyHeader,

		DecodeJSONKey: DecodeJSON,

		GetKey: Get,

		HeaderKey: Header,

		Non200OKKey: Non200OK,

		ReadAllBodyKey: ReadAllBody,

		SetHTTPDoForTestingKey: SetHTTPDoForTesting,
	}
}

var Body = web2.Body

var CopyHeader = web2.CopyHeader

var DecodeJSON = web2.DecodeJSON

var Get = web2.Get

var Header = web2.Header

var Non200OK = web2.Non200OK

var ReadAllBody = web2.ReadAllBody

var SetHTTPDoForTesting = web2.SetHTTPDoForTesting
