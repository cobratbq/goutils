package assert

func False(expected bool) {
	True(!expected)
}

func True(expected bool) {
	if !expected {
		panic("assertion failed")
	}
}
