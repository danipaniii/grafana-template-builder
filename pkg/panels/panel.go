package panels

import (
	"encoding/json"

	fieldconfig "github.com/danipaniii/grafana-template-builder/pkg/field-config"
)

type Panel interface {
	Render() map[string]interface{}
}

type BasePanel struct {
	Type        string                  `json:"type"`
	Title       string                  `json:"title"`
	DataSource  string                  `json:"datasource"`
	Description string                  `json:"description"`
	FieldConfig fieldconfig.FieldConfig `json:"fieldConfig"`
	GridPos     GridPos                 `json:"gridPos"`
	Options     map[string]interface{}  `json:"options"`
	Custom      interface{}             `json:"-"`
}

type GridPos struct {
	X int `json:"x"`
	Y int `json:"y"`
	H int `json:"h"`
	W int `json:"w"`
}

func (bp BasePanel) Render() map[string]interface{} {
	return structToMap(bp)
}

func Jsonify[T any](dashboard T) (string, error) {
	jsonData, err := json.MarshalIndent(dashboard, "", "    ")
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}

func RenderCustomFields() {
	// TODO
}

func structToMap(obj interface{}) map[string]interface{} {
	data, _ := json.Marshal(obj)
	var result map[string]interface{}
	json.Unmarshal(data, &result)
	return result
}
