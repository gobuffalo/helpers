package dirhashh

import (
	"cmd/go/internal/dirhash"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	DirFilesKey = "DirFiles"

	Hash1Key = "Hash1"

	HashDirKey = "HashDir"

	HashZipKey = "HashZip"
)

func New() hctx.Map {
	return hctx.Map{

		DirFilesKey: DirFiles,

		Hash1Key: Hash1,

		HashDirKey: HashDir,

		HashZipKey: HashZip,
	}
}

var DirFiles = dirhash.DirFiles

var Hash1 = dirhash.Hash1

var HashDir = dirhash.HashDir

var HashZip = dirhash.HashZip
