package main

import (
	"bufio"
	"errors"
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
	var a, b []int
	var c []string
	for i := 0; i < m; i++ {
		a = append(a, nextInt()-1)
		b = append(b, nextInt()-1)
		c = append(c, nextString())
	}

	ans, err := solve(n, m, a, b, c)

	if err != nil {
		Print(-1)
	} else {
		Print(ans)
	}
}

func solve(n, m int, a, b []int, c []string) (int, error) {
	type edge struct {
		to int
		c  string
	}
	e := make([][]edge, n)
	connected := make([][]bool, n)
	for i := range connected {
		connected[i] = make([]bool, n)
	}
	for i := 0; i < m; i++ {
		e[a[i]] = append(e[a[i]], edge{b[i], c[i]})
		e[b[i]] = append(e[b[i]], edge{a[i], c[i]})
		connected[a[i]][b[i]] = true
		connected[b[i]][a[i]] = true
	}
	const INF = 1 << 60
	dist := make([][]int, n)
	for i := range dist {
		dist[i] = make([]int, n)
		for j := range dist[i] {
			dist[i][j] = INF
		}
	}
	q := queue.New[[2]int]()
	q.Push([2]int{0, n - 1})
	dist[0][n-1] = 0
	for !q.Empty() {
		cur := q.Pop()
		for _, n1 := range e[cur[0]] {
			for _, n2 := range e[cur[1]] {
				if dist[n1.to][n2.to] < INF {
					continue
				}
				if n1.c == n2.c {
					q.Push([2]int{n1.to, n2.to})
					dist[n1.to][n2.to] = dist[cur[0]][cur[1]] + 1
				}
			}
		}
	}
	ans := INF
	for i := 0; i < n; i++ {
		ans = Min(ans, 2*dist[i][i])
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if connected[i][j] {
				ans = Min(ans, 2*dist[i][j]+1)
			}
			if connected[j][i] {
				ans = Min(ans, 2*dist[j][i]+1)
			}
		}
	}
	if ans == INF {
		return ans, errors.New("impossible")
	} else {
		return ans, nil
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func Print(x any) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
