package main

import (
	"bufio"
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
	var a, b []int
	for i := 0; i < m; i++ {
		a = append(a, nextInt()-1)
		b = append(b, nextInt()-1)
	}

	ans := solve(n, m, a, b)

	Print(ans)
}

func solve(n, m int, a, b []int) int {
	const p = int(1e9) + 7
	e := make([][]int, n)
	for i := 0; i < m; i++ {
		e[a[i]] = append(e[a[i]], b[i])
		e[b[i]] = append(e[b[i]], a[i])
	}

	q := queue.New[int]()
	q.Push(0)

	dist := make([]int, n)
	dp := make([]int, n)
	dp[0] = 1
	for !q.Empty() {
		cur := q.Pop()
		for _, next := range e[cur] {
			if dist[next] == 0 {
				dist[next] = dist[cur] + 1
				q.Push(next)
				dp[next] = dp[cur]
			} else if dist[next] == dist[cur]+1 {
				dp[next] += dp[cur]
				dp[next] %= p
			}
		}
	}
	return dp[n-1]
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
