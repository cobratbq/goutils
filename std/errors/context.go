package errors

// Context wraps an error to provide additional context information.
// TODO consider defining different variations of context: with-message, with-key-values-pairs, ... We can consider this basic context type as key-value pair with key 'message'. Then as we extract key-value pairs, we can include base contexts.
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
