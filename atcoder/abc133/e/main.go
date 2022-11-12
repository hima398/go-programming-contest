package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func solve(n, k int, a, b []int) int {
	const p = int(1e9) + 7
	m := n - 1
	for i := 0; i < m; i++ {
		a[i]--
		b[i]--
	}
	e := make([][]int, n)
	for i := 0; i < m; i++ {
		e[a[i]] = append(e[a[i]], b[i])
		e[b[i]] = append(e[b[i]], a[i])
	}
	dp := make([]int, n)
	var dfs func(cur, par, kd int) int
	dfs = func(cur, par, kd int) int {
		//fmt.Println(cur, par, kd)
		s := kd
		//自分自身を塗る色を減らす
		used := 1
		if cur > 0 {
			used++
		}
		for _, next := range e[cur] {
			if next == par {
				continue
			}
			s *= dfs(next, cur, k-used)
			s %= p
			used++
		}
		dp[cur] = s
		return dp[cur]
	}
	root := 0
	ans := dfs(root, -1, k)
	//fmt.Println(dp)
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, k := nextInt(), nextInt()
	var a, b []int
	for i := 0; i < n-1; i++ {
		a = append(a, nextInt())
		b = append(b, nextInt())
	}
	ans := solve(n, k, a, b)
	PrintInt(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func PrintInt(x int) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
