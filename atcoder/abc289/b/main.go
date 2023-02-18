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
	a := nextIntSlice(m)
	e := make([][]int, n)
	for i := 0; i < m; i++ {
		a[i]--
		b := a[i] + 1
		e[a[i]] = append(e[a[i]], b)
		e[b] = append(e[b], a[i])
	}
	visited := make([]bool, n)
	var stack []int
	var dfs func(cur int)
	dfs = func(cur int) {
		if visited[cur] {
			return
		}
		visited[cur] = true
		stack = append(stack, cur)
		for _, next := range e[cur] {
			if visited[next] {
				continue
			}
			dfs(next)
		}
	}
	var ans []int
	for i := 0; i < n; i++ {
		if visited[i] {
			continue
		}
		dfs(i)
		for len(stack) > 0 {
			ans = append(ans, stack[len(stack)-1]+1)
			stack = stack[:len(stack)-1]
		}
	}
	PrintHorizonaly(ans)

}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func nextIntSlice(n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = nextInt()
	}
	return s
}
func PrintInt(x int) {
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
