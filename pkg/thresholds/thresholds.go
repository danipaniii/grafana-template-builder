package thresholds

// First step in Step slice is gonna be used as "Base" value
type Thresholds struct {
	Mode  string `json:"mode"`
	Steps []Step `json:"steps"`
}

type Step struct {
	Color string  `json:"color"`
	Value float64 `json:"value"`
}

type TTuple struct {
	Color string
	Value float64
}

func BuildThreshold(mode string, values []TTuple) Thresholds {
	steps := []Step{}
	for _, val := range values {
		newStep := Step{
			Color: val.Color,
			Value: float64(val.Value),
		}

		steps = append(steps, newStep)
	}

	return Thresholds{
		Mode:  mode,
		Steps: steps,
	}
}
