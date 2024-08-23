package mappings

type Mapping struct {
	Type    string                 `json:"type"`
	Options map[string]interface{} `json:"options"`
}

type Result struct {
	Index int    `json:"index"`
	Text  string `json:"text"`
}

type ValTuple struct {
	Value    string
	NewValue string
}

type RngTuple struct {
	From     float64
	To       float64
	NewValue string
}

type RgxTuple struct {
	Pattern  string
	NewValue string
}

type SpcTuple struct {
	Match    string
	NewValue string
}

func BuildValueMappings(values []ValTuple) []Mapping {
	mappings := []Mapping{}

	for _, val := range values {
		newMapping := Mapping{
			Type: "value",
			Options: map[string]interface{}{
				val.Value: Result{
					Text: val.NewValue,
				},
			},
		}

		mappings = append(mappings, newMapping)
	}

	return mappings
}

func BuildRangeMappings(values []RngTuple) []Mapping {
	mappings := []Mapping{}

	for _, val := range values {
		newMapping := Mapping{
			Type: "range",
			Options: map[string]interface{}{
				"from": val.From,
				"to":   val.To,
				"result": Result{
					Text: val.NewValue,
				},
			},
		}

		mappings = append(mappings, newMapping)
	}

	return mappings
}

func BuildRegexMappings(values []RgxTuple) []Mapping {
	mappings := []Mapping{}

	for _, val := range values {
		newMapping := Mapping{
			Type: "regex",
			Options: map[string]interface{}{
				"pattern": val.Pattern,
				"result": Result{
					Text: val.NewValue,
				},
			},
		}

		mappings = append(mappings, newMapping)
	}

	return mappings
}

func BuildSpecialMappings(values []SpcTuple) []Mapping {
	mappings := []Mapping{}

	for _, val := range values {
		newMapping := Mapping{
			Type: "special",
			Options: map[string]interface{}{
				"match": val.Match,
				"result": Result{
					Text: val.NewValue,
				},
			},
		}

		mappings = append(mappings, newMapping)
	}

	return mappings
}
