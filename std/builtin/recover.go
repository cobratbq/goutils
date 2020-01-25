package builtin

import (
	"log"
)

// RecoverLogged recovers from panic and logs the payload.
func RecoverLogged(message string) {
	if m := recover(); m != nil {
		log.Printf(message, m)
	}
}
