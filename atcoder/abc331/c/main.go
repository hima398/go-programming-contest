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

	n := nextInt()
	a := nextIntSlice(n)

	ans := solve(n, a)

	PrintHorizonaly(ans)
}

func solve(n int, a []int) []int {
	rbTree := redblacktree.NewWithIntComparator()

	for _, ai := range a {
		if v, found := rbTree.Get(ai); found {
			rbTree.Put(ai, v.(int)+1)
		} else {
			rbTree.Put(ai, 1)
		}
	}

	it := rbTree.Iterator()
	it.End()
	m := make(map[int]int)
	var s int
	for it.Prev() {
		k, v := it.Key().(int), it.Value().(int)
		//fmt.Println(s, k, v)
		m[k] = s
		s += k * v
	}
	var ans []int
	for _, ai := range a {
		ans = append(ans, m[ai])
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

func PrintHorizonaly(x []int) {
	defer out.Flush()
	fmt.Fprintf(out, "%d", x[0])
	for i := 1; i < len(x); i++ {
		fmt.Fprintf(out, " %d", x[i])
	}
	fmt.Fprintln(out)
}
