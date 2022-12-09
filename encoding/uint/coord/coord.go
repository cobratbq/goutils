package coord

func Encode2D(length, first, second uint) uint {
	return second*length + first
}

func Decode2D(length, index uint) (uint, uint) {
	return index % length, index / length
}

func Encode2DInt(length, first, second int) uint {
	return uint(second)*uint(length) + uint(first)
}

func Decode2DInt(length int, index uint) (int, int) {
	return int(index % uint(length)), int(index / uint(length))
}
