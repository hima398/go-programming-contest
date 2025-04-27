package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	n, m := nextInt(), nextInt()
	p := nextIntSlice(n)

	ans := solve(n, m, p)

	Print(ans)
}

func solve(n, m int, p []int) int {
	//予算Mでx以下の商品を買い占められるか？
	check := func(x int) bool {
		var s int
		for _, pi := range p {
			k := sort.Search(m, func(i int) bool {
				return (2*i-1)*pi > x
			})
			k--
			s += k * k * pi
		}
		return s <= m
	}
	ok, ng := 0, m
	for ng-ok > 1 {
		mid := (ok + ng) / 2
		if check(mid) {
			ok = mid
		} else {
			ng = mid
		}
	}
	var ans int
	for _, pi := range p {
		k := sort.Search(m, func(i int) bool {
			return (2*i-1)*pi > ok
		})
		k--
		ans += k
	}
	return ans
}

func solve01(n, m int, p []int) int {
	sort.Ints(p)
	fmt.Println(p)
	type node struct {
		i, v int
	}
	q := priorityqueue.New[node](func(a, b node) int {
		if a.v == b.v {
			return 0
		}
		if a.v < b.v {
			return -1
		}
		return 1
	})
	for i, pi := range p {
		for k := 1; k*k*pi <= m; k++ {
			q.Push(node{i, k * k * pi})
		}
	}
	r, s := make([]int, n), make([]int, n)
	var t int
	for {
		cur := q.Pop()
		if t-s[cur.i]+cur.v > m {
			break
		}
		t = t - s[cur.i] + cur.v
		r[cur.i]++
		s[cur.i] = cur.v
	}
	fmt.Println(t, r, s)
	var ans int
	for _, v := range r {
		ans += v
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
