package panels

import (
	fieldconfig "github.com/danipaniii/grafana-template-builder/pkg/field-config"
	"github.com/danipaniii/grafana-template-builder/pkg/overrides"
)

type Table struct {
	Type        string                  `json:"type"`
	Title       string                  `json:"title"`
	DataSource  string                  `json:"datasource"`
	Description string                  `json:"description"`
	FieldConfig fieldconfig.FieldConfig `json:"fieldConfig"`
	Overrides   []overrides.Override    `json:"overrides"`
	GridPos     GridPos                 `json:"gridPos"`
}

func (t Table) Render() string {
	t_string, err := Jsonify(t)
	if err != nil {
		return ""
	}

	return t_string
}
