// SPDX-License-Identifier: AGPL-3.0-or-later

package bytes

import "io"

func ReadByte(in io.Reader, buffer []byte) error {
	_, err := io.ReadFull(in, buffer[:1:1])
	return err
}

func ReadBytes(in io.Reader, buffer []byte) error {
	_, err := io.ReadFull(in, buffer)
	return err
}
