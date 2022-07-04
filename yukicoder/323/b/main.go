package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func solveHonestly(n, k int) int {
	const INF = 1 << 60
	f := make([]int, n)
	f[k-1] = 1
	ans := 0
	for i := 2; i <= n; i++ {
		d := 0
		found := false
		g := make([]int, n)
		for j := range g {
			g[j] = INF
		}
		for j := 0; j < n; j++ {
			if f[j] > 0 {
				g[j] = 0
				found = true
				d = 0
			} else if found {
				g[j] = Min(g[j], d)
			}
			d++
		}
		d = 0
		found = false
		for j := n - 1; j >= 0; j-- {
			if f[j] > 0 {
				g[j] = 0
				found = true
				d = 0
			} else if found {
				g[j] = Min(g[j], d)
			}
			d++
		}
		idx := -1
		x := 0
		for j := 0; j < n; j++ {
			if x < g[j] {
				idx = j
				x = g[j]
			}
		}
		f[idx] = i
		if i == n {
			ans = idx + 1
		}
		//fmt.Println("i = ", i, " g = ", g)
	}
	//PrintHorizonaly(f)
	return ans
}

func solve(n, k int) int {
	if n == 1 {
		return 1
	}
	ans := n - 1
	if n-1 == k {
		ans = n
	}
	return ans
}
func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, k := nextInt(), nextInt()

	//ans := solveHonestly(n, k)
	ans := solve(n, k)

	PrintInt(ans)
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

func PrintHorizonaly(x []int) {
	defer out.Flush()
	fmt.Fprintf(out, "%d", x[0])
	for i := 1; i < len(x); i++ {
		fmt.Fprintf(out, " %d", x[i])
	}
	fmt.Fprintln(out)
}

func PrintVertically(x []int) {
	defer out.Flush()
	for _, v := range x {
		fmt.Fprintln(out, v)
	}
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
