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

	n, q := nextInt(), nextInt()
	a := nextIntSlice(n)
	var idx, x []int
	for i := 0; i < q; i++ {
		idx = append(idx, nextInt()-1)
		x = append(x, nextInt())
	}

	ans := solve(n, q, a, idx, x)

	for _, v := range ans {
		Print(v)
	}
}

func solve(n, q int, a, idx, x []int) []int {
	const max = 2 * int(1e5)
	m := make(map[int]int)
	for _, ai := range a {
		m[ai]++
	}
	tree := redblacktree.NewWithIntComparator()
	for k := 0; k <= max; k++ {
		if _, found := m[k]; !found {
			tree.Put(k, 0)
		}
	}
	var ans []int
	for i := 0; i < q; i++ {
		m[a[idx[i]]]--
		if a[idx[i]] <= max && m[a[idx[i]]] == 0 {
			tree.Put(a[idx[i]], 0)
		}
		a[idx[i]] = x[i]
		m[a[idx[i]]]++
		if _, found := tree.Get(a[idx[i]]); found {
			tree.Remove(a[idx[i]])
		}

		ans = append(ans, tree.Left().Key.(int))
	}
	return ans
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func nextIntSlice(n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = nextInt()
	}
	return s
}

func Print(x any) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
