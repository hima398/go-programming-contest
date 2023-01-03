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
	var l, r []int
	for i := 0; i < n; i++ {
		l = append(l, nextInt())
		r = append(r, nextInt())
	}
	ans := solve(n, l, r)
	PrintVertically(ans)
}

func solve(n int, l, r []int) [][2]int {
	f := make([]int, 2*int(1e5)+1)
	for i := 0; i < n; i++ {
		f[l[i]]++
		f[r[i]]--
	}
	for i := 1; i <= 2*int(1e5); i++ {
		f[i] += f[i-1]
	}
	var ans [][2]int
	isInArea := 0
	for i := 1; i <= 2*int(1e5); i++ {
		if isInArea == 0 && f[i] > 0 {
			isInArea = 1
			ans = append(ans, [2]int{i, 0})
		} else if isInArea == 1 && f[i] == 0 {
			ans[len(ans)-1][1] = i
			isInArea = 0
		}
	}
	return ans
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func PrintVertically(x [][2]int) {
	defer out.Flush()
	for _, v := range x {
		fmt.Fprintln(out, v[0], v[1])
	}
}
