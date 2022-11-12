package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func solve(n, x, m int, a []int) string {
	var use, magic int
	for i := n - 1; i >= 0; i-- {
		for a[i]/(1<<magic) >= x {
			magic++
			use += i + 1
		}
	}
	if use <= m {
		return "Yes"
	} else {
		return "No"
	}
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, x, m := nextInt(), nextInt(), nextInt()
	a := nextIntSlice(n)
	ans := solve(n, x, m, a)
	PrintString(ans)

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

func PrintString(x string) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
