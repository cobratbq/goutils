package errors

import (
	"fmt"

	"github.com/cobratbq/goutils/std/builtin"
)

// RequireSuccess checks that err is nil. If the error is non-nil, it will
// panic. `message` can have '%v' format specifier so that it can be
// substituted with the error message.
func RequireSuccess(err error, message string) {
	builtin.Require(err == nil, fmt.Sprintf(message, err))
}
