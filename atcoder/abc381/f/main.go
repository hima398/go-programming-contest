package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"sort"
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
	a := nextIntSlice(n)

	ans := solve(n, a)

	Print(ans)
}

func solve(n int, a []int) int {
	const maxA = 20
	INF := n + 1

	var ks []int
	for k := 0; k < 1<<maxA; k++ {
		ks = append(ks, k)
	}
	sort.Slice(ks, func(i, j int) bool {
		if bits.OnesCount(uint(ks[i])) == bits.OnesCount(uint(ks[j])) {
			return ks[i] < ks[j]
		}
		return bits.OnesCount(uint(ks[i])) < bits.OnesCount(uint(ks[j]))
	})
	//fmt.Println(ks)

	//i番目以降にjが2つ表れる最小のインデックス
	idxes := make([][]int, n+2)
	for i := range idxes {
		idxes[i] = make([]int, maxA)
	}
	for j := 0; j < maxA; j++ {
		idxes[n+1][j] = INF
		idxes[n][j] = INF
	}
	for i := n - 1; i >= 0; i-- {
		for j := 0; j < maxA; j++ {
			idxes[i][j] = idxes[i+1][j]
		}
		idxes[i][a[i]-1] = i + 1
	}
	//for _, v := range idxes {
	//	Print(v)
	//}

	dp := make([]int, 1<<maxA)
	for i := range dp {
		dp[i] = INF
	}
	dp[0] = 0
	for _, k := range ks {
		for j := 0; j < maxA; j++ {
			if (k>>j)&1 > 0 {
				continue
			}
			next := k | (1 << j)
			i := idxes[dp[k]][j]
			dp[next] = Min(dp[next], idxes[i][j])
		}
	}
	var ans int
	for k, v := range dp {
		if v <= n {
			ans = Max(ans, 2*bits.OnesCount(uint(k)))
		}
	}
	return ans
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

func Print(x any) {
	defer out.Flush()
	fmt.Fprintln(out, x)
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
