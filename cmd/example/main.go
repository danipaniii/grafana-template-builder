package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/danipaniii/grafana-template-builder/pkg/dashboard"
	fieldconfig "github.com/danipaniii/grafana-template-builder/pkg/field-config"
	"github.com/danipaniii/grafana-template-builder/pkg/mappings"
	"github.com/danipaniii/grafana-template-builder/pkg/overrides"
	"github.com/danipaniii/grafana-template-builder/pkg/panels"
	"github.com/danipaniii/grafana-template-builder/pkg/thresholds"
)

var grafana_url = "http://localhost:3000/api/dashboards/db"

func main() {
	fmt.Println("Hello Package")

	aBasePanel := panels.BasePanel{
		Title:      "Test",
		Type:       "table",
		DataSource: "grafana",
		GridPos: panels.GridPos{
			X: 12,
			Y: 0,
			H: 8,
			W: 12,
		},
		Options: map[string]interface{}{
			"cellHeight": "md",
		},
	}

	cBasePanel := panels.BasePanel{
		Title:      "Test",
		Type:       "timeseries",
		DataSource: "grafana",
		GridPos: panels.GridPos{
			X: 12,
			Y: 9,
			H: 8,
			W: 12,
		},
		Options: map[string]interface{}{
			"cellHeight": "sm",
		},
	}

	// a := panels.Table{
	// 	BasePanel: aBasePanel,
	// }

	bThreshold := thresholds.BuildThreshold("absolute", []thresholds.TTuple{{Color: "blue", Value: 0}, {Color: "yellow", Value: 50}, {Color: "green", Value: 80}})

	bMappings := []mappings.Mapping{}
	bValMappings := mappings.BuildValueMappings([]mappings.ValTuple{{Value: "1", NewValue: "A"}})
	bRngMappings := mappings.BuildRangeMappings([]mappings.RngTuple{{From: 35.0, To: 36.2, NewValue: "B"}})
	bRgxMappings := mappings.BuildRegexMappings([]mappings.RgxTuple{{Pattern: "\\d", NewValue: "C"}})
	bSpcMappings := mappings.BuildSpecialMappings([]mappings.SpcTuple{{Match: "null", NewValue: "LOL"}})

	bMappings = append(bMappings, bValMappings...)
	bMappings = append(bMappings, bRngMappings...)
	bMappings = append(bMappings, bRgxMappings...)
	bMappings = append(bMappings, bSpcMappings...)

	bOverride := overrides.Override{
		Matcher: overrides.Matcher{Id: "byName", Options: "A-series"},
		Properties: []overrides.Property{
			{
				Id:    "unit",
				Value: "m",
			},
		},
	}

	bCustom := map[string]interface{}{
		"fieldConfig": map[string]interface{}{
			"overrides": []overrides.Override{
				{
					Matcher: overrides.Matcher{Id: "byName", Options: "A-series"},
					Properties: []overrides.Property{
						{
							Id: "custom.cellOptions",
							Value: map[string]interface{}{
								"type": "color-background",
								"mode": "basic",
							},
						},
					},
				},
			},
		},
	}

	bBasePanel := panels.BasePanel{
		Title:      "Test2",
		Type:       "table",
		DataSource: "grafana",
		FieldConfig: fieldconfig.FieldConfig{
			Defaults: fieldconfig.Defaults{
				Thresholds: bThreshold,
				Mappings:   bMappings,
			},
			Overrides: []overrides.Override{
				bOverride,
			},
		},
		GridPos: panels.GridPos{
			X: 0,
			Y: 0,
			H: 8,
			W: 12,
		},
		Custom: bCustom,
	}

	new_dashboard := dashboard.CreateDashboard{
		Overwrite: true,
		Dashboard: dashboard.Dashboard{
			Title:    "Hello-Test",
			Panels:   []interface{}{aBasePanel.Render(), bBasePanel.Render(), cBasePanel.Render()},
			Editable: true,
		},
	}

	jsonDashboard, err := jsonify(new_dashboard) // make this method part of dashboard? Also for panel auto id? if necessary?
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	fmt.Println(string(jsonDashboard))

	// Building Request to push to grafana -> Create API methods
	req, err := http.NewRequest("POST", grafana_url, bytes.NewBuffer(jsonDashboard))
	if err != nil {
		fmt.Println("Error creating request")
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+"local_grafana_token")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Println("Dashboard JSON pushed successfully!")
	} else {
		fmt.Println("Failed to push Dashboard JSON. Status:", resp.StatusCode)
	}
}

func jsonify[T any](dashboard T) ([]byte, error) {
	jsonData, err := json.MarshalIndent(dashboard, "", "    ")
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}
