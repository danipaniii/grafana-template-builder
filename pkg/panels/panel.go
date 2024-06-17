package panels

import (
	"encoding/json"

	fieldconfig "github.com/danipaniii/grafana-template-builder/pkg/field-config"
	"github.com/danipaniii/grafana-template-builder/pkg/overrides"
)

type Panel interface {
	Render() string
}

type BasePanel struct {
	Type        string                  `json:"type"`
	Title       string                  `json:"title"`
	DataSource  string                  `json:"datasource"`
	Description string                  `json:"description"`
	FieldConfig fieldconfig.FieldConfig `json:"fieldConfig"`
	Overrides   []overrides.Override    `json:"overrides"`
	GridPos     GridPos                 `json:"gridPos"`
	Options     map[string]interface{}  `json:"options"`
}

type GridPos struct {
	X int `json:"x"`
	Y int `json:"y"`
	H int `json:"h"`
	W int `json:"w"`
}

func (bp BasePanel) Render() string {
	bp_string, err := Jsonify(bp)
	if err != nil {
		return ""
	}

	return bp_string
}

func Jsonify[T any](dashboard T) (string, error) {
	jsonData, err := json.MarshalIndent(dashboard, "", "    ")
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}
