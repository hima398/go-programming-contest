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

	ax, ay, bx, by := nextInt(), nextInt(), nextInt(), nextInt()
	n := nextInt()
	var x, y []int
	for i := 0; i < n; i++ {
		x = append(x, nextInt())
		y = append(y, nextInt())
	}
	ans := solve(ax, ay, bx, by, n, x, y)
	PrintInt(ans)
}

func solve(ax, ay, bx, by, n int, x, y []int) int {
	type pos struct {
		x, y int
	}
	a := pos{ax, ay}
	b := pos{bx, by}
	f := func(a, b, c, d pos) bool {
		s1 := (a.x-b.x)*(c.y-a.y) - (a.y-b.y)*(c.x-a.x)
		s2 := (a.x-b.x)*(d.y-a.y) - (a.y-b.y)*(d.x-a.x)
		if s1*s2 > 0 {
			return false
		}
		s3 := (c.x-d.x)*(a.y-c.y) - (c.y-d.y)*(a.x-c.x)
		s4 := (c.x-d.x)*(b.y-c.y) - (c.y-d.y)*(b.x-c.x)
		return s3*s4 < 0
	}
	var ans int
	for i := 0; i < n; i++ {
		c := pos{x[i], y[i]}
		next := (i + 1) % n
		d := pos{x[next], y[next]}
		if f(a, b, c, d) {
			ans++
		}
	}
	ans = ans/2 + 1
	return ans
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
