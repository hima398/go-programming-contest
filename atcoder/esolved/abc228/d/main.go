package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/psampaz/gods/trees/avltree"
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
	PrintVertically(ans)
}

func solve(q int, t, x []int) []int {
	n := 1 << 20

	tree := avltree.NewWithIntComparator()
	//tree := redblacktree.NewWithIntComparator()
	m := make(map[int]int)
	for i := 0; i < n; i++ {
		tree.Put(i, -1)
	}
	var ans []int
	for i := 0; i < q; i++ {
		switch t[i] {
		case 1:
			node, found := tree.Ceiling(x[i] % n)
			//TODO:右端の処理
			if !found {
				node, _ = tree.Ceiling(0)
			}
			m[node.Key.(int)] = x[i]
			tree.Remove(node.Key)
		case 2:
			idx := x[i] % n
			_, found := tree.Get(idx)
			if found {
				ans = append(ans, -1)
			} else {
				ans = append(ans, m[idx])
			}
		}
	}
	//fmt.Println(m)
	return ans
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func PrintVertically(x []int) {
	defer out.Flush()
	for _, v := range x {
		fmt.Fprintln(out, v)
	}
}
