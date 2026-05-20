// Package amrdec implements AMR-WB decoding. Implementation is based on opencore-amr.
package amrdec

/*
#include "dec_if.h"
*/
import "C"

import "unsafe"

// New creates a new Decoder. Caller must call Close to avoid resource leak.
func New() *Decoder {
	return &Decoder{
		p: C.D_IF_init(),
	}
}

// Decoder for AMR-WB audio codec.
type Decoder struct {
	p unsafe.Pointer
}

// Close the decoder and free up resources.
func (d *Decoder) Close() {
	if d.p != nil {
		C.D_IF_exit(d.p)
		d.p = nil
	}
}

// Decode AMR-WB block into PCM16 audio.
func (d *Decoder) Decode(audio *[320]int16, block *[61]byte, noData bool) {
	bfi := 0
	if noData {
		bfi = 1
	}
	C.D_IF_decode(d.p, (*C.uchar)(&block[0]), (*C.short)(&audio[0]), C.int(bfi))
}
