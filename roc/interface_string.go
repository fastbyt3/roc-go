// Code generated by "stringer -type Interface -trimprefix Interface -output interface_string.go"; DO NOT EDIT.

package roc

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[InterfaceConsolidated-1]
	_ = x[InterfaceAudioSource-11]
	_ = x[InterfaceAudioRepair-12]
	_ = x[InterfaceAudioControl-13]
}

const (
	_Interface_name_0 = "Consolidated"
	_Interface_name_1 = "AudioSourceAudioRepairAudioControl"
)

var (
	_Interface_index_1 = [...]uint8{0, 11, 22, 34}
)

func (i Interface) String() string {
	switch {
	case i == 1:
		return _Interface_name_0
	case 11 <= i && i <= 13:
		i -= 11
		return _Interface_name_1[_Interface_index_1[i]:_Interface_index_1[i+1]]
	default:
		return "Interface(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
