// SPDX-License-Identifier: LGPL-3.0-only

package slices

import (
	"bytes"
	"testing"

	"github.com/cobratbq/goutils/std/crypto/rand"
	assert "github.com/cobratbq/goutils/std/testing"
)

func TestIncrement(t *testing.T) {
	testdata := []struct {
		input    []uint8
		leresult []uint8
		beresult []uint8
	}{
		{input: []uint8{}, leresult: []uint8{}, beresult: []uint8{}},
		{input: []uint8{0}, leresult: []uint8{1}, beresult: []uint8{1}},
		{input: []uint8{100}, leresult: []uint8{101}, beresult: []uint8{101}},
		{input: []uint8{100, 0}, leresult: []uint8{101, 0}, beresult: []uint8{100, 1}},
		{input: []uint8{100, 0, 55}, leresult: []uint8{101, 0, 55}, beresult: []uint8{100, 0, 56}},
		{input: []uint8{255, 255, 255}, leresult: []uint8{0, 0, 0}, beresult: []uint8{0, 0, 0}},
		{input: []uint8{255, 255, 254}, leresult: []uint8{0, 0, 255}, beresult: []uint8{255, 255, 255}},
	}
	for d := range testdata {
		incle := bytes.Clone(testdata[d].input)
		IncrementLE(incle)
		assert.SlicesEqual(t, testdata[d].leresult, incle)
		if len(incle) > 0 {
			assert.False(t, Equal(testdata[d].input, incle))
		}
		incbe := bytes.Clone(testdata[d].input)
		IncrementBE(incbe)
		assert.SlicesEqual(t, testdata[d].beresult, incbe)
		if len(incle) > 0 {
			assert.False(t, Equal(testdata[d].input, incbe))
		}
		if len(incle) == 1 {
			assert.SlicesEqual(t, incle, incbe)
		}
	}
}

func TestDecrement(t *testing.T) {
	testdata := []struct {
		input    []uint8
		leresult []uint8
		beresult []uint8
	}{
		{input: []uint8{}, leresult: []uint8{}, beresult: []uint8{}},
		{input: []uint8{0}, leresult: []uint8{255}, beresult: []uint8{255}},
		{input: []uint8{100}, leresult: []uint8{99}, beresult: []uint8{99}},
		{input: []uint8{100, 0}, leresult: []uint8{99, 0}, beresult: []uint8{99, 255}},
		{input: []uint8{100, 0, 55}, leresult: []uint8{99, 0, 55}, beresult: []uint8{100, 0, 54}},
		{input: []uint8{255, 255, 255}, leresult: []uint8{254, 255, 255}, beresult: []uint8{255, 255, 254}},
		{input: []uint8{255, 255, 0}, leresult: []uint8{254, 255, 0}, beresult: []uint8{255, 254, 255}},
	}
	for d := range testdata {
		incle := bytes.Clone(testdata[d].input)
		DecrementLE(incle)
		assert.SlicesEqual(t, testdata[d].leresult, incle)
		if len(incle) > 0 {
			assert.False(t, Equal(testdata[d].input, incle))
		}
		incbe := bytes.Clone(testdata[d].input)
		DecrementBE(incbe)
		assert.SlicesEqual(t, testdata[d].beresult, incbe)
		if len(incle) > 0 {
			assert.False(t, Equal(testdata[d].input, incbe))
		}
		if len(incle) == 1 {
			assert.SlicesEqual(t, incle, incbe)
		}
	}
}

func TestSliceExtendFailing(t *testing.T) {
	defer assert.RequirePanic(t)
	s := make([]byte, 0, 4)
	s = Extend(s, 'a', 'b', 'c', 'd', 'e')
	t.FailNow()
}

func TestSliceExtendFromFailing(t *testing.T) {
	defer assert.RequirePanic(t)
	s := make([]byte, 0, 4)
	s = ExtendFrom(s, []byte{'a', 'b', 'c', 'd', 'e'})
	t.FailNow()
}

func TestSliceExtendMaybeFailing(t *testing.T) {
	var err error
	s := make([]byte, 0, 4)
	s, err = ExtendMaybe(s[0:0:4], 'a', 'b', 'c', 'd', 'e')
	assert.NotNil(t, err)
	assert.NotNil(t, s)
	assert.Equal(t, 0, len(s))
	s, err = ExtendFromMaybe(s[0:0:4], []byte{'a', 'b', 'c', 'd', 'e'})
	assert.NotNil(t, err)
	assert.NotNil(t, s)
	assert.Equal(t, 0, len(s))
}

func TestSliceReversing(t *testing.T) {
	var random [33]byte
	rand.MustReadBytes(random[:])
	revved := Reversed(random[:])
	assert.False(t, bytes.Equal(revved, random[:]))
	revrevved := Reversed(revved)
	assert.SlicesEqual(t, random[:], revrevved)
	Reverse(random[:])
	assert.SlicesEqual(t, revved, random[:])
}

func TestUniformDimensions2D(t *testing.T) {
	assert.True(t, UniformDimensions2D([][]uint{}))
	assert.True(t, UniformDimensions2D([][]uint{{}, {}, {}}))
	assert.True(t, UniformDimensions2D([][]uint{{1}, {2}, {3}}))
	assert.False(t, UniformDimensions2D([][]uint{{1}, {2}, {}}))
	assert.False(t, UniformDimensions2D([][]uint{{1, 3, 4}, {2, 3}, {}}))
}
