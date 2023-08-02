// SPDX-License-Identifier: LGPL-3.0-only

package os

const ColorReset = "\033[0m"

const ColorBlack = "\033[30m"
const ColorRed = "\033[31m"
const ColorGreen = "\033[32m"
const ColorYellow = "\033[33m"
const ColorBlue = "\033[34m"
const ColorMagenta = "\033[35m"
const ColorCyan = "\033[36m"
const ColorWhite = "\033[37m"

const ColorBlackBold = "\033[1;30m"
const ColorRedBold = "\033[1;31m"
const ColorGreenBold = "\033[1;32m"
const ColorYellowBold = "\033[1;33m"
const ColorBlueBold = "\033[1;34m"
const ColorMagentaBold = "\033[1;35m"
const ColorCyanBold = "\033[1;36m"
const ColorWhiteBold = "\033[1;37m"

const ColorBlackBackground = "\033[40m"
const ColorRedBackground = "\033[41m"
const ColorGreenBackground = "\033[42m"
const ColorYellowBackground = "\033[43m"
const ColorBlueBackground = "\033[44m"
const ColorMagentaBackground = "\033[45m"
const ColorCyanBackground = "\033[46m"
const ColorWhiteBackground = "\033[47m"

// ColorText prepends a terminal escape code for text color, followed by the provided text, and
// ending with a color reset code. Use the color-constants defined in this package for the
// `colorcode` parameter.
func ColorText(colorcode string, text string) string {
	return colorcode + text + ColorReset
}

// TODO consider if we want to add this, but below shows how escape codes are composed, so not removing yet.
//func ColorCompose(fg, bg byte, bold bool) string {
//	line := "\033["
//	if bold {
//		line += "1;"
//	}
//	line += "3" + string(fg) + ";"
//	line += "4" + string(bg) + "m"
//	return line
//}
