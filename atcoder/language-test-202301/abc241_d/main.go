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

func solve(q int, t, x, k []int) []int {

	increment := func(x int, t *redblacktree.Tree) {
		v, found := t.Get(x)
		if found {
			t.Put(x, v.(int)+1)
		} else {
			t.Put(x, 1)
		}
	}
	redBlackTree := redblacktree.NewWithIntComparator()
	var ans []int
	for i := 0; i < q; i++ {
		switch t[i] {
		case 1:
			increment(x[i], redBlackTree)
		case 2:
			node, found := redBlackTree.Floor(x[i])
			if !found {
				ans = append(ans, -1)
				continue
			}
			it := redBlackTree.IteratorAt(node)
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
			node, found := redBlackTree.Ceiling(x[i])
			if !found {
				ans = append(ans, -1)
				continue
			}
			it := redBlackTree.IteratorAt(node)
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

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	q := nextInt()
	var t, x, k []int
	for i := 0; i < q; i++ {
		t = append(t, nextInt())
		x = append(x, nextInt())
		if t[i] == 1 {
			k = append(k, -1)
		} else {
			k = append(k, nextInt())
		}
	}
	ans := solve(q, t, x, k)
	for _, v := range ans {
		print(v)
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func print(x any) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
