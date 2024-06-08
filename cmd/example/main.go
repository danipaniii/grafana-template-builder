package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/danipaniii/grafana-template-builder/pkg/dashboard"
	"github.com/danipaniii/grafana-template-builder/pkg/panels"
)

var grafana_url = "http://localhost:3000/api/dashboards/db"

func main() {
	fmt.Println("Hello Package")

	a := panels.Table{
		Title: "Test",
		Type:  "table",
		//DataSource: "grafana",
	}

	new_dashboard := dashboard.CreateDashboard{
		Overwrite: true,
		Dashboard: dashboard.Dashboard{
			Title:    "Hello-Test",
			Panels:   []panels.Panel{a},
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
