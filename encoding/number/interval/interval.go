package interval

import "github.com/cobratbq/goutils/types"

func ExpandClosed[N types.Integer](a, b N) []N {
	if a == b {
		return []N{a}
	}
	var parts []N
	for i := a; i <= b; i++ {
		parts = append(parts, i)
	}
	for i := a; i >= b; i-- {
		parts = append(parts, i)
	}
	return parts
}
