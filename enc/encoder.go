// Package amrenc implements AMR-WB encoding. Implementation is based on vo-amrwbenc.
package amrenc

/*
#include "enc_if.h"
*/
import "C"

import (
	"unsafe"
)

// Mode specifies the bit rate the codec supports.
type Mode byte

const (
	Mode6Kb  = Mode(0) // 6.60kbps
	Mode8Kb  = Mode(1) // 8.85kbps
	Mode12Kb = Mode(2) // 12.65kbps
	Mode14Kb = Mode(3) // 14.25kbps
	Mode16Kb = Mode(4) // 15.85bps
	Mode18Kb = Mode(5) // 18.25bps
	Mode20Kb = Mode(6) // 19.85kbps
	Mode23Kb = Mode(7) // 23.05kbps
	Mode24Kb = Mode(8) // 23.85kbps
)

// New creates a new Encoder. Caller must call Close to avoid resource leak.
func New(mode Mode) *Encoder {
	return &Encoder{
		p:    C.E_IF_init(),
		mode: mode,
	}
}

// Encoder for AMR-WB audio codec.
type Encoder struct {
	p    unsafe.Pointer
	mode Mode
	dtx  int8
}

// Close the encoder and free up resources.
func (e *Encoder) Close() {
	if e.p != nil {
		C.E_IF_exit(e.p)
		e.p = nil
	}
}

// SetMode sets the encoding mode.
func (e *Encoder) SetMode(mode Mode) {
	e.mode = mode
}

// SetDTX sets the DTX.
func (e *Encoder) SetDTX(enabled bool) {
	if enabled {
		e.dtx = 1
	} else {
		e.dtx = 0
	}
}

// Encode PCM16 audio frame into AMR-WB block, returning the number of bytes written.
func (e *Encoder) Encode(block *[61]byte, audio *[320]int16) int {
	return int(C.E_IF_encode(e.p, C.int(e.mode), (*C.short)(&audio[0]), (*C.uchar)(&block[0]), C.int(e.dtx)))
}
