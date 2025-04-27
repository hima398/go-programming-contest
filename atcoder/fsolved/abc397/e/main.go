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
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, k := nextInt(), nextInt()
	var u, v []int
	for i := 0; i < n*k-1; i++ {
		u = append(u, nextInt()-1)
		v = append(v, nextInt()-1)
	}

	if solve(n, k, u, v) {
		Print("Yes")
	} else {
		Print("No")
	}
}

func solve(n, k int, u, v []int) bool {
	e := make([][]int, n*k)
	m := n*k - 1
	for i := 0; i < m; i++ {
		e[u[i]] = append(e[u[i]], v[i])
		e[v[i]] = append(e[v[i]], u[i])
	}
	decomposed := make([]bool, n*k)
	l := make([]int, n*k)
	var dfs func(cur, par int) bool
	dfs = func(cur, par int) bool {
		ok := true
		var numChildren int
		for _, next := range e[cur] {
			if next == par {
				continue
			}
			ok = ok && dfs(next, cur)
			if !decomposed[next] {
				numChildren++
			}
		}
		if !ok {
			return false
		}
		switch numChildren {
		case 0:
			l[cur]++
			if l[cur] == k {
				decomposed[cur] = true
			}
			return true
		case 1:
			l[cur]++
			for _, next := range e[cur] {
				if next == par {
					continue
				}
				if decomposed[next] {
					continue
				}
				l[cur] += l[next]
			}
			if l[cur] == k {
				decomposed[cur] = true
			}
		case 2:
			l[cur]++
			for _, next := range e[cur] {
				if next == par {
					continue
				}
				if decomposed[next] {
					continue
				}
				l[cur] += l[next]
			}
			if l[cur] == k {
				decomposed[cur] = true
			} else {
				return false
			}
		default:
			// numChildren >= 3
			return false
		}
		return true
	}

	ans := dfs(0, -1)
	//fmt.Println(l)
	//fmt.Println(decomposed)
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
