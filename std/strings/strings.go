// SPDX-License-Identifier: AGPL-3.0-or-later

package strings

import "strings"

// AnyPrefix tests for any of series of prefixes.
func AnyPrefix(s string, prefixes ...string) bool {
	for _, p := range prefixes {
		if strings.HasPrefix(s, p) {
			return true
		}
	}
	return false
}

// OrDefault returns the provided text if non-empty, or the alternative otherwise.
func OrDefault(text, alt string) string {
	if text == "" {
		return alt
	}
	return text
}

// CondText returns a certain text given a provided condition is true or false. Both the true and
// false case texts can be provided.
func CondText(cond bool, truetext, falsetext string) string {
	if cond {
		return truetext
	} else {
		return falsetext
	}
}

// ContainsAll tests if all chars are present in s (subset).
func ContainsAll(s, chars string) bool {
	for _, r := range chars {
		if !strings.ContainsRune(s, r) {
			return false
		}
	}
	return true
}
