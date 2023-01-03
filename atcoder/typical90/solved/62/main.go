package main

import (
	"bufio"
	"errors"
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
		a = append(a, nextInt()-1)
		b = append(b, nextInt()-1)
	}
	ans, err := solve(n, a, b)
	if err != nil {
		PrintInt(-1)
		return
	}
	PrintVertically(ans)
}

func solve(n int, a, b []int) ([]int, error) {
	e := make([][]int, n)
	var q []int
	visited := make([]bool, n)
	for i := 0; i < n; i++ {
		e[a[i]] = append(e[a[i]], i)
		e[b[i]] = append(e[b[i]], i)
		if a[i] == i || b[i] == i {
			q = append(q, i)
			visited[i] = true
		}
	}
	var ans []int
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		ans = append(ans, cur+1)
		for _, next := range e[cur] {
			if visited[next] {
				continue
			}
			q = append(q, next)
			visited[next] = true
		}
	}
	if len(ans) != n {
		return nil, errors.New("Impossible")
	}
	for i := 0; i < len(a)/2; i++ {
		j := n - 1 - i
		ans[i], ans[j] = ans[j], ans[i]
	}
	return ans, nil
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

func PrintVertically(x []int) {
	defer out.Flush()
	for _, v := range x {
		fmt.Fprintln(out, v)
	}
}
