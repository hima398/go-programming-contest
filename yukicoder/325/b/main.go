package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func solve(n int, a []int) int {
	if n == 1 {
		return 0
	}
	s := make([]int, n)
	for i := n - 2; i >= 0; i-- {
		if a[i] != a[i+1] {
			s[i] = s[i+1] + 1
		}
	}
	//fmt.Println(a)
	//fmt.Println(s)

	var ans int
	for i := 0; i < n-1; i++ {
		if a[i] == 1 {
			ans += s[i]
			//ロボットがNに到達していなければs[i]+1回進められる
			if n-i-1 > s[i] {
				ans++
			}
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
	ans := solve(n, a)
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
