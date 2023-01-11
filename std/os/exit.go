// SPDX-License-Identifier: AGPL-3.0-or-later

package os

import "os"

// ExitWithError exits the program with specified errorcode, and prints the
// provided message to Stderr including a line-ending.
func ExitWithError(errcode int, message string) {
	os.Stderr.WriteString(message + "\n")
	os.Exit(errcode)
}
