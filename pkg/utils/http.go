package utils

import (
	"bytes"
	"encoding/json"
	"io"
)

func ConvertToReader(data interface{}) (io.Reader, error) {
	paramsJSON, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	// Tạo một đối tượng io.Reader từ dữ liệu JSON
	paramsReader := bytes.NewBuffer(paramsJSON)
	return paramsReader, nil
}
