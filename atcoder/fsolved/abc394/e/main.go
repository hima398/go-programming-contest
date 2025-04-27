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

	n := nextInt()
	var c []string
	for i := 0; i < n; i++ {
		c = append(c, nextString())
	}
	ans := solve(n, c)
	for _, v := range ans {
		PrintHorizonaly(v)
	}
}

func solve(n int, c []string) [][]int {
	ans := make([][]int, n)
	for i := 0; i < n; i++ {
		ans[i] = make([]int, n)
		for j := range ans[i] {
			ans[i][j] = -1
		}
	}

	q := queue.New[[2]int]()
	for i := 0; i < n; i++ {
		q.Push([2]int{i, i})
		ans[i][i] = 0
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			if c[i][j] != '-' {
				q.Push([2]int{i, j})
				ans[i][j] = 1
			}
		}
	}
	for !q.Empty() {
		cur := q.Pop()
		for ni := 0; ni < n; ni++ {
			for nj := 0; nj < n; nj++ {
				if ans[ni][nj] >= 0 {
					continue
				}
				if c[ni][cur[0]] == '-' || c[cur[1]][nj] == '-' {
					continue
				}
				if c[ni][cur[0]] == c[cur[1]][nj] {
					q.Push([2]int{ni, nj})
					ans[ni][nj] = ans[cur[0]][cur[1]] + 2
				}
			}
		}
	}

	return ans
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func PrintHorizonaly(x []int) {
	defer out.Flush()
	fmt.Fprintf(out, "%d", x[0])
	for i := 1; i < len(x); i++ {
		fmt.Fprintf(out, " %d", x[i])
	}
	fmt.Fprintln(out)
}
