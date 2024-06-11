package thresholds

// First step in Step slice is gonna be used as "Base" value
type Thresholds struct {
	Mode  string `json:"mode"`
	Steps []Step `json:"steps"`
}

type Step struct {
	Color string  `json:"color"`
	Value float32 `json:"value"`
}
