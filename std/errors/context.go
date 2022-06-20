package errors

// TODO consider dropping in favor of 'github.com/pkg/errors', although unfortunate that it is not part of stdlib

// Context wraps an error to provide additional context information.
func Context(cause error, message string) error {
	// TODO should we return the pointer instead?
	return context{cause: cause, message: message}
}

type context struct {
	cause   error
	message string
}

func (c context) Error() string {
	return c.message + ": " + c.cause.Error()
}

func (c context) Unwrap() error {
	return c.cause
}
