package dashboard

import "github.com/danipaniii/grafana-template-builder/pkg/panels"

type CreateDashboard struct {
	Dashboard Dashboard `json:"dashboard"`
	FolderUid string    `json:"folderUid"`
	Message   string    `json:"message"`
	Overwrite bool      `json:"overwrite"`
}

type Dashboard struct {
	Title    string         `json:"title"`
	Editable bool           `json:"editable"`
	Panels   []panels.Panel `json:"panels"`
}
