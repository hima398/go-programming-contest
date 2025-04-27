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
	var u, v []int
	for i := 0; i < m; i++ {
		u = append(u, nextInt()-1)
		v = append(v, nextInt()-1)
	}

	ans := solve(n, m, u, v)

	Print(ans)
}

func solve(n, m int, u, v []int) int {
	e := make([][]int, n)
	for i := 0; i < m; i++ {
		e[u[i]] = append(e[u[i]], v[i])
		e[v[i]] = append(e[v[i]], u[i])
	}

	visited := make([]int, n)
	for i := range visited {
		visited[i] = -1
	}
	bfs := func(start int) ([2]int, error) {
		var res [2]int
		q := queue.New[int]()
		q.Push(start)
		visited[start] = 0
		res[0]++

		for !q.Empty() {
			cur := q.Pop()
			for _, next := range e[cur] {
				if visited[next] >= 0 {
					if visited[next] == visited[cur] {
						return [2]int{-1, -1}, errors.New("not bipartite")
					}
					continue
				}
				q.Push(next)
				visited[next] = visited[cur] ^ 1
				res[visited[next]]++
			}
		}
		return res, nil
	}

	ans := n*(n-1)/2 - m
	for i := 0; i < n; i++ {
		if visited[i] >= 0 {
			continue
		}
		res, err := bfs(i)
		if err != nil {
			return 0
		} else {
			ans -= res[0] * (res[0] - 1) / 2
			ans -= res[1] * (res[1] - 1) / 2
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
