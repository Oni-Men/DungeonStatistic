package text

import (
	"bytes"
	"encoding/json"
)

func FormatJSON(data []byte) string {
	var buf bytes.Buffer
	err := json.Indent(&buf, data, "", "  ")
	if err != nil {
		return err.Error()
	}

	return buf.String()
}
