package stack

import (
	"github.com/cobratbq/goutils/std/errors"
)

// IsEmpty returns true iff stack is empty.
// FIXME needs testing
func IsEmpty[T any](stack []T) bool {
	return len(stack) == 0
}

// IsFull returns true iff stack is full, i.e. at capacity.
// FIXME needs testing
func IsFull[T any](stack []T) bool {
	return len(stack) == cap(stack)
}

// ErrFull indicates that stack is at full capacity, either at present or after operation.
var ErrFull = errors.NewStringError("stack is full")

// Push pushes an element onto the stack.
//
// Returns the updated stack, or ErrFull if stack is at full capacity.
// FIXME needs testing
func Push[T any](stack []T, val T) ([]T, error) {
	if len(stack) == cap(stack) {
		return stack, ErrFull
	}
	n := len(stack)
	// Note: cannot use `append` as it would reallocate the slice if we reach capacity
	stack = stack[:n+1]
	stack[n] = val
	return stack, nil
}

// PushMany pushes a sequence of elements onto the stack.
//
// Returns updated stack, or ErrFull if sequence is larger than capacity.
// FIXME needs testing
func PushMany[T any](stack []T, vals []T) ([]T, error) {
	if len(stack)+len(vals) > cap(stack) {
		return stack, ErrFull
	}
	n := len(stack)
	stack = stack[:n+len(vals)]
	for i, v := range vals {
		stack[n+i] = v
	}
	return stack, nil
}

// ErrEmpty indicates that stack is empty, either at present or after the operation.
var ErrEmpty = errors.NewStringError("stack is empty")

// Pop pops an element off the stack.
//
// Returns updated stack and element, or ErrEmpty if stack is empty.
// FIXME needs testing
func Pop[T any](stack []T) ([]T, T, error) {
	if len(stack) == 0 {
		var zero T
		return stack, zero, ErrEmpty
	}
	v := stack[len(stack)-1]
	return stack[:len(stack)-1], v, nil
}

// PopN pops `n` elements off the stack.
//
// Returns updated stack and elements, or ErrEmpty if more elements are popped than are present.
// FIXME needs testing
func PopN[T any](stack []T, n uint) ([]T, []T, error) {
	if uint(len(stack)) < n {
		return stack, nil, ErrEmpty
	}
	popped := make([]T, n)
	// TODO start with last on stack, or end with last on stack ... you'd think that popping multiple would result in reverse order result?
	for i := uint(0); i < n; i++ {
		popped[i] = stack[uint(len(stack))-1-i]
	}
	return stack[:uint(len(stack))-n], popped, nil
}

// Peek copies the top element of the stack.
//
// Returns the element on the top of the stack or ErrEmpty if stack is empty.
// FIXME needs testing
func Peek[T any](stack []T) (T, error) {
	if len(stack) == 0 {
		var zero T
		return zero, ErrEmpty
	}
	return stack[len(stack)-1], nil
}

// PeekN peeks for the top `n` elements on the stack.
//
// Returns slice of top `n` elements, or ErrEmpty if more elements are requested than are present.
// FIXME needs testing
func PeekN[T any](stack []T, n uint) ([]T, error) {
	if uint(len(stack)) < n {
		return nil, ErrEmpty
	}
	peeked := make([]T, n)
	// TODO start with last on stack, or end with last on stack ... you'd think that popping multiple would result in reverse order result?
	for i := uint(0); i < n; i++ {
		peeked[i] = stack[uint(len(stack))-1-i]
	}
	return peeked, nil
}
