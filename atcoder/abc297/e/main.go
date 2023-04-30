package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"

	"github.com/emirpasic/gods/trees/avltree"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, k := nextInt(), nextInt()
	a := nextIntSlice(n)

	ans := solve(n, k, a)

	PrintInt(ans)
}

func solve(n, k int, a []int) int {
	sort.Ints(a)
	q := avltree.NewWithIntComparator() //&PriorityQueue{}
	q.Put(0, 0)
	//ans := avltree.NewWithIntComparator()
	var ans []int

	for len(ans) <= k {
		cur := q.Left().Key.(int)
		q.Remove(q.Left().Key.(int))
		//ans.Put(cur, 0)
		ans = append(ans, cur)
		for j := 0; j < n; j++ {
			v := cur + a[j]
			q.Put(v, 0)
		}
		//fmt.Println("ans.Size(), q.Size() = ", ans.Size(), q.Size())
	}
	return ans[k]
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

func PrintInt(x int) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
