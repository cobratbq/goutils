package builtin

import (
	"reflect"
)

// ExtractStringKeys extracts keys from a map containing string keys. The
// function uses reflection to extract keys from the map.
func ExtractStringKeys(map_ interface{}) []string {
	keys := reflect.ValueOf(map_).MapKeys()
	result := make([]string, len(keys))
	for i, v := range keys {
		result[i] = v.Interface().(string)
	}
	return result
}
