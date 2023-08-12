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
	solve(n, m)
}

func solve(n, m int) {
	visited := make([]bool, n)
	var dfs func(cur int)
	dfs = func(cur int) {
		if cur == n-1 {
			fmt.Println("OK")
			//return
			os.Exit(0)
		}
		visited[cur] = true
		k := nextInt()
		v := nextIntSlice(k)
		for _, next := range v {
			if visited[next] {
				continue
			}
			fmt.Println(next + 1)
			dfs(next)
		}
	}
	dfs(0)
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

func PrintFloat64(x float64) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}

func PrintString(x string) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
