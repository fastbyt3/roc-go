// Code generated by "stringer -type FrameEncoding -trimprefix FrameEncoding -output frame_encoding_string.go"; DO NOT EDIT.

package roc

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[FrameEncodingPcmFloat-1]
}

const _FrameEncoding_name = "PcmFloat"

var _FrameEncoding_index = [...]uint8{0, 8}

func (i FrameEncoding) String() string {
	i -= 1
	if i < 0 || i >= FrameEncoding(len(_FrameEncoding_index)-1) {
		return "FrameEncoding(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _FrameEncoding_name[_FrameEncoding_index[i]:_FrameEncoding_index[i+1]]
}