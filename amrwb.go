// Package amrwb implements AMR-WB audio codec encoding and decoding.
package amrwb

const (
	PCMFrameSize = 320 // PCM frame size at 16kHz
	FrameSizeMax = 61  // max encoded AMR-WB block size
)

type PCMFrame = [PCMFrameSize]int16

var blockSizes = []byte{18, 24, 33, 37, 41, 47, 51, 59, 61, 6, 6, 0, 0, 0, 1, 1}

// BlockSize returns size of the first AMR-WB block in data.
func BlockSize(data []byte) int {
	if len(data) == 0 {
		return 0
	}
	mode := int((data[0] >> 3) & 0x0f)
	if mode >= len(blockSizes) {
		return -1
	}
	n := int(blockSizes[mode])
	if n > len(data) {
		return -1
	}
	return n
}

// Blocks counts the number of AMR-WB blocks in the frame.
// It also returns the size of all valid blocks.
func Blocks(data []byte) (sz, cnt int) {
	for len(data) > 0 {
		n := BlockSize(data)
		if n <= 0 {
			break
		}
		sz += n
		cnt++
		data = data[sz:]
	}
	return
}
