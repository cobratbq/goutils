package builtin

// Require check required condition and panics if condition does not hold.
func Require(condition bool, message string) {
	if !condition {
		panic(message)
	}
}

// RequireNonNil checks if provided value is nil, if so panics with provided
// message.
func RequireNonNil(val interface{}, message string) {
	if val == nil {
		panic(message)
	}
}
