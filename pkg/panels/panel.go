package panels

import (
	"encoding/json"

	fieldconfig "github.com/danipaniii/grafana-template-builder/pkg/field-config"
	"github.com/danipaniii/grafana-template-builder/pkg/transformations"
)

type Panel interface {
	Render() map[string]interface{}
}

type BasePanel struct {
	Type            string                           `json:"type"`
	Title           string                           `json:"title"`
	DataSource      string                           `json:"datasource"`
	Description     string                           `json:"description"`
	FieldConfig     fieldconfig.FieldConfig          `json:"fieldConfig"`
	GridPos         GridPos                          `json:"gridPos"`
	Options         map[string]interface{}           `json:"options"`
	Transformations []transformations.Transformation `json:"transformations"`
	Override        map[string]interface{}           `json:"-"`
}

type GridPos struct {
	X int `json:"x"`
	Y int `json:"y"`
	H int `json:"h"`
	W int `json:"w"`
}

func (bp BasePanel) Render() map[string]interface{} {
	panelMap := structToMap(bp)
	return mergeMaps(panelMap, bp.Override)
}

func structToMap(obj interface{}) map[string]interface{} {
	data, _ := json.Marshal(obj)
	var result map[string]interface{}
	json.Unmarshal(data, &result)
	return result
}

func mergeMaps(mapA, mapB map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})

	// Iterate over all keys in mapA
	for key, valA := range mapA {
		if valB, ok := mapB[key]; ok {
			// If the key exists in both maps
			switch valA := valA.(type) {
			case map[string]interface{}:
				// If both values are maps, merge them recursively
				if nestedMapB, ok := valB.(map[string]interface{}); ok {
					result[key] = mergeMaps(valA, nestedMapB)
				} else {
					// If not, prefer mapB's value
					result[key] = valB
				}
			case []interface{}:
				// If both values are slices, concatenate them
				if sliceB, ok := valB.([]interface{}); ok {
					result[key] = append(valA, sliceB...)
				} else {
					// If not, prefer mapB's value
					result[key] = valB
				}
			default:
				// If values are not maps or slices, prefer mapB's value
				result[key] = valB
			}
		} else {
			// If the key exists only in mapA, copy it
			result[key] = valA
		}
	}

	// Add any keys that exist only in mapB
	for key, valB := range mapB {
		if _, ok := mapA[key]; !ok {
			result[key] = valB
		}
	}

	return result
}

func Jsonify[T any](dashboard T) (string, error) {
	jsonData, err := json.MarshalIndent(dashboard, "", "    ")
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}
