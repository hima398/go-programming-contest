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
	PrintInt(ans)
}

func solve(n, m int, u, v []int) int {
	e := make([][]int, n)
	for i := 0; i < m; i++ {
		e[u[i]] = append(e[u[i]], v[i])
		e[v[i]] = append(e[v[i]], u[i])
	}

	colors := make([]int, n)
	var dfs func(cur, par, color int) (int, int, error)
	dfs = func(cur, par, color int) (int, int, error) {
		colors[cur] = color

		var b, w int
		if color == 1 {
			b++
		} else {
			w++
		}
		for _, next := range e[cur] {
			if next == par {
				continue
			}
			if colors[next] == -color {
				continue
			} else if colors[next] == color {
				return -1, -1, errors.New("is not bipartite")
			}
			nb, nw, err := dfs(next, cur, -color)
			if err != nil {
				return -1, -1, errors.New("is not bipartite")
			}
			b += nb
			w += nw
		}
		return b, w, nil
	}
	ans := n*(n-1)/2 - m
	for i := 0; i < n; i++ {
		if colors[i] != 0 {
			continue
		}
		b, w, err := dfs(i, -1, 1)
		if err != nil {
			return 0
		}
		ans -= b*(b-1)/2 + w*(w-1)/2
	}
	return ans
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func PrintInt(x int) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
