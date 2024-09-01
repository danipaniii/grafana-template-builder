package templating

type Templating struct {
	List []Template `json:"list"`
}

type Template struct {
	Current     Current  `json:"current"`
	Hide        int      `json:"hide"`
	IncludeAll  bool     `json:"includeAll"`
	Label       string   `json:"label"`
	Multi       bool     `json:"multi"`
	Name        string   `json:"name"`
	Options     []Option `json:"options"`
	Query       string   `json:"query"`
	SkipUrlSync bool     `json:"skipUrlSync"`
	Type        string   `json:"type"`
}

type Current struct {
	Selected bool   `json:"selected"`
	Text     string `json:"text"`
	Value    string `json:"value"`
}

type Option struct {
	Selected bool   `json:"true"`
	Text     string `json:"text"`
	Value    string `json:"value"`
}
