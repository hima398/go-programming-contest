package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func solve(n, m int, a, b []int) string {
	e := make([][]int, n)
	for i := 0; i < m; i++ {
		e[a[i]] = append(e[a[i]], b[i])
		e[b[i]] = append(e[b[i]], a[i])
	}
	type node struct {
		i, d int
	}
	d := make([]int, n)
	var q []node
	q = append(q, node{0, 1})
	d[0] = 1
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		for _, next := range e[cur.i] {
			if d[next] > 0 {
				continue
			}
			q = append(q, node{next, cur.d + 1})
			d[next] = cur.d + 1
		}
	}
	if d[n-1]-d[0] == 2 {
		return "POSSIBLE"
	} else {
		return "IMPOSSIBLE"
	}
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, m := nextInt(), nextInt()
	var a, b []int
	for i := 0; i < m; i++ {
		a = append(a, nextInt()-1)
		b = append(b, nextInt()-1)
	}
	ans := solve(n, m, a, b)
	PrintString(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func PrintString(x string) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
