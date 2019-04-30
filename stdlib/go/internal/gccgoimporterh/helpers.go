package gccgoimporterh

import (
	"go/internal/gccgoimporter"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	GetImporterKey = "GetImporter"
)

func New() hctx.Map {
	return hctx.Map{

		GetImporterKey: GetImporter,
	}
}

var GetImporter = gccgoimporter.GetImporter
