package panels

import "encoding/json"

type Panel interface {
	Render() string
}

func Jsonify[T any](dashboard T) (string, error) {
	jsonData, err := json.MarshalIndent(dashboard, "", "    ")
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}
