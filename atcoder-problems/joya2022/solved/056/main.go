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
	var a, b []int
	for i := 0; i < n; i++ {
		a = append(a, nextInt())
		b = append(b, nextInt())
	}
	ans := solve(n, a, b)
	PrintInt(ans)
}

func solve(n int, a, b []int) int {
	e := make(map[int][]int)
	for i := 0; i < n; i++ {
		e[a[i]] = append(e[a[i]], b[i])
		e[b[i]] = append(e[b[i]], a[i])
	}
	v := make(map[int]struct{})
	var q []int
	q = append(q, 1)
	v[1] = struct{}{}
	ans := 1
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		for _, next := range e[cur] {
			if _, found := v[next]; found {
				continue
			}
			q = append(q, next)
			v[next] = struct{}{}
			ans = Max(ans, next)
		}
	}
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

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
