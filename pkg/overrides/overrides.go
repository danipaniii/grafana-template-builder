package overrides

type Matcher struct {
	Id      string `json:"id"`
	Options string `json:"options"`
}

type Property struct {
	Id    string      `json:"id"`
	Value interface{} `json:"value"`
}

type Override struct {
	Matcher    Matcher    `json:"matcher"`
	Properties []Property `json:"properties"`
}
