// Code generated by "stringer -type FieldType"; DO NOT EDIT.

package cursors

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Float-0]
	_ = x[Integer-1]
	_ = x[Unsigned-2]
	_ = x[String-3]
	_ = x[Boolean-4]
	_ = x[Undefined-5]
}

const _FieldType_name = "FloatIntegerUnsignedStringBooleanUndefined"

var _FieldType_index = [...]uint8{0, 5, 12, 20, 26, 33, 42}

func (i FieldType) String() string {
	if i < 0 || i >= FieldType(len(_FieldType_index)-1) {
		return "FieldType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _FieldType_name[_FieldType_index[i]:_FieldType_index[i+1]]
}
