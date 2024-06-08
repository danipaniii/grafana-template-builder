package panels

type Table struct {
	Type        string `json:"type"`
	Title       string `json:"title"`
	DataSource  string `json:"datasource"`
	Description string `json:"description"`
}

type TableOptions struct {
	Type        string
	Title       string
	DataSource  string
	Description string
}

func (t Table) Render() string {
	t_string, err := Jsonify(t)
	if err != nil {
		return ""
	}

	return t_string
}

func CreatePanel(options TableOptions) *Table {
	return &Table{
		Type:        options.Type,
		Title:       options.Title,
		DataSource:  options.DataSource,
		Description: options.Description,
	}
}
