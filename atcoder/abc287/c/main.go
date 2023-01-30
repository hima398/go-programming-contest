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
	PrintString(ans)
}

func solve(n, m int, u, v []int) string {
	if n-m != 1 {
		return "No"
	}
	d := make([]int, n)
	e := make([][]int, n)
	for i := range u {
		d[u[i]]++
		d[v[i]]++
		e[u[i]] = append(e[u[i]], v[i])
		e[v[i]] = append(e[v[i]], u[i])
	}
	no, nt := 0, 0
	for _, v := range d {
		if v == 1 {
			no++
		} else if v == 2 {
			nt++
		}
	}
	if !(no == 2 && nt == n-2) {
		return "No"
	}
	visited := make([]bool, n)
	var dfs func(cur int)
	dfs = func(cur int) {
		if visited[cur] {
			return
		}
		visited[cur] = true
		for _, next := range e[cur] {
			dfs(next)
		}
	}
	dfs(0)
	ok := true
	for i := range visited {
		ok = ok && visited[i]
	}
	if ok {
		return "Yes"
	} else {
		return "No"
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func PrintString(x string) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
