// SPDX-License-Identifier: LGPL-3.0-only

package digit

import (
	"github.com/cobratbq/goutils/std/strings"
)

// FindDigitOverlapping finds the words representing digits in a text, with findings possibly overlapping.
func FindDigitOverlapping(dict map[string]uint8, text string) map[int]uint8 {
	finds := make(map[int]uint8)
	for word, value := range dict {
		// The use of FindOverlapping here gives a bit of a false impression, as we only search
		// overlapping for the same word. Other words are in the next iteration.
		for _, idx := range strings.FindOverlapping(text, word) {
			finds[idx] = value
		}
	}
	return finds
}
