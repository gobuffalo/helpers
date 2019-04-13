package encoders

import (
	"encoding/json"
	"html/template"
)

func ToJSON(v interface{}) (template.HTML, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return template.HTML(b), nil
}
