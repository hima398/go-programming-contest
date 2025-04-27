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

func main() {
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	var a, b []int
	for i := 0; i < n; i++ {
		a, b = append(a, nextInt()), append(b, nextInt())
	}

	ans := solve(n, a, b)

	PrintHorizonaly(ans)
}

func solve(n int, a, b []int) []int {
	type period struct {
		k, v int
	}
	var ps []period
	for i := 0; i < n; i++ {
		ps = append(ps, period{a[i], 1})
		ps = append(ps, period{a[i] + b[i], -1})
	}
	sort.Slice(ps, func(i, j int) bool {
		return ps[i].k < ps[j].k
	})

	ans := make([]int, n+1)
	var sum int
	for i := 0; i < len(ps)-1; i++ {
		sum += ps[i].v
		ans[sum] += ps[i+1].k - ps[i].k
	}
	ans = ans[1:]
	return ans
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func PrintHorizonaly(x []int) {
	defer out.Flush()
	fmt.Fprintf(out, "%d", x[0])
	for i := 1; i < len(x); i++ {
		fmt.Fprintf(out, " %d", x[i])
	}
	fmt.Fprintln(out)
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
