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

	n, k := nextInt(), nextInt()
	var s []string
	for i := 0; i < n; i++ {
		s = append(s, nextString())
	}
	ans := solve(n, k, s)
	PrintInt(ans)
}

func solve(n, k int, s []string) int {
	m := make([]map[rune]struct{}, n)
	for i := 0; i < n; i++ {
		m[i] = make(map[rune]struct{})
		for _, r := range s[i] {
			m[i][r] = struct{}{}
		}
	}
	pattern := 1 << n
	var ans int
	for mask := 0; mask < pattern; mask++ {
		m2 := make(map[rune]int)
		for i := 0; i < n; i++ {
			if mask>>i&1 == 0 {
				continue
			}
			for k := range m[i] {
				m2[k]++
			}
		}
		s := 0
		for _, v := range m2 {
			if v == k {
				s++
			}
		}
		ans = Max(ans, s)
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
