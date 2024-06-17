package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/danipaniii/grafana-template-builder/pkg/dashboard"
	fieldconfig "github.com/danipaniii/grafana-template-builder/pkg/field-config"
	"github.com/danipaniii/grafana-template-builder/pkg/mappings"
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

	bBasePanel := panels.BasePanel{
		Title:      "Test2",
		Type:       "table",
		DataSource: "grafana",
		FieldConfig: fieldconfig.FieldConfig{
			Defaults: fieldconfig.Defaults{
				Thresholds: thresholds.Thresholds{ // provide helper methods for easier generation also for mappings + overrides
					Mode: "absolute",
					Steps: []thresholds.Step{
						{
							Color: "blue",
						},
						{
							Color: "yellow",
							Value: 50,
						},
						{
							Color: "green",
							Value: 80,
						},
					},
				},
				Mappings: []mappings.Mapping{
					{
						Type: "value",
						Options: map[string]interface{}{
							"1": mappings.Result{
								Text: "A",
							},
						},
					},
					{
						Type: "range",
						Options: map[string]interface{}{
							"from": 35.9,
							"to":   36.2,
							"result": mappings.Result{
								Text: "B",
							},
						},
					},
					{
						Type: "regex",
						Options: map[string]interface{}{
							"pattern": "\\d",
							"result": mappings.Result{
								Text: "B",
							},
						},
					},
					{
						Type: "special",
						Options: map[string]interface{}{
							"match": "null",
							"result": mappings.Result{
								Text: "LOL",
							},
						},
					},
				},
			},
		},
		GridPos: panels.GridPos{
			X: 0,
			Y: 0,
			H: 8,
			W: 12,
		},
	}

	// b := panels.Table{
	// 	BasePanel: bBasePanel,
	// }

	new_dashboard := dashboard.CreateDashboard{
		Overwrite: true,
		Dashboard: dashboard.Dashboard{
			Title:    "Hello-Test",
			Panels:   []panels.Panel{aBasePanel, bBasePanel, cBasePanel},
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
