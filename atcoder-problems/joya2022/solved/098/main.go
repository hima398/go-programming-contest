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

	n := nextInt()
	ans := solve(n)
	PrintInt(ans)
}

func solve(n int) int {
	f := func(a, b int) int {
		return a*a*a + a*a*b + a*b*b + b*b*b
	}
	ans := int(1e18) + 1
	for a := 0; a <= int(1e6); a++ {
		b := sort.Search(int(1e6)+1, func(idx int) bool {
			return f(a, idx) >= n
		})
		ans = Min(ans, f(a, b))
	}
	return ans
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

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
