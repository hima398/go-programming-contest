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

	n, l := nextInt(), nextInt()
	a := nextIntSlice(n)

	ans := solve(n, l, a)
	PrintString(ans)
}

func solve(n, l int, a []int) string {
	cur := 0
	for _, ai := range a {
		if cur+ai > l && ai == 2 {
			return "No"
		}
		cur += ai + 1
	}
	return "Yes"
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
