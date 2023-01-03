package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/emirpasic/gods/trees/redblacktree"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	q := nextInt()
	t, x, k := make([]int, q), make([]int, q), make([]int, q)
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

func solve(q int, t, x, k []int) []int {
	rbTree := redblacktree.NewWithIntComparator()

	size := 0
	var ans []int
	for i := 0; i < q; i++ {
		switch t[i] {
		case 1:
			v, found := rbTree.Get(x[i])
			if found {
				rbTree.Put(x[i], v.(int)+1)
			} else {
				rbTree.Put(x[i], 1)
			}
			size++
		case 2:
			node, found := rbTree.Floor(x[i])
			if !found {
				ans = append(ans, -1)
				continue
			}
			//it := rbTree.IteratorAt(node)
			it := Iterator{rbTree, node, between}
			s := it.Value().(int)
			found = s >= k[i]
			for !found && it.Prev() {
				s += it.Value().(int)
				found = s >= k[i]
			}
			if found {
				ans = append(ans, it.Key().(int))
			} else {
				ans = append(ans, -1)
			}
		case 3:
			node, found := rbTree.Ceiling(x[i])
			if !found {
				ans = append(ans, -1)
				continue
			}
			//it := rbTree.IteratorAt(node)
			it := Iterator{rbTree, node, between}
			s := it.Value().(int)
			found = s >= k[i]
			for !found && it.Next() {
				s += it.Value().(int)
				found = s >= k[i]
			}
			if found {
				ans = append(ans, it.Key().(int))
			} else {
				ans = append(ans, -1)
			}
		}
	}
	return ans
}

type position byte

const (
	begin, between, end position = 0, 1, 2
)

type Iterator struct {
	tree *redblacktree.Tree
	node *redblacktree.Node
	pos  position
}

func (it *Iterator) Prev() bool {
	//左端
	if it.pos == begin {
		goto begin
	}
	if it.pos == end {
		right := it.tree.Right()
		if right == nil {
			goto begin
		}
		it.node = right
		goto between
	}
	if it.node.Left != nil {
		it.node = it.node.Left
		for it.node.Right != nil {
			it.node = it.node.Right
		}
		goto between
	}
	if it.node.Parent != nil {
		node := it.node
		for it.node.Parent != nil {
			it.node = it.node.Parent
			if it.tree.Comparator(node.Key, it.node.Key) >= 0 {
				goto between
			}
		}
	}

begin:
	it.node = nil
	it.pos = begin
	return false

between:
	it.pos = between
	return true
}

func (it *Iterator) Next() bool {
	if it.pos == end {
		goto end
	}
	if it.pos == begin {
		left := it.tree.Left()
		if left == nil {
			goto end
		}
		it.node = left
		goto between
	}
	if it.node.Right != nil {
		it.node = it.node.Right
		for it.node.Left != nil {
			it.node = it.node.Left
		}
		goto between
	}
	if it.node.Parent != nil {
		node := it.node
		for it.node.Parent != nil {
			it.node = it.node.Parent
			if it.tree.Comparator(node.Key, it.node.Key) <= 0 {
				goto between
			}
		}
	}

end:
	it.node = nil
	it.pos = end
	return false

between:
	it.pos = between
	return true
}

func (it *Iterator) Key() interface{} {
	return it.node.Key
}

func (it *Iterator) Value() interface{} {
	return it.node.Value
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func PrintVertically(x []int) {
	defer out.Flush()
	for _, v := range x {
		fmt.Fprintln(out, v)
	}
}
