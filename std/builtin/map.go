// SPDX-License-Identifier: LGPL-3.0-or-later
package builtin

import (
	"reflect"
)

// ExtractStringKeys extracts keys from a map containing string keys. The
// function uses reflection to extract keys from the map.
func ExtractStringKeys(stringmap interface{}) []string {
	keys := reflect.ValueOf(stringmap).MapKeys()
	result := make([]string, len(keys))
	for i, v := range keys {
		result[i] = v.Interface().(string)
	}
	return result
}
