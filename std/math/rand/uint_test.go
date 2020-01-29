package rand

import (
	"testing"
)

func TestMathRandUint64NonZero(t *testing.T) {
	v1 := MathRandUint64NonZero()
	v2 := MathRandUint64NonZero()
	v3 := MathRandUint64NonZero()
	if v1 == v2 || v2 == v3 || v1 == v3 {
		t.FailNow()
	}
}

func TestMathRandUint64NonZeroNotZero(t *testing.T) {
	for i := 0; i < 100; i++ {
		v := MathRandUint64NonZero()
		if v == 0 {
			t.Fail()
		}
	}
}
