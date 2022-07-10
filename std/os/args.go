// SPDX-License-Identifier: LGPL-3.0-or-later
package os

import (
	"os"
	"strings"
)

// MatchProcessCommandArg checks if the executing commandline argument is same as the specified
// command.
func MatchProcessCommandArg(command string) bool {
	return os.Args[0] == command || strings.HasSuffix(os.Args[0], string(os.PathSeparator)+command)
}
