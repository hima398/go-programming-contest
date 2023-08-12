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
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	a := nextIntSlice(n)
	s := make([]int, n)
	mx := make([]int, n)
	s[0] = a[0]
	mx[0] = Max(a[0], 0)
	for i := 1; i < n; i++ {
		s[i] = s[i-1] + a[i]
		mx[i] = Max(s[i-1]+a[i], mx[i-1])
	}
	ans := 0
	cur := 0
	for i := 0; i < n; i++ {
		ans = Max(ans, cur+mx[i])
		cur += s[i]
	}
	PrintInt(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
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

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
