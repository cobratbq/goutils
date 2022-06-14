package os

import (
	"os"
	"strings"
)

func MatchProcessCommandArg(command string) bool {
	return os.Args[0] == command || strings.HasSuffix(os.Args[0], string(os.PathSeparator)+command)
}
