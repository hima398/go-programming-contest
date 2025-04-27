package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	h, w, n := nextInt(), nextInt(), nextInt()
	var r, c []int
	for i := 0; i < n; i++ {
		r = append(r, nextInt())
		c = append(c, nextInt())
	}

	numCoins, route := solve(h, w, n, r, c)

	Print(numCoins)
	Print(route)
}

type coin struct {
	r, c int
}

func solve(h, w, n int, r, c []int) (int, string) {
	const INF = 1 << 60
	var cs []coin
	for i := 0; i < n; i++ {
		cs = append(cs, coin{r[i], c[i]})
	}
	sort.Slice(cs, func(i, j int) bool {
		if cs[i].r == cs[j].r {
			return cs[i].c < cs[j].c
		}
		return cs[i].r < cs[j].r
	})
	dp := make([]int, n)
	for i := range dp {
		dp[i] = INF
	}
	id := make([]int, n)
	for i := range id {
		id[i] = -1
	}
	pre := make([]int, n)
	for i := 0; i < n; i++ {
		idx := sort.Search(n, func(j int) bool {
			return cs[i].c < dp[j]
		})
		dp[idx] = cs[i].c
		id[idx] = i
		if idx > 0 {
			pre[i] = id[idx-1]
		} else {
			pre[i] = -1
		}
	}
	numCoins := n - 1
	for id[numCoins] < 0 {
		numCoins--
	}

	var path []coin
	path = append(path, coin{h, w})
	cur := id[numCoins]
	for cur >= 0 {
		path = append(path, cs[cur])
		cur = pre[cur]
	}
	//fmt.Println(path)

	ci, cj := 1, 1
	var ans []string
	for i := len(path) - 1; i >= 0; i-- {
		ans = append(ans, strings.Repeat("D", path[i].r-ci))
		ans = append(ans, strings.Repeat("R", path[i].c-cj))
		ci, cj = path[i].r, path[i].c
	}

	return numCoins + 1, strings.Join(ans, "")
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

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
