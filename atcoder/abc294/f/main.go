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

	n, m, k := nextInt(), nextInt(), nextInt()
	var a, b []int
	for i := 0; i < n; i++ {
		a = append(a, nextInt())
		b = append(b, nextInt())
	}
	var c, d []int
	for i := 0; i < m; i++ {
		c = append(c, nextInt())
		d = append(d, nextInt())
	}
	ans := solve(n, m, k, a, b, c, d)
	PrintFloat64(ans)
}

func solve(n, m, k int, a, b, c, d []int) float64 {
	ng, ok := 0.0, 1.0
	check := func(x float64) bool {
		s := x / (1 - x)
		var u []float64
		for i := 0; i < m; i++ {
			u = append(u, float64(c[i])-float64(d[i])*s)
		}
		sort.Float64s(u)
		var res int
		for i := 0; i < n; i++ {
			v := float64(a[i]) - float64(b[i])*s
			idx := sort.Search(m, func(j int) bool {
				return -v <= u[j]
			})
			res += m - idx
		}
		return res < k
	}
	for ok-ng > 1e-12 {
		mid := (ok + ng) / 2.0
		if check(mid) {
			ok = mid
		} else {
			ng = mid
		}
	}
	return 100.0 * ok
}

func solveHonestly(n, m, k int, a, b, c, d []int) float64 {
	type sugarWater struct {
		s, w int
	}
	var ss []sugarWater
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			ss = append(ss, sugarWater{a[i] + c[j], a[i] + b[i] + c[j] + d[j]})
		}
	}
	//fmt.Println(ss)
	var ans []float64
	for _, v := range ss {
		ans = append(ans, float64(v.s)/float64(v.w))
	}
	sort.Float64s(ans)
	return 100.0 * ans[len(ans)-k]
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func PrintFloat64(x float64) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
