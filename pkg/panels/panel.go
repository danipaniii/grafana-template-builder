package panels

import "encoding/json"

type Panel interface {
	Render() string
}

type GridPos struct {
	X int `json:"x"`
	Y int `json:"y"`
	H int `json:"h"`
	W int `json:"w"`
}

func Jsonify[T any](dashboard T) (string, error) {
	jsonData, err := json.MarshalIndent(dashboard, "", "    ")
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}
