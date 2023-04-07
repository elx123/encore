// Code generated by "stringer -type=BuiltinKind -output=types_string.go"; DO NOT EDIT.

package schema

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Invalid-0]
	_ = x[Any-1]
	_ = x[Bool-2]
	_ = x[Int-3]
	_ = x[Int8-4]
	_ = x[Int16-5]
	_ = x[Int32-6]
	_ = x[Int64-7]
	_ = x[Uint-8]
	_ = x[Uint8-9]
	_ = x[Uint16-10]
	_ = x[Uint32-11]
	_ = x[Uint64-12]
	_ = x[Float32-13]
	_ = x[Float64-14]
	_ = x[String-15]
	_ = x[Bytes-16]
	_ = x[Time-17]
	_ = x[UUID-18]
	_ = x[JSON-19]
	_ = x[UserID-20]
	_ = x[Error-21]
	_ = x[unsupported - -1]
}

const _BuiltinKind_name = "unsupportedInvalidAnyBoolIntInt8Int16Int32Int64UintUint8Uint16Uint32Uint64Float32Float64StringBytesTimeUUIDJSONUserIDError"

var _BuiltinKind_index = [...]uint8{0, 11, 18, 21, 25, 28, 32, 37, 42, 47, 51, 56, 62, 68, 74, 81, 88, 94, 99, 103, 107, 111, 117, 122}

func (i BuiltinKind) String() string {
	i -= -1
	if i < 0 || i >= BuiltinKind(len(_BuiltinKind_index)-1) {
		return "BuiltinKind(" + strconv.FormatInt(int64(i+-1), 10) + ")"
	}
	return _BuiltinKind_name[_BuiltinKind_index[i]:_BuiltinKind_index[i+1]]
}