package leet_code_attempts

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
	//
}

func Constructor() Codec {
	return Codec{}
}

func (c *Codec) serialize(root *TreeNode) string {
	return ""
}

func (c *Codec) deserialize(data string) *TreeNode {
	return nil
}
