package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func solve(k int) []int {
	const maxN = 50
	mk := k % 50
	ans := make([]int, maxN)
	for i := 0; i < mk; i++ {
		ans[i] = maxN
	}
	for i := mk; i < maxN; i++ {
		ans[i] = maxN - mk - 1
	}
	dk := k / 50
	for i := range ans {
		ans[i] += dk
	}

	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	k := nextInt()
	ans := solve(k)
	PrintInt(len(ans))
	PrintHorizonaly(ans)
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

func PrintHorizonaly(x []int) {
	defer out.Flush()
	fmt.Fprintf(out, "%d", x[0])
	for i := 1; i < len(x); i++ {
		fmt.Fprintf(out, " %d", x[i])
	}
	fmt.Fprintln(out)
}
