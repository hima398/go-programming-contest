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
	var t []string
	var a []int
	for i := 0; i < n; i++ {
		t = append(t, nextString())
		a = append(a, nextInt())
	}
	ans := solve(n, t, a)
	PrintVertically(ans)
}

func solve(n int, t []string, a []int) []int {
	const p = 10000
	x := 0
	var ans []int
	for i := 0; i < n; i++ {
		switch t[i] {
		case "+":
			x = (x + a[i]) % p
		case "-":
			x = (x - a[i] + p) % p
		case "*":
			x = x * a[i] % p
		}
		ans = append(ans, x)
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

func PrintVertically(x []int) {
	defer out.Flush()
	for _, v := range x {
		fmt.Fprintln(out, v)
	}
}
