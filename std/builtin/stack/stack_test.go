package stack

import (
	"testing"

	"github.com/cobratbq/goutils/std/builtin"
	"github.com/cobratbq/goutils/std/errors"
	assert "github.com/cobratbq/goutils/std/testing"
)

func TestPushOntoFullStack(t *testing.T) {
	var store [3]byte
	stack := store[:3]
	assert.True(t, IsFull(stack))
	assert.IsError(t, errors.ErrOverflow, builtin.Error(Push(stack, 1)))
}

func TestPushOntoEmptyStack(t *testing.T) {
	var store [3]byte
	stack := store[:0]
	assert.True(t, IsEmpty(stack))
	stack, err := Push(stack, 0)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(stack))
}

func TestPushMany(t *testing.T) {
	var store [3]byte = [...]byte{1, 4, 5}
	stack := store[:0]
	stack, err := PushMany(stack, []byte{2, 3})
	assert.Nil(t, err)
}

func TestPopFromEmptyStack(t *testing.T) {
	var store [3]byte
	stack := store[:0]
	assert.True(t, IsEmpty(stack))
	assert.IsError(t, errors.ErrUnderflow, builtin.Error2(Pop(stack)))
}

func TestPopFromFullStack(t *testing.T) {
	var store [3]byte = [...]byte{1, 2, 3}
	stack := store[:3]
	assert.True(t, IsFull(stack))
	stack, v, err := Pop(stack)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(stack))
	assert.Equal(t, 3, v)
}

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
