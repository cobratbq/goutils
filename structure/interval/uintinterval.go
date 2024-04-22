package interval

// UintInterval represents a closed interval.
type UintInterval struct {
	start uint
	size  uint
}

func NewUintInterval(start uint, size uint) UintInterval {
	return UintInterval{start, size}
}

func (r *UintInterval) Start() uint {
	return r.start
}

func (r *UintInterval) End() uint {
	return r.start + r.size
}

func (r *UintInterval) Size() uint {
	return r.size
}

func (r *UintInterval) Equal(o *UintInterval) bool {
	return *r == *o
}

func (r *UintInterval) Contains(o *UintInterval) bool {
	return r.start <= o.start && o.End() <= r.End()
}

func (r *UintInterval) ContainsValue(v uint) bool {
	return r.start <= v && v <= r.End()
}

func (r *UintInterval) Overlaps(o *UintInterval) bool {
	return (r.start <= o.start && r.End() >= o.start) || (o.start <= r.start && o.End() >= r.start)
}
