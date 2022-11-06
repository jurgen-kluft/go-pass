// Copyright 2011 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package coding

import (
	"bytes"
	"testing"

	"github.com/jurgen-kluft/go-pass/qrterminal/qrcode/gf256"
)

func TestEncode(t *testing.T) {
	data := []byte{0x10, 0x20, 0x0c, 0x56, 0x61, 0x80, 0xec, 0x11, 0xec, 0x11, 0xec, 0x11, 0xec, 0x11, 0xec, 0x11}
	check := []byte{0xa5, 0x24, 0xd4, 0xc1, 0xed, 0x36, 0xc7, 0x87, 0x2c, 0x55}
	rs := gf256.NewRSEncoder(Field, len(check))
	out := make([]byte, len(check))
	rs.ECC(data, out)
	if !bytes.Equal(out, check) {
		t.Errorf("have %x want %x", out, check)
	}
}
