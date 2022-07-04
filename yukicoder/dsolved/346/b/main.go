package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func solve(n int, a, b []int) int {
	canWin := make([]bool, n)
	for i := 1; i < n; i++ {
		if b[i-1] < b[i] {
			canWin[i] = a[i-1]+b[i] < a[i]+b[i-1]
		} else {
			canWin[i] = a[i-1]+b[i-1] < a[i]+b[i]
		}
	}
	//fmt.Println(canWin)
	//番兵
	canWin = append(canWin, false)

	var ans int
	var cnt int
	for i := 1; i < n; i++ {
		if canWin[i] {
			cnt++
		}
		if !canWin[i+1] {
			ans += cnt * (cnt + 1) / 2
			cnt = 0
		}
	}
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	a := nextIntSlice(n)
	b := nextIntSlice(n)

	ans := solve(n, a, b)
	PrintInt(ans)
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

func PrintInt(x int) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
