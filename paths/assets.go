package paths

import (
	"encoding/json"
	"path/filepath"

	"github.com/gobuffalo/syncx"
)

var assets = syncx.StringMap{}

func Asset(file string) string {
	filePath, ok := assets.Load(file)
	if filePath == "" || !ok {
		filePath = file
	}
	return filepath.ToSlash(filepath.Join("/assets", filePath))
}

func loadManifest(manifest string) error {
	m := map[string]string{}

	err := json.Unmarshal([]byte(manifest), &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		assets.Store(k, v)
	}
	return nil
}
