package transformations

type Transformation struct {
	Id      string      `json:"id"`
	Options interface{} `json:"options"`
}

type OrganizeOptions struct {
	RenameByName  map[string]string `json:"renameByName"`
	ExcludeByName map[string]bool   `json:"excludeByName"`
	IncludeByName map[string]bool   `json:"includeByName"`
	IndexByname   map[string]int    `json:"indexByName"`
}

type GroupByOptions struct {
	Fields map[string]FieldOptions `json:"fields"`
}

type FieldOptions struct {
	Aggregations []string `json:"aggregations"`
	Operation    string   `json:"operation"`
}

type CalculateFieldOptions struct {
	Alias  string                 `json:"alias"`
	Mode   string                 `json:"mode"`
	Binary map[string]interface{} `json:"binary"`
	Unary  map[string]interface{} `json:"unary"`
	Reduce map[string]interface{} `json:"reduce"`
}

type ConcatenateOptions struct {
	FrameNameLabel string `json:"frameNameLabel"`
	FrameNameMode  string `json:"frameNameMode"`
}
