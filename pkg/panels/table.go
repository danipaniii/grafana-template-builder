package panels

type Table struct {
	Type        string `json:"type"`
	Title       string `json:"title"`
	DataSource  string `json:"datasource"`
	Description string `json:"description"`
}

func (t Table) Render() string {
	t_string, err := Jsonify(t)
	if err != nil {
		return ""
	}

	return t_string
}
