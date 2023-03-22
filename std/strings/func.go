// SPDX-License-Identifier: AGPL-3.0-or-later

package strings

func SupersetOf(chars string) func(s string) bool {
	return func(s string) bool {
		return ContainsAll(s, chars)
	}
}

func NotSupersetOf(chars string) func(s string) bool {
	return func(s string) bool {
		return !ContainsAll(s, chars)
	}
}

func SubsetOf(chars string) func(s string) bool {
	return func(s string) bool {
		return ContainsAll(chars, s)
	}
}

func NotSubsetOf(chars string) func(s string) bool {
	return func(s string) bool {
		return !ContainsAll(chars, s)
	}
}

func IdenticalTo(chars string) func(s string) bool {
	// TODO consider if we want to test for equal length first? Set should not contain duplicates, so then this requirement holds.
	return func(s string) bool {
		return ContainsAll(s, chars) && ContainsAll(chars, s)
	}
}

func NotIdenticalTo(chars string) func(s string) bool {
	// TODO consider if we want to test for equal length first? Set should not contain duplicates, so then this requirement holds.
	return func(s string) bool {
		return !ContainsAll(s, chars) || !ContainsAll(chars, s)
	}
}
