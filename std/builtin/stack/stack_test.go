package stack

import (
	"testing"

	"github.com/cobratbq/goutils/std/builtin"
	assert "github.com/cobratbq/goutils/std/testing"
)

func TestPush(t *testing.T) {
	var store [3]uint
	stack := store[:1]
	stack2 := builtin.Expect(Push(stack, 2))
	stack2[0] = 1
	assert.Equal(t, stack[0], stack2[0])
	stack3 := builtin.Expect(Push(stack2, 3))
	stack3[0] = 2
	assert.Equal(t, stack[0], stack2[0])
	assert.Equal(t, stack2[0], stack3[0])
	assert.Equal(t, stack3[0], stack[0])
	assert.NotNil(t, builtin.Error(Push(stack3, 4)))
	assert.Equal(t, stack[0], stack2[0])
	assert.Equal(t, stack2[0], stack3[0])
	assert.Equal(t, stack3[0], stack[0])
}
