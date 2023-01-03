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

	n, q := nextInt(), nextInt()
	s := nextString()
	var t, x []int
	for i := 0; i < q; i++ {
		t = append(t, nextInt())
		x = append(x, nextInt())
	}
	ans := solve(n, q, s, t, x)
	PrintVertically(ans)
}

func solve(n, q int, s string, t, x []int) []string {
	offset := 0
	var ans []string
	for i := 0; i < q; i++ {
		switch t[i] {
		case 1:
			offset = (offset - x[i] + n) % n
		case 2:
			idx := (offset + x[i] - 1 + n) % n
			ans = append(ans, string(s[idx]))
		}
	}
	return ans
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func PrintString(x string) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}

func PrintVertically(x []string) {
	defer out.Flush()
	for _, v := range x {
		fmt.Fprintln(out, v)
	}
}
