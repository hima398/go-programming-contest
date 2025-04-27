package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"

	"github.com/liyue201/gostl/ds/priorityqueue"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	s := nextIntSlice(n)
	t := nextIntSlice(n)

	ans := solve(n, s, t)

	for _, v := range ans {
		Print(v)
	}
}

func solve(n int, s, t []int) []int {
	const INF = math.MaxInt
	type node struct {
		t, i int
	}
	q := priorityqueue.New[node](func(a, b node) int {
		if a.t == b.t {
			return 0
		}
		if a.t < b.t {
			return -1
		}
		return 1
	})
	for i := 0; i < n; i++ {
		q.Push(node{t[i], i})
	}
	ans := make([]int, n)
	for i := range ans {
		ans[i] = INF
	}
	rem := n
	for rem > 0 {
		cur := q.Pop()
		if ans[cur.i] == INF {
			ans[cur.i] = cur.t
			rem--
		}
		q.Push(node{cur.t + s[cur.i], (cur.i + 1) % n})
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
