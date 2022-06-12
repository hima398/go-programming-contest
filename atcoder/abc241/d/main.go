package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func solve(q int, t, x, k []int) (ans []int) {
	//rbTree := redblacktree.NewWithIntComparator()
	rbTree := NewRedBlackTree()

	for i := 0; i < q; i++ {
		switch t[i] {
		case 1:
			if v, found := rbTree.Get(x[i]); found {
				rbTree.Put(x[i], v+1)
			} else {
				rbTree.Put(x[i], 1)
			}
			//rbTree: RedBlackTree
			//│   ┌── 30
			//└── 20
			//    └── 10
			//        └── 1
			//rbTree: RedBlackTree
			//│   ┌── key:30, value:1
			//└── key:20, value:2
			//    └── key:10, value:1
			//        └── key:1, value:1
			fmt.Printf("rbTree: %v\n", rbTree)
		case 2:
			node, _ := rbTree.Floor(x[i])
			it := rbTree.IteratorAt(node)
			s := it.Value()
			found := s >= k[i]
			fmt.Println(it)
			for !found && it.Prev() {
				s += it.Value()
				found = s >= k[i]
			}
			if found {
				ans = append(ans, it.Key())
			} else {
				ans = append(ans, -1)
			}
		case 3:
			node, _ := rbTree.Ceil(x[i])
			it := rbTree.IteratorAt(node)
			s := it.Value()
			found := s >= k[i]
			for !found && it.Next() {
				s += it.Value()
				found = s >= k[i]
			}
			if found {
				ans = append(ans, it.Key())
			} else {
				ans = append(ans, -1)
			}
		}
	}

	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	q := nextInt()
	t := make([]int, q)
	x := make([]int, q)
	k := make([]int, q)
	for i := 0; i < q; i++ {
		t[i] = nextInt()
		if t[i] == 1 {
			x[i] = nextInt()
		} else {
			x[i], k[i] = nextInt(), nextInt()
		}
	}
	ans := solve(q, t, x, k)
	PrintVertically(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextIntSlice(n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = nextInt()
	}
	return s
}

func nextFloat64() float64 {
	sc.Scan()
	f, _ := strconv.ParseFloat(sc.Text(), 64)
	return f
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func PrintInt(x int) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}

func PrintFloat64(x float64) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}

func PrintString(x string) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}

func PrintHorizonaly(x []int) {
	defer out.Flush()
	fmt.Fprintf(out, "%d", x[0])
	for i := 1; i < len(x); i++ {
		fmt.Fprintf(out, " %d", x[i])
	}
	fmt.Fprintln(out)
}

func PrintVertically(x []int) {
	defer out.Flush()
	for _, v := range x {
		fmt.Fprintln(out, v)
	}
}

const (
	red   = 0
	black = 1
)

//RedBlackTree
type RedBlackTreeNode struct {
	Key    int
	Value  int
	color  byte
	Left   *RedBlackTreeNode
	Right  *RedBlackTreeNode
	Parent *RedBlackTreeNode
}

func (node *RedBlackTreeNode) String() string {
	return fmt.Sprintf("key:%v, value:%v", node.Key, node.Value)
}

func (node *RedBlackTreeNode) sibling() *RedBlackTreeNode {
	if node == nil || node.Parent == nil {
		return nil
	}
	if node == node.Parent.Left {
		return node.Parent.Right
	}
	return node.Parent.Left
}

func (node *RedBlackTreeNode) parentSibling() *RedBlackTreeNode {
	if node == nil || node.Parent == nil || node.Parent.Parent == nil {
		return nil
	}
	return node.Parent.sibling()
}

func (node *RedBlackTreeNode) grandparent() *RedBlackTreeNode {
	if node == nil || node.Parent == nil {
		return nil
	}
	return node.Parent.Parent
}

const (
	begin   = 0
	between = 1
	end     = 2
)

type RedBlackTreeIterator struct {
	tree     *RedBlackTree
	node     *RedBlackTreeNode
	position byte
}

func (it *RedBlackTreeIterator) Key() int {
	return it.node.Key
}

func (it *RedBlackTreeIterator) Value() int {
	return it.node.Value
}

func (it *RedBlackTreeIterator) Prev() bool {
	if it.position == begin {
		it.node = nil
		it.position = begin
		fmt.Println("it.position == begin")
		return false
	}
	if it.position == end {
		right := it.tree.Right()
		if right == nil {
			it.node = nil
			it.position = begin
			fmt.Println("it.position == end, right==nil")
			return false
		}
		it.node = right
		it.position = between
		return true
	}
	if it.node.Left != nil {
		it.node = it.node.Left
		for it.node.Right != nil {
			it.node = it.node.Right
		}
		it.position = between
		return true
	}
	if it.node.Parent != nil {
		node := it.node
		for it.node.Parent != nil {
			it.node = it.node.Parent
			if node.Key >= it.node.Key {
				it.position = between
				return true
			}
		}
	}
	fmt.Printf("it.tree: %v\n", it.tree)
	fmt.Println("end")

	it.node = nil
	it.position = begin
	return false
}

func (iterator *RedBlackTreeIterator) Prev2() bool {
	if iterator.position == begin {
		goto begin
	}
	if iterator.position == end {
		right := iterator.tree.Right()
		if right == nil {
			goto begin
		}
		iterator.node = right
		goto between
	}
	if iterator.node.Left != nil {
		iterator.node = iterator.node.Left
		for iterator.node.Right != nil {
			iterator.node = iterator.node.Right
		}
		goto between
	}
	if iterator.node.Parent != nil {
		node := iterator.node
		for iterator.node.Parent != nil {
			iterator.node = iterator.node.Parent
			if node.Key >= iterator.node.Key {
				goto between
			}
		}
	}

begin:
	iterator.node = nil
	iterator.position = begin
	return false

between:
	iterator.position = between
	return true

}

func (it *RedBlackTreeIterator) Next() bool {
	if it.position == end {
		it.node = nil
		it.position = end
		return false
	}
	if it.position == begin {
		left := it.tree.Left()
		if left == nil {
			it.node = nil
			it.position = end
			return false
		}
		it.node = left
		it.position = between
		return true
	}
	if it.node.Right != nil {
		it.node = it.node.Right
		for it.node.Left != nil {
			it.node = it.node.Left
		}
		it.position = between
		return true

	}
	if it.node.Parent != nil {
		node := it.node
		for it.node.Parent != nil {
			it.node = it.node.Parent
			if node.Key <= it.node.Key {
				it.position = between
				return true
			}
		}
	}
	it.node = nil
	it.position = end
	return false
}

func (iterator *RedBlackTreeIterator) Next2() bool {
	if iterator.position == end {
		goto end
	}
	if iterator.position == begin {
		left := iterator.tree.Left()
		if left == nil {
			goto end
		}
		iterator.node = left
		goto between
	}
	if iterator.node.Right != nil {
		iterator.node = iterator.node.Right
		for iterator.node.Left != nil {
			iterator.node = iterator.node.Left
		}
		goto between
	}
	if iterator.node.Parent != nil {
		node := iterator.node
		for iterator.node.Parent != nil {
			iterator.node = iterator.node.Parent
			if node.Key <= iterator.node.Key {
				goto between
			}
		}
	}

end:
	iterator.node = nil
	iterator.position = end
	return false

between:
	iterator.position = between
	return true

}

type RedBlackTree struct {
	Root *RedBlackTreeNode
	size int
}

func NewRedBlackTree() *RedBlackTree {
	return &RedBlackTree{}
}

func (tree *RedBlackTree) Size() int {
	return tree.size
}

func (tree *RedBlackTree) Get(key int) (int, bool) {
	node := tree.lookup(key)
	if node == nil {
		return 0, false
	}
	return node.Value, true
}

func (tree *RedBlackTree) Left() *RedBlackTreeNode {
	var parent *RedBlackTreeNode
	current := tree.Root
	for current != nil {
		parent = current
		current = current.Left
	}
	return parent
}

func (tree *RedBlackTree) Right() *RedBlackTreeNode {
	var parent *RedBlackTreeNode
	current := tree.Root
	for current != nil {
		parent = current
		current = current.Right
	}
	return parent
}

func (tree *RedBlackTree) Floor(key int) (*RedBlackTreeNode, bool) {
	var floor *RedBlackTreeNode
	found := false
	node := tree.Root
	for node != nil {
		switch {
		case key == node.Key:
			return node, true
		case key < node.Key:
			node = node.Left
		case key > node.Key:
			floor, found = node, true
			node = node.Right
		}
	}
	if found {
		return floor, found
	} else {
		return nil, false
	}
}

func (tree *RedBlackTree) Ceil(key int) (*RedBlackTreeNode, bool) {
	var ceil *RedBlackTreeNode
	found := false
	node := tree.Root
	for node != nil {
		switch {
		case key == node.Key:
			return node, true
		case key < node.Key:
			ceil, found = node, true
			node = node.Left
		case key > node.Key:
			node = node.Right
		}
	}
	if found {
		return ceil, found
	} else {
		return nil, false
	}
}

func (tree *RedBlackTree) Put(key, value int) {
	//空の木に新規挿入
	if tree.Root == nil {
		tree.Root = &RedBlackTreeNode{Key: key, Value: value, color: red}
		tree.insert(tree.Root)
		tree.size = 1
		return
	}

	//木を辿って更新
	node := tree.Root
	loop := true
	var addNode *RedBlackTreeNode
	for loop {
		switch {
		case key == node.Key:
			node.Key = key
			node.Value = value
			return
		case key < node.Key:
			if node.Left == nil {
				node.Left = &RedBlackTreeNode{Key: key, Value: value, color: red}
				addNode = node.Left
				loop = false
			} else {
				node = node.Left
			}
		case key > node.Key:
			if node.Right == nil {
				node.Right = &RedBlackTreeNode{Key: key, Value: value, color: red}
				addNode = node.Right
				loop = false
			} else {
				node = node.Right
			}
		}
	}
	tree.insert(addNode)
	tree.size++
}

func (tree *RedBlackTree) Iterator() RedBlackTreeIterator {
	return RedBlackTreeIterator{tree, nil, begin}
}

func (tree *RedBlackTree) IteratorAt(node *RedBlackTreeNode) RedBlackTreeIterator {
	return RedBlackTreeIterator{tree, node, between}
}

func (tree *RedBlackTree) String() string {
	str := "RedBlackTree\n"
	if tree.size > 0 {
		output(tree.Root, "", true, &str)
	}
	return str
}

func (tree *RedBlackTree) lookup(key int) *RedBlackTreeNode {
	current := tree.Root
	for current != nil {
		switch {
		case key == current.Key:
			return current
		case key < current.Key:
			current = current.Left
		case key > current.Key:
			current = current.Right
		}
	}
	return nil
}

func (tree *RedBlackTree) replaceNode(old, new *RedBlackTreeNode) {
	if old.Parent == nil {
		tree.Root = new
	} else {
		if old == old.Parent.Left {
			old.Parent.Left = new
		} else {
			old.Parent.Right = new
		}
	}
	if new != nil {
		new.Parent = old.Parent
	}
}

func (tree *RedBlackTree) rotateLeft(node *RedBlackTreeNode) {
	right := node.Right
	tree.replaceNode(node, right)
	node.Right = right.Left
	if right.Left != nil {
		right.Left.Parent = node
	}
	right.Left = node
	node.Parent = right
}

func (tree *RedBlackTree) rotateRight(node *RedBlackTreeNode) {
	left := node.Left
	tree.replaceNode(node, left)
	node.Left = left.Right
	if left.Right != nil {
		left.Right.Parent = node
	}
	left.Right = node
	node.Parent = left
}

func (tree *RedBlackTree) insert(node *RedBlackTreeNode) {
	if node.Parent == nil {
		node.color = black
		return
	}
	if tree.getColor(node.Parent) == black {
		return
	}
	parentSibling := node.parentSibling()
	if tree.getColor(parentSibling) == red {
		node.Parent.color = black
		parentSibling.color = black
		//TODO:node.grandparent()がnilでないか確認
		node.grandparent().color = red
		tree.insert(node.grandparent())
		return
	}
	grandparent := node.grandparent()
	if node == node.Parent.Right && node.Parent == grandparent.Left {
		tree.rotateLeft(node.Parent)
		node = node.Left
	} else if node == node.Parent.Left && node.Parent == grandparent.Right {
		tree.rotateRight(node.Parent)
		node = node.Right
	}
	node.Parent.color = black
	grandparent.color = red
	if node == node.Parent.Left && node.Parent == grandparent.Left {
		tree.rotateRight(grandparent)
	} else if node == node.Parent.Right && node.Parent == grandparent.Right {
		tree.rotateLeft(grandparent)
	}
}

func (tree *RedBlackTree) getColor(node *RedBlackTreeNode) byte {
	if node == nil {
		return black
	}
	return node.color
}

func output(node *RedBlackTreeNode, prefix string, isTail bool, str *string) {
	if node.Right != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "│   "
		} else {
			newPrefix += "    "
		}
		output(node.Right, newPrefix, false, str)
	}
	*str += prefix
	if isTail {
		*str += "└── "
	} else {
		*str += "┌── "
	}
	*str += node.String() + "\n"
	if node.Left != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "    "
		} else {
			newPrefix += "│   "
		}
		output(node.Left, newPrefix, true, str)
	}
}
