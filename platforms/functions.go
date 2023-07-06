package platforms

import (
	"encoding/json"
	"io"
)

func parseJson(b io.ReadCloser) Platform {
	body, _ := io.ReadAll(b)

	var data Platform

	json.Unmarshal(body, &data)

	return data
}
