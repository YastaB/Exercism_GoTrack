package binarysearchtree

type SearchTreeData struct {
	left  *SearchTreeData
	data  int
	right *SearchTreeData
}

func Bst(data int) *SearchTreeData {
	return &SearchTreeData{
		left:  nil,
		data:  data,
		right: nil,
	}
}
func (root *SearchTreeData) Insert(data int) {
	cur := root
	for {
		if cur.data >= data {
			if cur.left == nil {
				cur.left = &SearchTreeData{
					left:  nil,
					data:  data,
					right: nil,
				}
				break
			}
			cur = cur.left
		} else{
			if cur.right == nil {
				cur.right = &SearchTreeData{
					left:  nil,
					data:  data,
					right: nil,
				}
				break
			}
			cur = cur.right
		}
	}


}
func (root *SearchTreeData) MapString(fun func(int) string) []string {
	return nil

}
func (root *SearchTreeData) MapInt(func(int) int) []int {
	return nil
}
