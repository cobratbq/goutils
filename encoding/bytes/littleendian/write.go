// SPDX-License-Identifier: LGPL-3.0-or-later

package littleendian

import "io"

func WriteUint8(out io.Writer, value uint8) error {
	_, err := out.Write([]byte{value})
	return err
}

func WriteUint16(out io.Writer, value uint16) error {
	b0 := uint8(0x00ff & value)
	b1 := uint8((0xff00 & value) >> 8)
	_, err := out.Write([]byte{b0, b1})
	return err
}

func WriteUint32(out io.Writer, value uint32) error {
	b0 := uint8(0x000000ff & value)
	b1 := uint8((0x0000ff00 & value) >> 8)
	b2 := uint8((0x00ff0000 & value) >> 16)
	b3 := uint8((0xff000000 & value) >> 24)
	_, err := out.Write([]byte{b0, b1, b2, b3})
	return err
}

func WriteUint64(out io.Writer, value uint64) error {
	b0 := uint8(0x00000000000000ff & value)
	b1 := uint8((0x000000000000ff00 & value) >> 8)
	b2 := uint8((0x0000000000ff0000 & value) >> 16)
	b3 := uint8((0x00000000ff000000 & value) >> 24)
	b4 := uint8((0x000000ff00000000 & value) >> 32)
	b5 := uint8((0x0000ff0000000000 & value) >> 40)
	b6 := uint8((0x00ff000000000000 & value) >> 48)
	b7 := uint8((0xff00000000000000 & value) >> 56)
	_, err := out.Write([]byte{b0, b1, b2, b3, b4, b5, b6, b7})
	return err
}
