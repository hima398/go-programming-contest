package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func solve(n, s int, a []int) ([]int, error) {
	const maxS = 60 * int(1e4)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, maxS+1)
	}
	dp[0][0] = 1
	for i := 1; i <= n; i++ {
		for j := 0; j <= maxS; j++ {
			dp[i][j] = dp[i-1][j]
			pj := j - a[i-1]
			if pj < 0 {
				continue
			}
			dp[i][j] += dp[i-1][pj]
		}
	}
	if dp[n][s] == 0 {
		return nil, errors.New("Impossible")
	}
	ps := s
	var ans []int
	for cur := n; cur > 0; cur-- {
		t := ps - a[cur-1]
		if t < 0 {
			continue
		}
		if dp[cur-1][t] > 0 {
			ps = t
			ans = append(ans, cur)
		}
	}
	sort.Slice(ans, func(i, j int) bool {
		return ans[i] < ans[j]
	})
	return ans, nil
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, s := nextInt(), nextInt()
	a := nextIntSlice(n)
	ans, err := solve(n, s, a)
	if err != nil {
		PrintInt(-1)
		return
	}

	PrintInt(len(ans))
	PrintHorizonaly(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
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

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
