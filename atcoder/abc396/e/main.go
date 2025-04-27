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
	var x, y, z []int
	for i := 0; i < m; i++ {
		x = append(x, nextInt()-1)
		y = append(y, nextInt()-1)
		z = append(z, nextInt())
	}

	ans, err := solve(n, m, x, y, z)

	if err != nil {
		Print(-1)
	} else {
		PrintHorizonaly(ans)
	}
}

type edge struct {
	t, w int
}

var visited []bool
var candidate []int

func bfs(root int, e [][]edge) ([]int, error) {
	var res []int
	q := queue.New[int]()
	q.Push(root)
	visited[root] = true
	res = append(res, root)
	for !q.Empty() {
		cur := q.Pop()
		for _, next := range e[cur] {
			if visited[next.t] {
				//TODO:問題の矛盾をチェックする
				if candidate[next.t] != candidate[cur]^next.w {
					return nil, errors.New("impossible")
				}
				continue
			}
			q.Push(next.t)
			visited[next.t] = true
			candidate[next.t] = candidate[cur] ^ next.w
			res = append(res, next.t)
		}
	}
	return res, nil
}

func solve(n, m int, x, y, z []int) ([]int, error) {
	e := make([][]edge, n)
	for i := 0; i < m; i++ {
		e[x[i]] = append(e[x[i]], edge{y[i], z[i]})
		e[y[i]] = append(e[y[i]], edge{x[i], z[i]})
	}
	visited = make([]bool, n)
	candidate = make([]int, n)

	ans := make([]int, n)
	for i := 0; i < n; i++ {
		if visited[i] {
			continue
		}
		s, err := bfs(i, e)
		if err != nil {
			return nil, errors.New("impossible")
		}
		for shift := 0; shift <= 32; shift++ {
			var sum int
			for _, j := range s {
				if candidate[j]&(1<<shift) > 0 {
					sum++
				}
			}
			if sum <= len(s)-sum {
				//集合sのスタートを0にする方が最小になる
				for _, j := range s {
					ans[j] |= candidate[j] & (1 << shift)
				}
			} else {
				//集合sのスタートを1にする方が最小になる
				for _, j := range s {
					ans[j] |= candidate[j]&(1<<shift) ^ (1 << shift)
				}
			}
		}
	}
	return ans, nil
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
