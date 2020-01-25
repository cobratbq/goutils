package io

import (
	"io"
	"log"

	"github.com/cobratbq/goutils/std/errors"
)

// CloseDiscarded closes the closer and discards any possible error.
func CloseDiscarded(c io.Closer) {
	c.Close()
}

// ClosePanicked closes the closer and panics with specified message in case of
// an error.
func ClosePanicked(c io.Closer, message string) {
	errors.RequireSuccess(c.Close(), message)
}

// CloseLogged closes the closer and logs specified message in case of error.
func CloseLogged(c io.Closer, message string) {
	if err := c.Close(); err != nil {
		log.Printf(message, err)
	}
}
