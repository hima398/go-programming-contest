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

	t := nextInt()
	var a, s []int
	for i := 0; i < t; i++ {
		a = append(a, nextInt())
		s = append(s, nextInt())
	}
	ans := solve(t, a, s)
	PrintVertically(ans)
}

func solve(t int, a, s []int) []string {
	var ans []string
	for i := 0; i < t; i++ {
		ok := a[i]<<1 <= s[i]
		ok = ok && (s[i]-a[i]<<1)&a[i] == 0
		if ok {
			ans = append(ans, "Yes")
		} else {
			ans = append(ans, "No")
		}
	}
	return ans
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func PrintVertically(x []string) {
	defer out.Flush()
	for _, v := range x {
		fmt.Fprintln(out, v)
	}
}
