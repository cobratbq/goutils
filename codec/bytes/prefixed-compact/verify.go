// SPDX-License-Identifier: LGPL-3.0-only

// TODO at some point implement read-functions for reading from io.Reader
package prefixed

import (
	"io"

	"github.com/cobratbq/goutils/assert"
	"github.com/cobratbq/goutils/std/errors"
	io_ "github.com/cobratbq/goutils/std/io"
	"github.com/cobratbq/goutils/std/log"
)

func verifyKeyValue(in io.Reader, hdr *Header) error {
	assert.Equal(TYPE_KEYVALUE, hdr.Vtype)
	var err error
	for {
		if _, err = io_.DiscardN(in, int64(hdr.Size)); err != nil {
			return err
		}
		if hdr.Terminated {
			break
		}
		if *hdr, err = ReadHeader(in); err != nil {
			return err
		} else if hdr.Vtype != TYPE_KEYVALUE {
			return errors.ErrIllegal
		}
	}
	return VerifyValue(in)
}

// TODO not tested yet
func VerifyValue(in io.Reader) error {
	log.Traceln("Reading value-header")
	var hdr Header
	var err error
	for terminated := false; !terminated; terminated = hdr.Terminated {
		// FIXME verify same valuetype|multiplicity upon continuation
		if hdr, err = ReadHeader(in); err != nil {
			return err
		}
		switch hdr.Vtype {
		case TYPE_BYTES:
			if _, err = io_.DiscardN(in, int64(hdr.Size)); err != nil {
				return err
			}
		case TYPE_KEYVALUE:
			return verifyKeyValue(in, &hdr)
		case TYPE_SEQUENCE:
			for i := uint16(0); i < hdr.Size; i++ {
				if err = VerifyValue(in); err != nil {
					return err
				}
			}
		case TYPE_MAP:
			for i := uint16(0); i < hdr.Size; i++ {
				if hdr, err = ReadHeader(in); err != nil {
					return err
				} else if hdr.Vtype != TYPE_KEYVALUE {
					return errors.ErrIllegal
				}
				if err = verifyKeyValue(in, &hdr); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

// TODO not tested yet
func Verify(in io.Reader) error {
	var err error
	for {
		err = VerifyValue(in)
		if errors.Is(err, io.EOF) {
			return nil
		} else if err != nil {
			return err
		}
	}
}
