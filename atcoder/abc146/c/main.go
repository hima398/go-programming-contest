package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	a, b, x := nextInt(), nextInt(), nextInt()
	ans := solve(a, b, x)
	PrintInt(ans)
}

func solve(a, b, x int) int {
	ok, ng := 0, int(1e9)+1
	d := func(n int) int {
		sn := strconv.Itoa(n)
		return len(sn)
	}
	check := func(n int) bool {
		return x >= a*n+b*d(n)
	}
	for ng-ok > 1 {
		mid := (ok + ng) / 2
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

func PrintInt(x int) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
