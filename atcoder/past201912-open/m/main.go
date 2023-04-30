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
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, m := nextInt(), nextInt()
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
	ans := solve(n, m, a, b, c, d)
	PrintFloat64(ans)
}

func solve(n, m int, a, b, c, d []int) float64 {
	ok, ng := 0.0, 1e18
	check := func(x float64) bool {
		var normal []float64
		for i := 0; i < n; i++ {
			normal = append(normal, float64(b[i])-x*float64(a[i]))
		}
		//sort.Slice(normal, func(i, j int) bool {
		//	return normal[i] > normal[j]
		//})
		//fmt.Println("normal = ", normal)
		sort.Sort(sort.Reverse(sort.Float64Slice(normal)))

		var helper []float64
		for i := 0; i < m; i++ {
			helper = append(helper, float64(d[i])-x*float64(c[i]))
		}
		//sort.Slice(helper, func(i, j int) bool {
		//	return helper[i] > helper[j]
		//})
		//fmt.Println("helper = ", helper)
		sort.Sort(sort.Reverse(sort.Float64Slice(helper)))

		s1 := 0.0
		for i := 0; i < 5; i++ {
			s1 += normal[i]
		}
		//fmt.Println("s1 = ", s1)
		if s1 >= 0 {
			return true
		}
		s2 := 0.0
		for i := 0; i < 4; i++ {
			s2 += normal[i]
		}
		s2 += helper[0]
		//fmt.Println("s2 = ", s2)
		if s2 >= 0.0 {
			return true
		}

		return false
	}
	for i := 0; i < 100; i++ {
		mid := (ng + ok) / 2.0
		//fmt.Println(i, ok, ng, check(mid))
		if check(mid) {
			ok = mid
		} else {
			ng = mid
		}
	}
	return ok
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextIntSlice(n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = nextInt()
	}
	return s
}

func PrintFloat64(x float64) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
