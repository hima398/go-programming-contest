package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/liyue201/gostl/ds/queue"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, m := nextInt(), nextInt()
	var u, v, w []int
	for i := 0; i < m; i++ {
		u = append(u, nextInt()-1)
		v = append(v, nextInt()-1)
		w = append(w, nextInt())
	}

	ans := solve(n, m, u, v, w)

	//検算
	for i := 0; i < m; i++ {
		if ans[v[i]]-ans[u[i]] != w[i] {
			fmt.Printf("(u, v) = (%d, %d): w = %d, ans = %d", u[i], v[i], w[i], ans[v[i]]-ans[u[i]])
		}
	}

	PrintHorizonaly(ans)
}

type edge struct {
	t, w int
}

func solve(n, m int, u, v, w []int) []int {
	e := make([][]edge, n)
	for i := 0; i < m; i++ {
		e[u[i]] = append(e[u[i]], edge{v[i], w[i]})
		e[v[i]] = append(e[v[i]], edge{u[i], -w[i]})
	}
	q := queue.New[int]()
	visited := make([]bool, n)

	ans := make([]int, n)
	for i := 0; i < n; i++ {
		if visited[i] {
			continue
		}
		q.Push(i)
		visited[i] = true
		for !q.Empty() {
			cur := q.Pop()
			for _, next := range e[cur] {
				if visited[next.t] {
					continue
				}
				q.Push(next.t)
				ans[next.t] = ans[cur] + next.w
				visited[next.t] = true
			}
		}
	}
	return ans
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func Print(x any) {
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
