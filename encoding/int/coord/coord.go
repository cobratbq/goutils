package coord

func Encode2D(length, first, second int) int {
	return second*length + first
}

func Decode2D(length, index int) (int, int) {
	return index % length, index / length
}

func Encode2DUint(length, first, second uint) int {
	return int(second*length + first)
}

func Decode2DUint(length uint, index int) (uint, uint) {
	return uint(index) % length, uint(index) / length
}
