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

	n := nextInt()
	ans := solve(n)
	PrintInt(ans)
}

func solve(n int) int {
	var ans int
	for i := 1; i <= n; i++ {
		k := i
		for x := 2; x*x <= k; x++ {
			for k%(x*x) == 0 {
				k /= x * x
			}
		}
		var s int
		for y := 1; k*y*y <= n; y++ {
			s++
		}
		ans += s
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
