// SPDX-License-Identifier: LGPL-3.0-only

//go:build !disable_assert

package assert

import (
	"testing"

	assert "github.com/cobratbq/goutils/std/testing"
	"github.com/cobratbq/goutils/types"
)

func TestTrue(t *testing.T) {
	True(true)
}

func TestTruePanics(t *testing.T) {
	defer assert.RequirePanic(t)
	True(false)
	t.FailNow()
}

func TestFalse(t *testing.T) {
	False(false)
}

func TestFalsePanics(t *testing.T) {
	defer assert.RequirePanic(t)
	False(true)
	t.FailNow()
}

func TestPositive(t *testing.T) {
	Positive(1)
	Positive(2)
	Positive(10)
	Positive(types.MaxInt - 1)
	Positive(types.MaxInt)
}

func TestPositiveZero(t *testing.T) {
	defer assert.RequirePanic(t)
	Positive(0)
	t.FailNow()
}

func TestPositiveNegativeOne(t *testing.T) {
	defer assert.RequirePanic(t)
	Positive(-1)
	t.FailNow()
}

func TestPositiveMinInt(t *testing.T) {
	defer assert.RequirePanic(t)
	Positive(types.MinInt)
	t.FailNow()
}

func TestNonNegativeNegativeOne(t *testing.T) {
	defer assert.RequirePanic(t)
	NonNegative(-1)
	t.FailNow()
}

func TestNonNegativeMinInt(t *testing.T) {
	defer assert.RequirePanic(t)
	NonNegative(types.MinInt)
	t.FailNow()
}

func TestNonNegative(t *testing.T) {
	NonNegative(0)
	NonNegative(1)
	NonNegative(100)
	NonNegative(types.MaxInt)
}

func TestNegative(t *testing.T) {
	Negative(-1)
	Negative(-10)
	Negative(types.MinInt)
}

func TestNegativeZero(t *testing.T) {
	defer assert.RequirePanic(t)
	Negative(0)
	t.FailNow()
}

func TestNegativePositiveOne(t *testing.T) {
	defer assert.RequirePanic(t)
	Negative(1)
	t.FailNow()
}

func TestNegativeMaxInt(t *testing.T) {
	defer assert.RequirePanic(t)
	Negative(types.MaxInt)
	t.FailNow()
}

func TestEmptySlice(t *testing.T) {
	EmptySlice([]byte{})
}

func TestEmptyMap(t *testing.T) {
	m := make(map[uint]string, 0)
	EmptyMap(m)
}
