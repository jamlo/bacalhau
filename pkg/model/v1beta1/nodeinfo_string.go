// Code generated by "stringer -type=NodeType --trimprefix=NodeType --output nodeinfo_string.go"; DO NOT EDIT.

package v1beta1

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[nodeTypeUnknown-0]
	_ = x[NodeTypeCompute-1]
}

const _NodeType_name = "nodeTypeUnknownCompute"

var _NodeType_index = [...]uint8{0, 15, 22}

func (i NodeType) String() string {
	if i < 0 || i >= NodeType(len(_NodeType_index)-1) {
		return "NodeType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _NodeType_name[_NodeType_index[i]:_NodeType_index[i+1]]
}