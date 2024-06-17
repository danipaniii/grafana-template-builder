package panels

type Table struct {
	BasePanel
}

func (t Table) Render() string {
	t_string, err := Jsonify(t)
	if err != nil {
		return ""
	}

	return t_string
}
