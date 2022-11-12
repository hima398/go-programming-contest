package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func solve(n int, l, r []int) int {
	type period struct {
		l, r int
	}
	var ps []period
	for i := 0; i < n; i++ {
		ps = append(ps, period{l[i], r[i]})
	}
	sort.Slice(ps, func(i, j int) bool {
		if ps[i].r == ps[j].r {
			return ps[i].l < ps[j].l
		}
		return ps[i].r < ps[j].r
	})
	var t int
	var ans int
	for _, p := range ps {
		if t > p.l {
			continue
		}
		ans++
		t = p.r
	}
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	var l, r []int
	for i := 0; i < n; i++ {
		l = append(l, nextInt())
		r = append(r, nextInt())
	}
	ans := solve(n, l, r)
	PrintInt(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func PrintInt(x int) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
