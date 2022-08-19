package mid

import (
	"encoding/json"
	"go-practice/main/common"
)

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *common.TreeNode) string {
	return ""
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *common.TreeNode {
	return nil
}

func slice2Str(slice []int) string {
	bytes, ok := json.Marshal(slice)
	if ok != nil {
		return ""
	} else {
		return string(bytes)
	}
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
