package interval

// UintInterval represents an unsigned integer interval.
type UintInterval struct {
	start uint
	size  uint
}

// NewUint constructs a new UintInterval instance.
func NewUint(start uint, size uint) UintInterval {
	return UintInterval{start, size}
}

// First returns the start/first value within the interval.
func (r *UintInterval) First() uint {
	return r.start
}

// Last returns the last value within the interval. (i.e. inclusive interval end)
func (r *UintInterval) Last() uint {
	return r.start + r.size - 1
}

// Start returns the start/first value within the interval.
func (r *UintInterval) Start() uint {
	return r.start
}

// End returns the first value outside of (past) the interval. (i.e. exclusive interval end)
func (r *UintInterval) End() uint {
	// FIXME is end inclusive?
	return r.start + r.size
}

// Size returns the size of the interval.
func (r *UintInterval) Size() uint {
	return r.size
}

// Equal returns true iff two intervals are equal.
func (r *UintInterval) Equal(o *UintInterval) bool {
	return *r == *o
}

// Contains returns true iff `o` is contained within the interval.
func (r *UintInterval) Contains(o *UintInterval) bool {
	return r.start <= o.start && o.Last() <= r.Last()
}

// ContainsValue returns true iff the value is contained within the interval.
func (r *UintInterval) ContainsValue(v uint) bool {
	return r.start <= v && v <= r.Last()
}

// Overlaps returns true iff there is some overlap between `o` and the interval.
func (r *UintInterval) Overlaps(o *UintInterval) bool {
	return (r.start <= o.start && r.Last() >= o.start) || (o.start <= r.start && o.Last() >= r.start)
}
