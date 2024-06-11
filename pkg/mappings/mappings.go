package mappings

type Mapping struct {
	Type    string                 `json:"type"`
	Options map[string]interface{} `json:"options"`
}

type Result struct {
	Index int    `json:"index"`
	Text  string `json:"text"`
}

// func jsonify[T any](inp T) {
// 	jsonData, err := json.MarshalIndent(inp, "", "    ")
// 	if err != nil {
// 		return
// 	}

// 	fmt.Println(jsonData)
// }
