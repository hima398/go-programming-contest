package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/emirpasic/gods/trees/redblacktree"
)

var sc = bufio.NewScanner(os.Stdin)

func solve(n int, s string) int {
	var ir []int
	for i, c := range s {
		if c == 'R' {
			ir = append(ir, i)
		}
	}
	//fmt.Println(ir)
	tree := redblacktree.NewWithIntComparator()
	incrementTree := func(k int, t *redblacktree.Tree) {
		v, found := t.Get(k)
		if found {
			t.Put(k, v.(int)+1)
		} else {
			t.Put(k, 1)
		}
	}

	decrementTree := func(k int, t *redblacktree.Tree) {
		v, found := t.Get(k)
		if found {
			if v.(int) > 1 {
				t.Put(k, v.(int)-1)
			} else if v == 1 {
				t.Remove(k)
			}
		}
	}
	for _, idx := range ir {
		d := 0
		l, r := idx-1, idx+1
		for l >= 0 && s[l] == 'A' && r < n && s[r] == 'C' {
			d++
			l--
			r++
		}
		if d > 0 {
			incrementTree(d, tree)
		}
	}
	//fmt.Println(q)
	ans := 0
	isOdd := true

	for tree.Size() > 0 {
		if isOdd {
			//奇数回目の操作はARC->R
			node := tree.Right()
			key := node.Key.(int)
			if key > 1 {
				decrementTree(key, tree)
				incrementTree(key-1, tree)
			} else {
				decrementTree(key, tree)
			}
		} else {
			//偶数回目の操作はARC->AC
			node := tree.Left()
			key := node.Key.(int)
			decrementTree(key, tree)
		}
		isOdd = !isOdd
		ans++
	}
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	s := nextString()

	ans := solve(n, s)
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}
