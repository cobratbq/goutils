package set

// SupersetOf constructs a closure with set `a`, such that it can be used repeatedly to test whether
// a given set `b` is a superset of `a`.
func SupersetOf[K comparable](a map[K]struct{}) func(set map[K]struct{}) bool {
	return func(b map[K]struct{}) bool {
		return ContainsAll(b, a)
	}
}

// NotSupersetOf constructs a closure with set `a`, such that it can be used repeatedly to test
// whether a given set `b` is NOT a superset of `a`.
func NotSupersetOf[K comparable](a map[K]struct{}) func(set map[K]struct{}) bool {
	return func(b map[K]struct{}) bool {
		return !ContainsAll(b, a)
	}
}

// SubsetOf constructs a closure with set `a`, such that it can be used repeated to test whether
// a given set `b` is a subset of `a`.
func SubsetOf[K comparable](a map[K]struct{}) func(set map[K]struct{}) bool {
	return func(b map[K]struct{}) bool {
		return ContainsAll(a, b)
	}
}

// NotSubsetOf constructs a closure with set `a`, such that it can be used repeated to test whether
// a given set `b` is NOT a subset of `a`.
func NotSubsetOf[K comparable](a map[K]struct{}) func(set map[K]struct{}) bool {
	return func(b map[K]struct{}) bool {
		return !ContainsAll(a, b)
	}
}

// IdenticalTo constructs a closure with set `a`, such that it can be used repeated to test whether
// a given set `b` is identical to `a`.
func IdenticalTo[K comparable](a map[K]struct{}) func(set map[K]struct{}) bool {
	return func(b map[K]struct{}) bool {
		if len(b) != len(a) {
			return false
		}
		return ContainsAll(a, b) && ContainsAll(b, a)
	}
}

// NotIdenticalTo constructs a closure with set `a`, such that it can be used repeated to test
// whether a given set `b` is NOT identical to `a`.
func NotIdenticalTo[K comparable](a map[K]struct{}) func(set map[K]struct{}) bool {
	return func(b map[K]struct{}) bool {
		if len(b) != len(a) {
			return true
		}
		return !ContainsAll(a, b) || !ContainsAll(b, a)
	}
}
