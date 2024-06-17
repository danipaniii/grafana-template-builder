package fieldconfig

import (
	"github.com/danipaniii/grafana-template-builder/pkg/mappings"
	"github.com/danipaniii/grafana-template-builder/pkg/thresholds"
)

type FieldConfig struct {
	Defaults Defaults `json:"defaults"`
}

type Defaults struct {
	Thresholds thresholds.Thresholds `json:"thresholds"`
	Mappings   []mappings.Mapping    `json:"mappings"`
	//Custom     map[string]interface{} `json:"custom"`
}
