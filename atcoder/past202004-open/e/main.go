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
	a := nextIntSlice(n)
	ans := solve(n, a)
	PrintHorizonaly(ans)
}

func solve(n int, a []int) []int {
	for i := range a {
		a[i]--
	}
	e := make([][]int, n)
	for i, v := range a {
		e[i] = append(e[i], v)
	}
	//fmt.Println(e)
	bfs := func(start int) int {
		var q []int
		d := make([]int, n)
		for i := range d {
			d[i] = -1
		}
		q = append(q, start)
		d[start] = 0
		for len(q) > 0 {
			cur := q[0]
			q = q[1:]
			for _, next := range e[cur] {
				if d[next] > 0 {
					continue
				}
				d[next] = d[cur] + 1
				q = append(q, next)
			}
		}
		//fmt.Println(start, d)
		return d[start]
	}
	var ans []int
	for i := 0; i < n; i++ {
		ans = append(ans, bfs(i))
	}
	return ans
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

func PrintHorizonaly(x []int) {
	defer out.Flush()
	fmt.Fprintf(out, "%d", x[0])
	for i := 1; i < len(x); i++ {
		fmt.Fprintf(out, " %d", x[i])
	}
	fmt.Fprintln(out)
}
