package row

import "github.com/danipaniii/grafana-template-builder/pkg/panels"

type Row struct {
	Id        int                `json:"id"`
	Title     string             `json:"title"`
	Type      string             `json:"type"`
	Panels    []panels.BasePanel `json:"panels"`
	Collapsed bool               `json:"collapsed"`
	GridPos   panels.GridPos     `json:"gridPos"`
}

func CreateSimpleRow(title string, collapsed bool, gridPos panels.GridPos) Row {
	return Row{
		Type:      "row",
		Title:     title,
		Collapsed: collapsed,
		GridPos:   gridPos,
	}
}
