// SPDX-License-Identifier: LGPL-3.0-only

package strings

import "strings"

func FindNonOverlapping(text string, substr string) []int {
	var finds []int
	var idx, next int
	for idx >= 0 && idx < len(text) {
		if next = strings.Index(text[idx:], substr); next < 0 {
			break
		}
		finds = append(finds, idx+next)
		idx += next + len(substr)
	}
	return finds
}

func FindOverlapping(text string, substr string) []int {
	var finds []int
	var idx, next int
	for idx >= 0 && idx < len(text) {
		if next = strings.Index(text[idx:], substr); next < 0 {
			break
		}
		finds = append(finds, idx+next)
		idx += next + 1
	}
	return finds
}

// AnyPrefix tests for any of provided prefixes.
func AnyPrefix(text string, prefixes ...string) bool {
	for p := range prefixes {
		if strings.HasPrefix(text, prefixes[p]) {
			return true
		}
	}
	return false
}

// AnySuffix tests for any of provided suffixes.
func AnySuffix(text string, suffixes ...string) bool {
	for s := range suffixes {
		if strings.HasSuffix(text, suffixes[s]) {
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
