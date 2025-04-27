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
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	q := nextInt()
	var t, x []int
	for i := 0; i < q; i++ {
		t = append(t, nextInt())
		x = append(x, nextInt())
	}
	ans := solve(q, t, x)
	for _, v := range ans {
		Print(v)
	}
}

func solve(q int, t, x []int) []int {
	const n = 1 << 20

	tree := redblacktree.NewWithIntComparator()
	a := make([]int, n)
	for i := 0; i < n; i++ {
		tree.Put(i, -1)
		a[i] = -1
	}

	var ans []int
	for i := 0; i < q; i++ {
		switch t[i] {
		case 1:
			var node *redblacktree.Node
			if v, found := tree.Ceiling(x[i] % n); !found {
				node, _ = tree.Ceiling(0)
			} else {
				node = v
			}
			a[node.Key.(int)] = x[i]
			tree.Remove(node.Key)
		case 2:
			idx := x[i] % n
			if _, found := tree.Get(idx); found {
				ans = append(ans, -1)
			} else {
				ans = append(ans, a[idx])
			}
		}
	}
	return ans
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func Print(x any) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
