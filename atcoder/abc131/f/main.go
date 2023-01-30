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

	n := nextInt()
	var x, y []int
	for i := 0; i < n; i++ {
		x = append(x, nextInt())
		y = append(y, nextInt())
	}
	ans := solve(n, x, y)
	PrintInt(ans)
}

func solve(n int, x, y []int) int {
	max := int(1e5)
	e := make([][]int, 2*max+1)
	for i := 0; i < n; i++ {
		y := y[i] + max
		e[x[i]] = append(e[x[i]], y)
		e[y] = append(e[y], x[i])
	}
	visited := make([]bool, 2*max+1)
	var l, r int
	var dfs func(cur int)
	dfs = func(cur int) {
		if visited[cur] {
			return
		}
		visited[cur] = true
		if cur <= max {
			l++
		} else {
			r++
		}
		for _, next := range e[cur] {
			dfs(next)
		}
	}
	var ans int
	for _, xi := range x {
		l, r = 0, 0
		dfs(xi)
		//fmt.Println("x, l, r = ", xi, l, r)
		ans += l * r
	}
	ans -= n
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
