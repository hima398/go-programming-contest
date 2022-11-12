package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func solve(n, k int, a []int) []int {
	ok, ng := 0, 1<<60
	check := func(x int) bool {
		var kd int
		for _, ai := range a {
			kd += Min(x, ai)
		}
		return kd <= k
	}
	for ng-ok > 1 {
		mid := (ng + ok) / 2
		if check(mid) {
			ok = mid
		} else {
			ng = mid
		}
	}

	ans := make([]int, n)
	rem := k
	for i := range ans {
		diff := Min(ok, a[i])
		ans[i] = a[i] - diff
		rem -= diff
	}

	var idx int
	for rem > 0 {
		for ans[idx] <= 0 {
			idx = (idx + 1) % n
		}
		ans[idx]--
		rem--
		idx = (idx + 1) % n
	}

	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, k := nextInt(), nextInt()
	a := nextIntSlice(n)
	ans := solve(n, k, a)
	PrintHorizonaly(ans)
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

func PrintInt(x int) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}

func PrintHorizonaly(x []int) {
	defer out.Flush()
	fmt.Fprintf(out, "%d", x[0])
	for i := 1; i < len(x); i++ {
		fmt.Fprintf(out, " %d", x[i])
	}
	fmt.Fprintln(out)
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
