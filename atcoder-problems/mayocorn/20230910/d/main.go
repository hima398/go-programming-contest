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
	const max = int(1e6)
	e := make([][]int, n)
	for i := 0; i < m; i++ {
		e[u[i]] = append(e[u[i]], v[i])
		e[v[i]] = append(e[v[i]], u[i])
	}
	var ans int
	visited := make([]bool, n)
	var dfs func(cur int)
	dfs = func(cur int) {
		//答えは最大でもmaxまで
		if ans >= max {
			return
		}
		visited[cur] = true
		ans++
		//ここが10以下程度
		for _, next := range e[cur] {
			if visited[next] {
				continue
			}
			dfs(next)
		}
		visited[cur] = false
	}
	//パスを全探索
	dfs(0)
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
